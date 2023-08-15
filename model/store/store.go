package store

type Store struct {
	ID 			int
	AccountName string  `json:"account_name"`
	ServiceName string  `json:"service_name"`
}