package handler

import (
	"kapi/ent"
	db "kapi/progresql"
	"log"
)

var (
	clientDB *ent.Client
	err error
)

func init() {
	log.Println("InitialHandler")
	clientDB, err = db.InitDatabase()
	if err != nil {
		log.Fatal(err)
	}
}

func CloseHandler() {
	db.CloseDatabase(clientDB)
}
