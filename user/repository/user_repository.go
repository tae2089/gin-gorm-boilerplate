package repository

import (
	"github.com/tae2089/gin-boilerplate/common/config"
	domain "github.com/tae2089/gin-boilerplate/user/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	Begin() (tx *gorm.DB)
	Save(user *domain.User) error
	FindById(id uint) (*domain.User, error)
	FindByEmail(email string) (*domain.User, error)
	FindAll() ([]*domain.User, error)
	Delete(user *domain.User) error
}

func NewUserRepository() UserRepository {
	return &userRepositoryImpl{}
}

type userRepositoryImpl struct {
	client *gorm.DB
}

func (*userRepositoryImpl) Begin() (tx *gorm.DB) {
	client := config.GetDB()
	tx = client.Begin()
	return tx
}

func (*userRepositoryImpl) Save(user *domain.User) error {
	client := config.GetDB()
	return client.Save(user).Error
}

func (*userRepositoryImpl) FindByEmail(email string) (*domain.User, error) {
	var user domain.User
	client := config.GetDB()
	err := client.Where("email =?", email).First(&user).Error
	return &user, err
}

func (*userRepositoryImpl) FindById(id uint) (*domain.User, error) {
	var user domain.User
	client := config.GetDB()
	if err := client.Where("id =?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (*userRepositoryImpl) FindAll() ([]*domain.User, error) {
	var users []*domain.User
	client := config.GetDB()
	if err := client.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (*userRepositoryImpl) Delete(user *domain.User) error {
	client := config.GetDB()
	return client.Delete(user).Error
}
