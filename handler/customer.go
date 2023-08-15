package handler

import (
	mCustomer "kapi/model/customer"
	db	"kapi/progresql"
	repo "kapi/repository"
	"log"

	"github.com/labstack/echo/v4"
)

var customerInput mCustomer.Customer
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