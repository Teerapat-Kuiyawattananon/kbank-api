package repository

import (
	"context"
	"kapi/ent"
	"kapi/ent/customer"
	"log"

	mCustomer "kapi/model"
)


type customerRepository struct {
	clientDB *ent.Client
}

func NewCustomerRepository(DB *ent.Client) customerRepository {
	return customerRepository{
		clientDB: DB,
	}
}

func (repo customerRepository) CreateCustomer(customer mCustomer.Customer) (*ent.Customer, error) {
	cus, err := repo.clientDB.Customer.Create().
					SetFirstName(customer.FirstName).
					SetLastName(customer.LastName).
					SetTitleName(customer.TitleName).
					SetMobileNumber(customer.MobileNumber).
					Save(context.Background())
	if err != nil {
		log.Println(err)
	}

	log.Printf("Created Customer id: %d success", cus.ID)
	return cus, nil
}

func (repo customerRepository) GetCustomerByID(id int) (*ent.Customer, error) {
	cus, err := repo.clientDB.Customer.Query().
					Where(customer.ID(id)).
					Only(context.Background())
	if err != nil {
		log.Println(err)
	}

	return cus, nil
}

func (repo customerRepository) GetCustomers() ([]*ent.Customer, error) {
	customers, err := repo.clientDB.Customer.Query().
						All(context.Background())
	if err != nil {
		log.Println(err)
	}

	return customers, nil
}