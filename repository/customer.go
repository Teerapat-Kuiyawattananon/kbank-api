package repository

import (
	"context"
	"kapi/ent"
	"log"
	"time"

	mCustomer "kapi/model/customer"
)


type customerRepository struct {
	clientDB *ent.Client
}

func NewCustomerRepository(DB *ent.Client) customerRepository {
	return customerRepository{
		clientDB: DB,
	}
}

func (repo customerRepository) CreateCustomer(customer mCustomer.Customer) (error) {
	cus, err := repo.clientDB.Customer.Create().
					SetFirstName(customer.FirstName).
					SetLastName(customer.LastName).
					SetTitleName(customer.TitleName).
					SetMobileNumber(customer.MobileNumber).
					SetCreatedAt(time.Now().UTC().Add(time.Hour * 7)).
					Save(context.Background())
	if err != nil {
		log.Fatal(err)
		return err
	}

	log.Printf("Created Customer id: %d success", cus.ID)
	return nil
}