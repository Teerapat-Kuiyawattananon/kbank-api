package handler

import (
	model "kapi/model"
	db "kapi/progresql"
	repo "kapi/repository"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

var billInput model.BillDetailRequest
func HandlerCreateBillDetail(c echo.Context) error {
	if err := c.Bind(&billInput) ; err != nil {
		log.Fatal(err)
	}

	clientDB, err := db.InitDatabase()
	if err != nil {
		log.Fatal(err)
	}
	
	billDetailRepo := repo.NewBillDetailRepository(clientDB)
	billRepo := repo.NewBillRepository(clientDB)

	billDetail, err := billDetailRepo.CreateBillDetail(billInput)
	if err != nil {
		log.Fatal(err)
	}

	bill := model.Bill{
		BillerId: billInput.BillerId,
		Ref1: billDetail.CustomerID,
		Ref2: billDetail.ID,
	}
	if err := billRepo.CreateBill(bill) ; err != nil {
		log.Fatal(err)
	}

	return nil
}

func HandlerGetAllBillDetails(c echo.Context) error {
	var billDetails []model.BillDetail
	clientDB, err := db.InitDatabase()
	if err != nil {
		log.Fatal(err)
	}

	billDetail := repo.NewBillDetailRepository(clientDB)
	entBillDetails, err := billDetail.GetAllBills()
	if err != nil {
		log.Println(err)
	}

	for _, bill := range entBillDetails {
		billDetails = append(billDetails, model.BillDetail{
			ID: bill.ID,
			ChannelCode: bill.ChannelCode,
			SenderBankCode: bill.SenderBankCode,
			Status: bill.Status,
			CustomerId: bill.CustomerID,
			TranAmount: bill.TranAmount,
			CreatedAt: bill.CreatedAt,
			UpdatedAt: bill.UpdatedAt,
		})
	}

	return c.JSON(http.StatusOK, billDetails)
}