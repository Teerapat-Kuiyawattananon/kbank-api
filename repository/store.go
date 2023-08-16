package repository

import (
	"context"
	"kapi/ent"
	"kapi/ent/store"
	"log"

	mStore "kapi/model"
)


type storeRepository struct {
	clientDB *ent.Client
}

func NewStoreRepository(DB *ent.Client) storeRepository {
	return storeRepository{
		clientDB: DB,
	}
}

func (repo storeRepository) CreateStore(input mStore.Store) (*ent.Store, error) {
	store, err := repo.clientDB.Store.Create().
					SetAccountName(input.AccountName).
					SetServiceName(input.ServiceName).
					Save(context.Background())
	if err != nil {
		log.Println(err)
	}

	log.Printf("Created Store id: %d success", store.ID)
	return store, nil
}

func (repo storeRepository) GetStoreByID(id int) (*ent.Store, error) {
	store, err := repo.clientDB.Store.Query().
					Where(store.ID(id)).
					Only(context.Background())
	if err != nil {
		log.Println(err)
	}

	return store, nil
}

func (repo storeRepository) GetStores() ([]*ent.Store, error) {
	stores, err := repo.clientDB.Store.Query().
					All(context.Background())
	if err != nil {
		log.Panicln(err)
	}

	return stores, nil
}
