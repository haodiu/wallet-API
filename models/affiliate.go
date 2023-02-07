package models

import "database/sql"

type Affiliate struct {
	ID          int            `json:"id,omitempty"`
	AffName     string         `json:"aff_name,omitempty"`
	District    string         `json:"district,omitempty"`
	Address     string         `json:"address,omitempty"`
	PhoneNumber string         `json:"phone_number,omitempty"`
	Fax         string         `json:"fax,omitempty"`
	Email       string         `json:"email,omitempty"`
	CreateAt    string         `json:"create_at"`
	UpdateAt    sql.NullString `json:"update_at"`
}
