package handler

import (
	model "kapi/model"
	repo "kapi/repository"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

var input model.Bill

func HandlerCreateBill(c echo.Context) error {
	var bill model.Bill
	if err := c.Bind(&input); err != nil {
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

func HandlerGetAllBills(c echo.Context) error {
	var bills []model.Bill

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
