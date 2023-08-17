package handler

import (
	model "kapi/model"
	db "kapi/progresql"
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

	clientDB, err := db.InitDatabase()
	if err != nil {
		log.Fatal(err)
	}
	
	billerAccountRepo := repo.NewBillerAccountRepository(clientDB)

	billerAccount, err := billerAccountRepo.CreateStore(billAccountInput)
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
	clientDB, err := db.InitDatabase()
	if err != nil {
		log.Fatal(err)
	}
	
	billerAccountRepo := repo.NewBillerAccountRepository(clientDB)

	entBillerAccount, err := billerAccountRepo.GetStores()
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