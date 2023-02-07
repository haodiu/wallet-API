package services

import "wallet/models"

type TransactionService interface {
	CreateTransaction(*models.Transaction) error
	GetTransactions(string) ([]models.Transaction, error)
	GetAll() ([]models.Transaction, error)
	DeleteTransaction(string) error
}
