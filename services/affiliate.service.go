package services

import "wallet/models"

type AffiliateService interface {
	CreateAffiliate(*models.Affiliate) error
	GetAffiliate(string) (*models.Affiliate, error)
	GetAll() ([]models.Affiliate, error)
	UpdateAffiliate(string, *models.Affiliate) error
	DeleteAffiliate(string) error
}