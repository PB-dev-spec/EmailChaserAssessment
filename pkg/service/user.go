package service

import (
	"mark_emailchaser/pkg/models"
	"mark_emailchaser/pkg/repository"
)

func CreateUser(newUser models.User) (*models.User, error) {
	if err := repository.CreateUser(&newUser); err != nil {
		return nil, err
	}

	return &newUser, nil
}

func GetUserByID(id uint) (*models.User, error) {
	return repository.FindUserByID(id)
}

func GetUserByEmail(email string) (*models.User, error) {
	return repository.FindUserByEmail(email)
}

func DeleteUser(id uint) error {
	return repository.DeleteUser(id)
}
