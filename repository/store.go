package repository

import (
	"context"
	"kapi/ent"
	"log"

	mStore "kapi/model/store"
)


type storeRepository struct {
	clientDB *ent.Client
}

func NewStoreRepository(DB *ent.Client) storeRepository {
	return storeRepository{
		clientDB: DB,
	}
}

func (repo storeRepository) CreateStore(input mStore.Store) (error) {
	store, err := repo.clientDB.Store.Create().
					SetAccountName(input.AccountName).
					SetServiceName(input.ServiceName).
					Save(context.Background())
	if err != nil {
		log.Fatal(err)
		return err
	}

	log.Printf("Created Store id: %d success", store.ID)
	return nil
}