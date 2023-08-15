package customer

import "time"

type Customer struct {
	ID 			int			`json:"id"`
	FirstName	string		`json:"firstName"`
	LastName	string		`json:"lastName"`
	TitleName 	string		`json:"titleName"`
	MobileNumber string		`json:"mobileNumber"`
	CreatedAt 	time.Time	`json:"createdAt"`
}