package repository

import (
	"context"
	"kapi/ent"
	"kapi/ent/biller_account"
	"log"

	model "kapi/model"
)


type billerAccountRepository struct {
	clientDB *ent.Client
}

func NewBillerAccountRepository(DB *ent.Client) billerAccountRepository {
	return billerAccountRepository{
		clientDB: DB,
	}
}

func (repo billerAccountRepository) CreateStore(input model.BillAccount) (*ent.Biller_account, error) {
	store, err := repo.clientDB.Biller_account.Create().
					SetName(input.Name).
					SetServiceName(input.ServiceName).
					Save(context.Background())
	if err != nil {
		log.Println(err)
	}

	log.Printf("Created Store id: %d success", store.ID)
	return store, nil
}

func (repo billerAccountRepository) GetStoreByID(id int) (*ent.Biller_account, error) {
	billerAccount, err := repo.clientDB.Biller_account.Query().
					Where(biller_account.ID(id)).
					Only(context.Background())
	if err != nil {
		log.Println(err)
	}

	return billerAccount, nil
}

func (repo billerAccountRepository) GetStores() ([]*ent.Biller_account, error) {
	billerAccounts, err := repo.clientDB.Biller_account.Query().
					All(context.Background())
	if err != nil {
		log.Panicln(err)
	}

	return billerAccounts, nil
}