package models

import "database/sql"

type Customer struct {
	ID          int            `json:"id,omitempty"`
	FirstName   string         `json:"firstname,omitempty"`
	LastName    string         `json:"lastname,omitempty"`
	DateOfBirth string         `json:"date_of_birth,omitempty"`
	Nationality string         `json:"nationality,omitempty"`
	Address     string         `json:"address,omitempty"`
	Balance     string         `json:"balance,omitempty"`
	Avatar      string         `json:"avatar"`
	CreateAt    string         `json:"create_at"`
	UpdateAt    sql.NullString `json:"update_at"`
}
