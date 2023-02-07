package services

import (
	"context"
	"database/sql"
	"time"
	"wallet/models"
)

type AffiliateServiceImpl struct {
	affiliateDB *sql.DB
	ctx context.Context
}

func (a AffiliateServiceImpl) CreateAffiliate(affiliate *models.Affiliate) error {
	_, err := a.affiliateDB.ExecContext(a.ctx, "INSERT INTO affiliate (affiname, district, address, phoneNumber, fax, email, create_at) VALUES (?, ?, ?, ?, ?, ?, ?)", affiliate.AffName, affiliate.District, affiliate.Address, affiliate.PhoneNumber, affiliate.Fax, affiliate.Email, time.Now())
	if err != nil {
		return err
	}
	return nil
}

func (a AffiliateServiceImpl) GetAffiliate(s string) (*models.Affiliate, error) {
	panic("implement me")
}

func (a AffiliateServiceImpl) GetAll() ([]models.Affiliate, error) {
	var affiliates []models.Affiliate
	rows, err := a.affiliateDB.QueryContext(a.ctx, "SELECT * FROM affiliate")
	if err != nil {
		return affiliates, err
	}
	defer rows.Close()
	for rows.Next() {
		var affiliate models.Affiliate
		err := rows.Scan(&affiliate.ID, &affiliate.AffName, &affiliate.District, &affiliate.Address, &affiliate.PhoneNumber, &affiliate.Fax, &affiliate.Email, &affiliate.CreateAt, &affiliate.UpdateAt)
		if err != nil {
			return affiliates, err
		}
		affiliates = append(affiliates, affiliate)
	}
	return affiliates, nil
}

func (a AffiliateServiceImpl) UpdateAffiliate(id string, affiliate *models.Affiliate) error {
	_, err := a.affiliateDB.ExecContext(a.ctx, "UPDATE affiliate SET affiname = ?, district = ?, address = ?, phoneNumber = ?, fax= ?, email = ?, update_at = ? WHERE id = ?", affiliate.AffName, affiliate.District, affiliate.Address, affiliate.PhoneNumber, affiliate.Fax, affiliate.Email, time.Now(), id)
	if err != nil {
		return err
	}
	return nil
}

func (a AffiliateServiceImpl) DeleteAffiliate(id string) error {
	_, err := a.affiliateDB.ExecContext(a.ctx, "DELETE FROM affiliate WHERE id = ?", id)
	if err != nil {
		return err
	}
	return nil
}

func NewAffiliateService(affiliateDB *sql.DB, ctx context.Context) AffiliateService {
	return &AffiliateServiceImpl{
		affiliateDB: affiliateDB,
		ctx: ctx,
	}
}

