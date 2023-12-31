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

// @Summary 	Create new customer.
// @Description Create new customer that use bill.
// @Tags 		Customers
// @Accept 		*/*
// @Produce 	json
// @Param body 	body model.Customer true "JSON request body for payment request"
// @Success 	201 {array} model.Customer "Success"
// @Failure     400 {string} string "Create Customer failed"
// @Router 		/api/customers [post]
func HandlerCreateCustomer(c echo.Context) error {
	if err := c.Bind(&customerInput); err != nil {
		log.Fatal(err)
	}

	cusRepo := repo.NewCustomerRepository(clientDB)

	cus, err := cusRepo.CreateCustomer(customerInput)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"Error message" : "Create Customer failed",
		})
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

// @Summary 	Show all customer.
// @Description Show all customer in system.
// @Tags 		Customers
// @Accept 		*/*
// @Produce 	json
// @Success 	200 {array} model.Customer "Success"
// @Failure	 	500 {string} string "Customer not found"
// @Router 		/api/customers [get]
func HandlerGetCustomers(c echo.Context) error {
	var customers []model.Customer

	cusRepo := repo.NewCustomerRepository(clientDB)
	entCustomers, err := cusRepo.GetCustomers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"Error message" : "Customer not found",
		})
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

// @Summary 	Show customer by ID.
// @Description Show customer by ID in system.
// @Tags 		Customers
// @Accept 		*/*
// @Produce 	json
// @Param body 	body model.Customer true "JSON request body for payment request"
// @Success 	200 {array} model.Customer "Success"
// @Failure     400 {string} string "Customer by ID failed"
// @Failure	 	500 {string} string "Customer by ID not found"
// @Router 		/api/customers/:id [get]
func HandlerGetCustomerByID(c echo.Context) error {
	var customer model.Customer

	cusRepo := repo.NewCustomerRepository(clientDB)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"Error message" : "Customer by ID failed",
		})
	}

	entCus, err := cusRepo.GetCustomerByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"Error message" : "Customer by ID not found",
		})
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
