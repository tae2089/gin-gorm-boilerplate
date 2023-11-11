package service

import (
	"errors"

	"github.com/tae2089/bob-logging/logger"
	"github.com/tae2089/gin-boilerplate/user/dto"
	"github.com/tae2089/gin-boilerplate/user/model"
	"github.com/tae2089/gin-boilerplate/user/repository"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService interface {
	Join(requestJoin *dto.RequestJoin) (string, error)
	Login(requestLogin *dto.RequestLogin) (string, error)
}

type userServiceImpl struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userServiceImpl{
		userRepository: userRepository,
	}
}

func (u *userServiceImpl) Join(requestJoin *dto.RequestJoin) (string, error) {
	user, err := u.userRepository.FindByEmail(requestJoin.Email)
	if err == nil {
		return "", errors.New("already exists user.")
	}
	user = &model.User{
		Name:     requestJoin.Username,
		Email:    requestJoin.Email,
		Phone:    requestJoin.Phone,
		Roles:    []string{"editor", "viewer"},
		Password: requestJoin.Password,
	}
	err = u.userRepository.Save(user)
	if err != nil {
		logger.Error(err)
		return "", errors.New("saving user failed.")
	}
	return "success", nil
}

func (s *userServiceImpl) Login(requestLogin *dto.RequestLogin) (string, error) {
	user, err := s.userRepository.FindByEmail(requestLogin.Email)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return "", errors.New("user not found")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(requestLogin.Password))
	if err != nil {
		logger.Error(err)
		return "", errors.New("wrong password")
	}
	return user.ID.String(), nil
}
