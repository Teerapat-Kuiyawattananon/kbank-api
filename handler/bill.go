package handler

import (
	
	model "kapi/model"
	db	"kapi/progresql"
	repo "kapi/repository"
	"log"

	"github.com/labstack/echo/v4"
)

var input model.Bill
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
