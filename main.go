package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	h "kapi/handler"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	if err := godotenv.Load(".env") ; err != nil {
		log.Fatal(err)
	}
	fmt.Println("Start")

	e := echo.New()

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