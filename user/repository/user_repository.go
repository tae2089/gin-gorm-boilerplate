package repository

import (
	"github.com/tae2089/gin-boilerplate/common/config"
	domain "github.com/tae2089/gin-boilerplate/user/model"
	"gorm.io/gorm"
)

func Begin() (tx *gorm.DB) {
	client := config.GetDB()
	tx = client.Begin()
	return tx
}

func Save(user *domain.User) error {
	client := config.GetDB()
	return client.Save(user).Error
}

func FindByEmail(email string) (*domain.User, error) {
	var user domain.User
	client := config.GetDB()
	err := client.Where("email =?", email).First(&user).Error
	return &user, err
}

func FindById(id uint) (*domain.User, error) {
	var user domain.User
	client := config.GetDB()
	if err := client.Where("id =?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func FindAll() ([]*domain.User, error) {
	var users []*domain.User
	client := config.GetDB()
	if err := client.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func Delete(user *domain.User) error {
	client := config.GetDB()
	return client.Delete(user).Error
}
