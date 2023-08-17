package handler

import (
	// mStore "kapi/model/store"
	model "kapi/model"
	db "kapi/progresql"
	repo "kapi/repository"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

var input model.Bill

// @Summary 	Create new bill.
// @Description Create new bill detail.
// @Tags 		Bills
// @Accept 		*/*
// @Produce 	json
// @Param body 	body model.Bill true "JSON request body for payment request"
// @Success 	201 {array} model.Bill "Success"
// @Router 		/api/bills [post]
func HandlerCreateBill(c echo.Context) error {
	var bill model.Bill
	if err := c.Bind(&input); err != nil {
		log.Fatal(err)
	}

	clientDB, err := db.InitDatabase()
	if err != nil {
		log.Fatal(err)
	}

	billRepo := repo.NewBillRepository(clientDB)

	entBill, err := billRepo.CreateBill(input)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"Error message" : "Create Bill failed",
		})
	}
	
	bill = model.Bill{
		ID:             entBill.ID,
		BillerId:       entBill.BillerID,
		Reference1:     entBill.Reference1,
		Reference2:     entBill.Reference2,
		TransactionID:  entBill.TransactionID,
		ChannelCode:    entBill.ChannelCode,
		SenderBankCode: entBill.SenderBankCode,
		Status:         entBill.Status,
		TranAmount:     entBill.TranAmount,
		CreatedAt:      entBill.CreatedAt,
		UpdatedAt:      entBill.UpdatedAt,
	}
	return c.JSON(http.StatusOK, bill)
}

// @Summary 	Show all bill.
// @Description show all bill in system.
// @Tags 		Bills
// @Accept 		*/*
// @Produce 	json
// @Success 	200 {array} model.Bill "Success"
// @Router 		/api/bills [get]
func HandlerGetAllBills(c echo.Context) error {
	var bills []model.Bill
	clientDB, err := db.InitDatabase()
	if err != nil {
		log.Fatal(err)
	}

	bill := repo.NewBillRepository(clientDB)
	entBills, err := bill.GetAllBills()
	if err != nil {
		log.Println(err)
	}

	for _, bill := range entBills {
		bills = append(bills, model.Bill{
			ID:             bill.ID,
			BillerId:       bill.BillerID,
			Reference1:     bill.Reference1,
			Reference2:     bill.Reference2,
			TransactionID:  bill.TransactionID,
			ChannelCode:    bill.ChannelCode,
			SenderBankCode: bill.SenderBankCode,
			Status:         bill.Status,
			TranAmount:     bill.TranAmount,
			CreatedAt:      bill.CreatedAt,
			UpdatedAt:      bill.UpdatedAt,
		})
	}

	return c.JSON(http.StatusOK, bills)
}

func HandlerUpdateBill(c echo.Context) error {
	var bill model.Bill
	if err := c.Bind(&input); err != nil {
		log.Fatal(err)
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err, id)
	}

	clientDB, err := db.InitDatabase()
	if err != nil {
		log.Fatal(err)
	}

	billRepo := repo.NewBillRepository(clientDB)
	entBill, err := billRepo.UpdateBill(input, id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"Error message" : "Create Bill failed",
		})
	}

	bill = model.Bill{
		ID:             entBill.ID,
		BillerId:       entBill.BillerID,
		Reference1:     entBill.Reference1,
		Reference2:     entBill.Reference2,
		TransactionID:  entBill.TransactionID,
		ChannelCode:    entBill.ChannelCode,
		SenderBankCode: entBill.SenderBankCode,
		Status:         entBill.Status,
		TranAmount:     entBill.TranAmount,
		CreatedAt:      entBill.CreatedAt,
		UpdatedAt:      entBill.UpdatedAt,
	}

	return c.JSON(http.StatusOK, bill)
}