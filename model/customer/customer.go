package customer

import "time"

type Customer struct {
	ID 			int
	FirstName	string
	LastName	string
	TitleName 	string
	MobileNumber string
	CreatedAt 	time.Time
}