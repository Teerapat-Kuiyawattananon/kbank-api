package model

type Bill struct {
	ID int			
	BillerId int	`json:"biller_id"`
	Ref1 int 		`json:"ref1"`
	Ref2 int		`json:"ref2"`
}