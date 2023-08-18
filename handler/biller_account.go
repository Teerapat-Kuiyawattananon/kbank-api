package handler

import (
	model "kapi/model"
	repo "kapi/repository"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

var billAccountInput model.BillAccount

// @Summary 	Create new biller_account.
// @Description Create new biller_account, it's for store.
// @Tags 		Billers
// @Accept 		*/*
// @Produce 	json
// @Param body 	body model.BillAccount true "JSON request body for payment request"
// @Success 	201 {array} model.BillAccount "Success"
// @Router 		/api/billers [post]
func HandlerCreateBillerAccount(c echo.Context) error {
	if err := c.Bind(&billAccountInput) ; err != nil {
		log.Fatal(err)
	}
	
	billerAccountRepo := repo.NewBillerAccountRepository(clientDB)

	billerAccount, err := billerAccountRepo.CreateBillerAccount(billAccountInput)
	if err != nil {
		log.Println(err)
	}

	billerAccountRes := model.BillAccount{
		ID: billerAccount.ID,
		Name: billerAccount.Name,
		ServiceName: billerAccount.ServiceName,
	}

	return c.JSON(http.StatusCreated, billerAccountRes)
}

// @Summary 	Show all biller_account.
// @Description Show all biller_account.
// @Tags 		Billers
// @Accept 		*/*
// @Produce 	json
// @Success 	200 {array} model.BillAccount "Success"
// @Router 		/api/billers [get]
func HandlerGetBillers(c echo.Context) error {
	var stores []model.BillAccount
	
	billerAccountRepo := repo.NewBillerAccountRepository(clientDB)

	entBillerAccount, err := billerAccountRepo.GetBillers()
	if err != nil {
		log.Panicln(err)
	}

	for _, biller := range entBillerAccount {
		stores = append(stores, model.BillAccount{
			ID: biller.ID,
			Name: biller.Name,
			ServiceName: biller.ServiceName,
		})
	}

	return c.JSON(http.StatusOK, stores)
}