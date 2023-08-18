package progresql

import (
	"context"
	"kapi/ent"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func InitDatabase() (*ent.Client, error) {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}

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

func CloseDatabase(clientDB *ent.Client) {
	clientDB.Close()
}

func MockUpTestEX01() {

	ctx := context.Background()

	client, err := InitDatabase()
	if client.Biller_account.Query().CountX(ctx) != 0 {
		return
	}
	if err != nil {
		log.Fatal(err)
	}

	//Create Store.
	store, err := client.Biller_account.Create().
		SetID(98499).
		SetName("ThaiTee").
		SetServiceName("Abank").
		Save(ctx)
	if err != nil {
		log.Fatal(err)
	}

	//Create Customer.
	customer, err := client.Customer.Create().
		SetID(300000025751).
		SetFirstName("Teerapat").
		SetLastName("Ku").
		SetMobileNumber("0888888888").
		Save(ctx)
	if err != nil {
		log.Fatal(err)
	}

	//Create bill.
	bill, err := client.Bill.Create().
		SetBillerID(store.ID).
		SetReference1(customer.ID).
		SetReference2(1733084).
		SetTranAmount(120.00).
		SetStatus("waiting").
		Save(ctx)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Created Bill id:%d Success", bill.ID)

}
