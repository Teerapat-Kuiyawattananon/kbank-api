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

func (repo billerAccountRepository) CreateBillerAccount(input model.BillAccount) (*ent.Biller_account, error) {
	billerAccount, err := repo.clientDB.Biller_account.Create().
					SetName(input.Name).
					SetServiceName(input.ServiceName).
					Save(context.Background())
	if err != nil {
		log.Println(err)
	}

	log.Printf("Created Biller id: %d success", billerAccount.ID)
	return billerAccount, nil
}

func (repo billerAccountRepository) GetBillerByID(id int) (*ent.Biller_account, error) {
	billerAccount, err := repo.clientDB.Biller_account.Query().
					Where(biller_account.ID(id)).
					Only(context.Background())
	if err != nil {
		log.Println(err)
	}

	return billerAccount, nil
}

func (repo billerAccountRepository) GetBillers() ([]*ent.Biller_account, error) {
	billerAccounts, err := repo.clientDB.Biller_account.Query().
					All(context.Background())
	if err != nil {
		log.Panicln(err)
	}

	return billerAccounts, nil
}
