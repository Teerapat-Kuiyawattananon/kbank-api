package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	h "kapi/handler"
	echoSwagger "github.com/swaggo/echo-swagger"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"

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
// @schemes http

//go:generate swag init -g docs\docs.go


func main() {
	if err := godotenv.Load(".env") ; err != nil {
		log.Fatal(err)
	}
	fmt.Println("Start")

	e := echo.New()

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.GET("/api/test", func(c echo.Context) error {
		return c.JSON(http.StatusOK, echo.Map{
			"data" : "tset",
			"test" : "1",
		})
	})

	e.POST("/api/billpayment/lookup", h.HandlerLookup)
	e.POST("/api/billpayment/payment", h.HandlerPayment)

	e.POST("/api/billpayment/create/bill", h.CreateBill)
	e.POST("/api/billpayment/create/customer", h.CreateCustomer)


	e.Start(":" + os.Getenv("PORT"))
}