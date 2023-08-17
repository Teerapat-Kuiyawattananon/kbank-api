package handler

import (
	model "kapi/model"
	repo "kapi/repository"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

var (
	customerInput model.Customer
)

func HandlerCreateCustomer(c echo.Context) error {
	if err := c.Bind(&customerInput); err != nil {
		log.Fatal(err)
	}

	cusRepo := repo.NewCustomerRepository(clientDB)

	cus, err := cusRepo.CreateCustomer(customerInput)
	if err != nil {
		log.Fatal(err)
	}
	customer := model.Customer{
		ID:           cus.ID,
		FirstName:    cus.FirstName,
		LastName:     cus.LastName,
		TitleName:    cus.TitleName,
		MobileNumber: cus.MobileNumber,
		CreatedAt:    cus.CreatedAt,
	}
	return c.JSON(http.StatusCreated, customer)
}

func HandlerGetCustomers(c echo.Context) error {
	var customers []model.Customer

	cusRepo := repo.NewCustomerRepository(clientDB)
	entCustomers, err := cusRepo.GetCustomers()
	if err != nil {
		log.Println(err)
	}

	for _, customer := range entCustomers {
		customers = append(customers, model.Customer{
			ID:           customer.ID,
			FirstName:    customer.FirstName,
			LastName:     customer.LastName,
			TitleName:    customer.TitleName,
			MobileNumber: customer.MobileNumber,
			CreatedAt:    customer.CreatedAt,
		})
	}

	return c.JSON(http.StatusOK, customers)
}

func HandlerGetCustomerByID(c echo.Context) error {
	var customer model.Customer

	cusRepo := repo.NewCustomerRepository(clientDB)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
	}

	entCus, err := cusRepo.GetCustomerByID(id)
	if err != nil {
		log.Println(err)
	}
	if (entCus == nil) {
		return c.NoContent(http.StatusOK)
	}
	if (entCus != nil) {
		customer = model.Customer{
			ID: entCus.ID,
			FirstName:    entCus.FirstName,
			LastName:     entCus.LastName,
			TitleName:    entCus.TitleName,
			MobileNumber: entCus.MobileNumber,
			CreatedAt:    entCus.CreatedAt,
		}
	}

	return c.JSON(http.StatusOK, customer)
}
