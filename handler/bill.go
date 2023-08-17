package handler

import (
	// mStore "kapi/model/store"
	model "kapi/model"
	db "kapi/progresql"
	repo "kapi/repository"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

var input model.Bill

func HandlerCreateBill(c echo.Context) error {
	if err := c.Bind(&input); err != nil {
		log.Fatal(err)
	}

	clientDB, err := db.InitDatabase()
	if err != nil {
		log.Fatal(err)
	}

	billRepo := repo.NewBillRepository(clientDB)

	if err := billRepo.CreateBill(input); err != nil {
		log.Fatal(err)
	}

	return nil
}

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
