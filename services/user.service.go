package services

import "example.com/yash-apis/models"

type UserService interfeace {
	CreateUser(*models.User) error
	GetUser(*string) (*models.User, error)
    	
}
