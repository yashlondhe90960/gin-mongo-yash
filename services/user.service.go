package services

import "example.com/yash-apis/models"

type UserService interfeace {
	CreateUser(*models.User) error
	GetUser(*string) (*models.User, error)
    	GetAll() ([] *models.User, error)
	UpdateUser(*models.User) error
	DeleteUser(*string) error
}
