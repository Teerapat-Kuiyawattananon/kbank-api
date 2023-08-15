package handler

import (
	// mStore "kapi/model/store"
	mBill "kapi/model/bill"
	db	"kapi/progresql"
	repo "kapi/repository"
	"log"

	"github.com/labstack/echo/v4"
)

var input mBill.Bill
func HandlerCreateBill(c echo.Context) error {
	if err := c.Bind(&input) ; err != nil {
		log.Fatal(err)
	}

	clientDB, err := db.InitDatabase()
	if err != nil {
		log.Fatal(err)
	}
	
	billRepo := repo.NewBillRepository(clientDB)

	if err := billRepo.CreateBill(input) ; err != nil {
		log.Fatal(err)
	}

	return nil
}