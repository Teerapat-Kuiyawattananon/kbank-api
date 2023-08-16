package handler

import (
	model "kapi/model"
	db	"kapi/progresql"
	repo "kapi/repository"
	"log"

	"github.com/labstack/echo/v4"
)

var customerInput model.Customer
func HandlerCreateCustomer(c echo.Context) error {
	if err := c.Bind(&customerInput) ; err != nil {
		log.Fatal(err)
	}

	clientDB, err := db.InitDatabase()
	if err != nil {
		log.Fatal(err)
	}
	
	cusRepo := repo.NewCustomerRepository(clientDB)

	if err := cusRepo.CreateCustomer(customerInput) ; err != nil {
		log.Fatal(err)
	}

	return nil
}