package services

import "wallet/models"

type DocumentService interface {
	CreateDocument(*models.Document) error
	GetDocuments(string) ([]models.Document, error)
	DeleteDocument(string) error
}
