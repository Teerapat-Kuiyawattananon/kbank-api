package customer

import "time"

type Customer struct {
	ID 			int				
	FirstName	string			`json:"first_name"`
	LastName	string			`json:"last_name"`
	TitleName 	string			`json:"title_name"`
	MobileNumber string			`json:"mobile_number"`
	CreatedAt 	time.Time
}