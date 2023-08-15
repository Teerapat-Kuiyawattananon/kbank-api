package handler


import (
	"fmt"
	model_bill "kapi/model/bill"
	model_customer "kapi/model/customer"

	repo "kapi/repository"
	"strconv"

	db "kapi/progresql"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

var inputCustomer model_customer.Customer


func CreateBill(c echo.Context) error {

	if err := c.Bind(&inquiryRequest) ; err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	// convert bind input from string to int
	billerId, _ := strconv.Atoi(inquiryRequest.BillerId)
	ref1, _ := strconv.Atoi(inquiryRequest.Reference1)
	ref2, _ := strconv.Atoi(inquiryRequest.Reference2)


	clientDB, err := db.InitDatabase()
	if err != nil {
		log.Fatal("err when init db", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal server error"})
	}

	billRepo := repo.NewBillRepository(clientDB)

	// create new bill
	newBill := billRepo.CreateBill(model_bill.Bill{
		BillerId: billerId,
		Ref1: ref1,
		Ref2: ref2,
	})

	fmt.Println("new bill", newBill)
	return c.JSON(http.StatusCreated, newBill)
}


func CreateCustomer(c echo.Context) error {

	if err := c.Bind(&inputCustomer) ; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	clientDB, err := db.InitDatabase()
	if err != nil {
		log.Fatal("err when init db", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal server error"})
	}

	customerRepo := repo.NewCustomerRepository(clientDB)
	newCustomer := customerRepo.CreateCustomer(inputCustomer)

	fmt.Println("error after creted customer : ", newCustomer)
	return c.JSON(http.StatusCreated, map[string]string{"message": "Create customer success"})
}