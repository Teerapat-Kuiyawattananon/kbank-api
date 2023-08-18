package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	h "kapi/handler"
	db "kapi/progresql"
	echoSwagger "github.com/swaggo/echo-swagger"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	_ "kapi/docs"
)

// @title Bill Payment API
// @version 1.0
// @description This is Bill Payment API Server From KBank-API.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
// @schemes http http 
func main() {
	if err := godotenv.Load(".env") ; err != nil {
		log.Fatal(err)
	}
	fmt.Println("Start")
	db.MockUpTestEX01()
	e := echo.New()

	   // Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.GET("/swagger/*", echoSwagger.WrapHandler)

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
	e.POST("/api/bills", h.HandlerCreateBill)
	e.GET("/api/bills", h.HandlerGetAllBills)

	// -CUSTOMER
	e.POST("/api/customers", h.HandlerCreateCustomer)
	e.GET("/api/customers", h.HandlerGetCustomers)
	e.GET("/api/customers/:id", h.HandlerGetCustomerByID)

	// -BILLER_ACCOUNT
	e.POST("/api/billers",h.HandlerCreateBillerAccount)
	e.GET("/api/billers", h.HandlerGetBillers)
	
	e.Start(":" + os.Getenv("PORT"))
}