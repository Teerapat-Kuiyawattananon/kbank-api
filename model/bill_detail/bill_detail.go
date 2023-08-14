package billdetail

import "time"

type BillDetail struct {
	ID int
	ChannelCode string
	SenderBankCode string
	Status string
	Customer_id int
	TranAmount float64
	CreatedAt time.Time
	UpdatedAt time.Time
}