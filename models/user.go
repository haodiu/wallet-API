package models

import "database/sql"

type User struct {
	ID        int            `json:"id,omitempty"`
	Username  string         `json:"username,omitempty"`
	Email     string         `json:"email,omitempty"`
	Password  string         `json:"password,omitempty"`
	IsAdmin   bool           `json:"is_admin"`
	IsAddInfo bool           `json:"is_add_info"`
	CreateAt  string         `json:"create_at"`
	UpdateAt  sql.NullString `json:"update_at"`
}
