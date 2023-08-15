package progresql

import (
	"context"
	"kapi/ent"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func InitDatabase() (*ent.Client, error) {
	ctx := context.Background()

	// init Database
	clientDB, err := ent.Open("postgres", os.Getenv("PSQL_DB_CONNECT"))
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
		return nil, err
	}

	// Run the auto migration tool.

	if err := clientDB.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
		return clientDB, err
	}

	return clientDB, nil
}

func MockUpTestEX01() {
	
	ctx := context.Background()

	client, err := InitDatabase()
	if (client.Store.Query().CountX(ctx) != 0) {
		return
	}
	if err != nil {
		log.Fatal(err)
	}

	//Create Store
	store, err := client.Store.Create().
						SetID(98499).
						SetAccountName("ThaiTee").
						SetServiceName("Abank").
						Save(ctx)
	if err != nil {
		log.Fatal(err)
	}

	//Create Customer
	customer, err := client.Customer.Create().
						SetID(300000025751).
						SetFirstName("Teerapat").
						SetLastName("Ku").
						SetMobileNumber("0888888888").
						Save(ctx)
	if err != nil {
		log.Fatal(err)
	}

	//Create billDetail
	billDetail, err := client.BillDetail.Create().
						SetID(1733084).
						SetCustomerID(customer.ID).
						SetTranAmount(120.00).
						SetStatus("waiting").
						Save(ctx)
	if err != nil {
		log.Fatal(err)
	}

	//Create bill
	bill, err := client.Bill.Create().
					SetBillerID(store.ID).
					SetRef1(customer.ID).
					SetRef2(billDetail.ID).
					Save(ctx)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Created Bill id:%d Success", bill.ID)
	
}

func MockUpTest() {
	ctx := context.Background()

	client, err := InitDatabase()
	if err != nil {
		log.Fatal(err)
	}

	//Create Store
	store, err := client.Store.Create().
						SetAccountName("ThaiTee").
						SetServiceName("Abank").
						Save(ctx)
	if err != nil {
		log.Fatal(err, store)
	}

	//Create Customer
	// customer, err := client.Customer.Create().
	// 					SetFirstName("Teerapat").
	// 					SetLastName("Ku").
	// 					SetMobileNumber("0888888888").
	// 					Save(ctx)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	//Create billDetail
	billDetail, err := client.BillDetail.Create().
						SetCustomerID(300000025751).
						SetTranAmount(120.00).
						SetStatus("waiting").
						Save(ctx)
	if err != nil {
		log.Fatal(err, billDetail)
	}

	//Create bill
	bill, err := client.Bill.Create().
					SetBillerID(98499).
					SetRef1(300000025751).
					SetRef2(1733084).
					Save(ctx)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Created Bill id:%d Success", bill.ID)
	
}


 