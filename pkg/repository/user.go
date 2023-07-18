package repository

import (
	"mark_emailchaser/pkg/common"
	"mark_emailchaser/pkg/models"
)

func CreateUser(user *models.User) error {
	return common.DB.Create(user).Error
}

func FindUserByID(id uint) (*models.User, error) {
	var user models.User
	err := common.DB.Where("id = ?", id).First(&user).Error
	return &user, err
}

func FindUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := common.DB.Where("email = ?", email).First(&user).Error
	return &user, err
}

func DeleteUser(id uint) error {
	return common.DB.Delete(&models.User{}, id).Error
}
