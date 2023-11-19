package service

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/tae2089/gin-boilerplate/common/domain"
	"github.com/tae2089/gin-boilerplate/common/util"
	oauthDomain "github.com/tae2089/gin-boilerplate/oauth/domain"
	"github.com/tae2089/gin-boilerplate/user/model"
	"github.com/tae2089/gin-boilerplate/user/repository"
	"golang.org/x/oauth2"
	"gorm.io/gorm"
)

type OauthProvider interface {
	GetAccessToken(code string) (*oauth2.Token, error)
	GetUserInfo(token *oauth2.Token) ([]byte, error)
	GetRedirectURL() (string, string)
}

const (
	OAUTH_GOOGLE_URL = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="
	OAUTH_GITHUB_URL = "https://api.github.com/user"
)

type OauthService struct {
	githubOauth OauthProvider
	googleOauth OauthProvider
	jwtUtil     util.JwtUtil
	userRepo    repository.UserRepository
}

func NewOauthService(githubOauth OauthProvider, googleOauth OauthProvider, jwtUtil util.JwtUtil, userRepository repository.UserRepository) *OauthService {
	return &OauthService{
		githubOauth: githubOauth,
		googleOauth: googleOauth,
		jwtUtil:     jwtUtil,
		userRepo:    userRepository,
	}
}

func (o *OauthService) GithubLogin() string {
	redirectURL, _ := o.githubOauth.GetRedirectURL()
	return redirectURL
}

func (o *OauthService) GoogleLogin() (string, string) {
	redirectURL, state := o.googleOauth.GetRedirectURL()
	return redirectURL, state
}

func (o *OauthService) GithubLoginCallback(code string) (domain.JwtToken, error) {
	oauthToken, err := o.githubOauth.GetAccessToken(code)
	if err != nil {
		return domain.JwtToken{}, err
	}
	body, err := o.githubOauth.GetUserInfo(oauthToken)
	if err != nil {
		return domain.JwtToken{}, err
	}
	var githubUserInfo oauthDomain.GithubUserInfo
	err = o.bindOAuthUserInfo(body, &githubUserInfo)
	if err != nil {
		return domain.JwtToken{}, err
	}
	id, err := o.findOrCreateUser(githubUserInfo.Email, githubUserInfo.Name)
	if err != nil {
		return domain.JwtToken{}, err
	}
	token, err := o.jwtUtil.CreateAccessToken(id.String(), true)
	if err != nil {
		return domain.JwtToken{}, err
	}
	return token, nil
}

func (o *OauthService) GoogleLoginCallback(code string) (domain.JwtToken, error) {
	oauthToken, err := o.googleOauth.GetAccessToken(code)
	if err != nil {
		return domain.JwtToken{}, err
	}
	body, err := o.googleOauth.GetUserInfo(oauthToken)
	if err != nil {
		return domain.JwtToken{}, err
	}
	var googleUserInfo oauthDomain.GoogleUserInfo
	err = o.bindOAuthUserInfo(body, &googleUserInfo)
	if err != nil {
		return domain.JwtToken{}, err
	}
	id, err := o.findOrCreateUser(googleUserInfo.Email, googleUserInfo.Name)
	if err != nil {
		return domain.JwtToken{}, err
	}
	token, err := o.jwtUtil.CreateAccessToken(id.String(), true)
	if err != nil {
		return domain.JwtToken{}, err
	}
	return token, nil
}

func (o *OauthService) bindOAuthUserInfo(body []byte, userInfo interface{}) error {
	if body == nil || len(body) == 0 {
		return errors.New("bindOAuthUserInfo: no data in body to unmarshal")
	}

	err := json.Unmarshal(body, userInfo)
	if err != nil {
		return fmt.Errorf("bindOAuthUserInfo: error unmarshaling data - %v", err)
	}

	return nil
}

func (o *OauthService) findOrCreateUser(email, name string) (uuid.UUID, error) {
	foundUser, err := o.userRepo.FindByEmail(email)
	if err == nil {
		return foundUser.ID, nil
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		newID := uuid.New()
		newUser := &model.User{
			ID:    newID,
			Email: email,
			Name:  name,
			Roles: []string{"user"},
		}
		err = o.userRepo.Save(newUser)
		if err != nil {
			return uuid.Nil, err
		}
		return newID, nil
	}

	return uuid.Nil, err
}
