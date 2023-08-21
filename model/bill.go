package model

import "time"

type Bill struct {
	ID 				int			`json:"id"`
	BillerId 		int			`json:"billerId"`
	Reference1 		int			`json:"reference1"`
	Reference2 		int			`json:"reference2"`
	TransactionID 	string 		`json:"trasactionId"`
	ChannelCode 	string		`json:"channelCode"`
	SenderBankCode 	string		`json:"senderBankCode"`
	Status 			string		`json:"status"`
	TranAmount 		float64		`json:"tranAmount"`
	CreatedAt 		time.Time	`json:"createdAt"`
	UpdatedAt 		time.Time	`json:"updatedAt"`
}

type DateParams struct {
	StartDate time.Time `query:"start_date" time_format:"2006-01-02"`
	EndDate   time.Time `query:"end_date" time_format:"2006-01-02"`
}
