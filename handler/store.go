package handler

import (
	model "kapi/model"
	db "kapi/progresql"
	repo "kapi/repository"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

var storeInput model.Store
func HandlerCreateStore(c echo.Context) error {
	if err := c.Bind(&storeInput) ; err != nil {
		log.Fatal(err)
	}

	clientDB, err := db.InitDatabase()
	if err != nil {
		log.Fatal(err)
	}
	
	storeRepo := repo.NewStoreRepository(clientDB)

	store, err := storeRepo.CreateStore(storeInput)
	if err != nil {
		log.Println(err)
	}

	storeRes := model.Store{
		ID: store.ID,
		AccountName: store.AccountName,
		ServiceName: store.ServiceName,
	}

	return c.JSON(http.StatusCreated, storeRes)
}

func HandlerGetStores(c echo.Context) error {
	var stores []model.Store
	clientDB, err := db.InitDatabase()
	if err != nil {
		log.Fatal(err)
	}
	
	storeRepo := repo.NewStoreRepository(clientDB)

	entStores, err := storeRepo.GetStores()
	if err != nil {
		log.Panicln(err)
	}

	for _, store := range entStores {
		stores = append(stores, model.Store{
			ID: store.ID,
			AccountName: store.AccountName,
			ServiceName: store.ServiceName,
		})
	}

	return c.JSON(http.StatusOK, stores)
}