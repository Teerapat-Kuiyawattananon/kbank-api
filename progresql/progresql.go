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
 