package handler

import (
	model "kapi/model"
	repo "kapi/repository"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

var billAccountInput model.BillAccount
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