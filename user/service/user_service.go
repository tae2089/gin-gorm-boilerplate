package service

import (
	"errors"

	"github.com/tae2089/bob-logging/logger"
	"github.com/tae2089/gin-boilerplate/user/dto"
	"github.com/tae2089/gin-boilerplate/user/model"
	userRepository "github.com/tae2089/gin-boilerplate/user/repository"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Join(requestJoin *dto.RequestJoin) (string, error) {
	user, err := userRepository.FindByEmail(requestJoin.Email)
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
	err = userRepository.Save(user)
	if err != nil {
		logger.Error(err)
		return "", errors.New("saving user failed.")
	}
	return "success", nil
}

func Login(requestLogin *dto.RequestLogin) (string, error) {
	user, err := userRepository.FindByEmail(requestLogin.Email)
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
