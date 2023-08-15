package bill

type Bill struct {
	ID int			`json:"id"`
	BillerId int	`json:"billerId"`
	Ref1 int		`json:"reference1"`
	Ref2 int		`json:"reference2"`
}