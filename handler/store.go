package handler


import (
	"fmt"
	model_bill "kapi/model/bill"
	model_customer "kapi/model/customer"
	model_store "kapi/model/store"
	model_bill_detail "kapi/model/bill_detail"

	repo "kapi/repository"
	"strconv"

	db "kapi/progresql"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

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
		log.Fatal(err)
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

	inputCustomer := model_customer.Customer{}

	if err := c.Bind(&inputCustomer) ; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	clientDB, err := db.InitDatabase()
	if err != nil {
		log.Fatal(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal server error"})
	}

	customerRepo := repo.NewCustomerRepository(clientDB)
	newCustomer := customerRepo.CreateCustomer(inputCustomer)

	fmt.Println("error after creted customer : ", newCustomer)
	return c.JSON(http.StatusCreated, map[string]string{"message": "Create customer success"})
}

func CreateStore(c echo.Context) error {
	inputStore := model_store.Store{}

	if err := c.Bind(&inputStore) ; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	clientDB, err := db.InitDatabase()
	if err != nil {
		log.Fatal(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal server error"})
	}

	storeRepo := repo.NewStoreRepository(clientDB)
	newStore := storeRepo.CreateStore(inputStore)

	fmt.Println("error after creted store : ", newStore)
	return c.JSON(http.StatusCreated, map[string]string{"message": "Create store success"})
}

func CreateBillDetail(c echo.Context) error {

	if err := c.Bind(&inquiryRequest) ; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	clientDB, err := db.InitDatabase()
	if err != nil {
		log.Fatal(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal server error"})
	}

	// convert string to float
	tranAmount, _ := strconv.ParseFloat(inquiryRequest.TranAmount, 64)

	billDetailRepo := repo.NewBillDetailRepository(clientDB)
	newBillDetail := billDetailRepo.CreateBillDetail(model_bill_detail.BillDetail{
		TranAmount: tranAmount,
		ChannelCode: inquiryRequest.ChannelCode,
		Customer_id: 1,
		SenderBankCode: inquiryRequest.SenderBankCode,
		Status: "waiting",
	})

	fmt.Println("error after creted bill detail : ", newBillDetail)
	return c.JSON(http.StatusCreated, map[string]string{"message": "Create bill detail success"})
}