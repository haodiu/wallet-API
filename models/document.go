package models

import "database/sql"

type Document struct {
	ID               int            `json:"id,omitempty"`
	DocType          string         `json:"doc_type,omitempty"`
	DocNumber        string         `json:"doc_number,omitempty"`
	IssuingAuthority string         `json:"issuing_authority,omitempty"`
	ExpiryDate       string         `json:"expiry_date,omitempty"`
	Img              string         `json:"img,omitempty"`
	UserID           int            `json:"user_id,omitempty"`
	CreateAt         string         `json:"create_at"`
	UpdateAt         sql.NullString `json:"update_at"`
}

//type NullString struct {
//	String string
//	Valid  bool // Valid is true if String is not NULL
//}