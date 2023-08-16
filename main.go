package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	h "kapi/handler"
	db "kapi/progresql"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	if err := godotenv.Load(".env") ; err != nil {
		log.Fatal(err)
	}
	fmt.Println("Start")
	db.MockUpTestEX01()
	e := echo.New()

	e.GET("/api/test", func(c echo.Context) error {
		return c.JSON(http.StatusOK, echo.Map{
			"data" : "tset",
			"test" : "1",
		})
	})
	
	// Kbank API (bill-payment-service)
	e.POST("/api/billpayment/lookup", h.HandlerLookup)
	e.POST("/api/billpayment/payment", h.HandlerPayment)

	// Store API
	// -BILL
	e.POST("/api/bills", h.HandlerCreateBillDetail)
	e.GET("/api/bills", h.HandlerGetAllBillDetails)
	// -CUSTOMER
	e.POST("/api/customers", h.HandlerCreateCustomer)
	e.GET("/api/customers", h.HandlerGetCustomers)
	e.GET("/api/customers/:id", h.HandlerGetCustomerByID)

	// -STORE
	e.POST("/api/stores",h.HandlerCreateStore)
	e.GET("/api/stores", h.HandlerGetStores)
	e.Start(":" + os.Getenv("PORT"))
}