package services

import "wallet/models"

type UserService interface {
	CreateUser(*models.User) error
	GetUser(string) (*models.User, error)
	CheckUser(*models.User) (*models.User, error)
	UpdateUser(string, *models.User) error
}
