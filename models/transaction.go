package models

type Transaction struct {
	ID           int    `json:"id,omitempty"`
	SenderName   string `json:"sender_name,omitempty"`
	ReceiverName string `json:"receiver_name,omitempty"`
	Date         string `json:"date,omitempty"`
	Money        int    `json:"money,omitempty"`
	Message      string `json:"message,omitempty"`
	Status       string `json:"status"`
	CreateAt     string `json:"create_at"`
}
