package model

import "time"

type BillDetail struct {
	ID int					`json:"id"`
	ChannelCode string		`json:"channelCode"`
	SenderBankCode string	`json:"senderBankCode"`
	Status string			`json:"status"`
	CustomerId int			`json:"customerId"`
	TranAmount float64		`json:"tranAmount"`
	CreatedAt time.Time		`json:"createdAt"`
	UpdatedAt time.Time		`json:"updatedAt"`
}

type BillDetailRequest struct {
	ID int					`json:"id"`
	ChannelCode string		`json:"channelCode"`
	SenderBankCode string	`json:"senderBankCode"`
	Status string			`json:"status"`
	CustomerId int			`json:"customerId"`
	TranAmount float64		`json:"tranAmount"`
	CreatedAt time.Time		`json:"createdAt"`
	UpdatedAt time.Time		`json:"updatedAt"`
	BillerId int			`json:"billerId"`
}