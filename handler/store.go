package handler

import (
	model "kapi/model"
	db	"kapi/progresql"
	repo "kapi/repository"
	"log"

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

	if err := storeRepo.CreateStore(storeInput) ; err != nil {
		log.Fatal(err)
	}

	return nil
}