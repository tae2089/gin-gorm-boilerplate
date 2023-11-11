package repository

import (
	domain "github.com/tae2089/gin-boilerplate/user/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	Begin() (tx *gorm.DB)
	Save(user *domain.User) error
	FindById(id uint) (*domain.User, error)
	FindByEmail(email string) (*domain.User, error)
	FindAll() ([]*domain.User, error)
	Delete(userId string) error
}

func NewUserRepository(client *gorm.DB) UserRepository {
	return &userRepositoryImpl{
		client,
	}
}

type userRepositoryImpl struct {
	client *gorm.DB
}

func (u *userRepositoryImpl) Begin() (tx *gorm.DB) {
	tx = u.client.Begin()
	return tx
}

func (u *userRepositoryImpl) Save(user *domain.User) error {
	return u.client.Save(user).Error
}

func (u *userRepositoryImpl) FindByEmail(email string) (*domain.User, error) {
	var user domain.User
	err := u.client.Where("email =?", email).First(&user).Error
	return &user, err
}

func (u *userRepositoryImpl) FindById(id uint) (*domain.User, error) {
	var user domain.User
	if err := u.client.Where("id =?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *userRepositoryImpl) FindAll() ([]*domain.User, error) {
	var users []*domain.User
	if err := u.client.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (u *userRepositoryImpl) Delete(userId string) error {
	return u.client.Where("id =?", userId).Delete(&domain.User{}).Error
}
