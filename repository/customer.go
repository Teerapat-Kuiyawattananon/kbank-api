package repository

import (
	"context"
	"kapi/ent"
	"kapi/ent/customer"
	"log"
	"time"

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
					SetCreatedAt(func () time.Time {
						strTime := time.Now().Add(time.Hour * 7).Format(time.RFC3339)
						t, _ := time.Parse(time.RFC3339, strTime)
						// loc, _ := time.LoadLocation("Asia/Bangkok")
						return t
					}()).
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