package services

import (
	"github.com/rohankhatua/yt-fetcher/internals/db/models"
	"github.com/rohankhatua/yt-fetcher/internals/db"
)

type UserService struct {}

func (s* UserService) CreateUser(user *models.User) error {
	return db.DB.Create(user).Error
}

func (s* UserService) GetUserByID (id uint) (*models.User, error) {
	user := &models.User{}
	err := db.DB.First(user, id).Error
	return user, err
}

func (s* UserService) UpdateUser (user *models.User) error {
	return db.DB.Save(user).Error
}

func (s* UserService) DeleteUser (id uint) error {
	return db.DB.Delete(&models.User{}, id).Error
}