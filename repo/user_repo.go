package repo

import (
	"vioxcd/dpl/config"
	"vioxcd/dpl/models"
)

func AddUser(user *models.User) error {
	result := config.DB.Create(user)
	return result.Error
}

func GetUser(user *models.User) error {
	result := config.DB.Joins("Role").Find(user)
	return result.Error
}

func Login(user *models.User) error {
	result:= config.DB.Where("email = ? AND password = ?", user.Email, user.Password).Joins("Role").First(user)
	return result.Error
}
