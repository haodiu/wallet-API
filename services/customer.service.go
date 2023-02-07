package services

import "wallet/models"

type CustomerService interface {
	CreateCustomer(*models.Customer) error
	GetCustomer(string) (*models.Customer, error)
	GetAll() ([]models.Customer, error)
	UpdateCustomer(string, *models.Customer) error
	DeleteCustomer(string) error
}