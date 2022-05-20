package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/sRRRs-7/golang-postgresql/api"
	db "github.com/sRRRs-7/golang-postgresql/db/sqlc"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://root:secret@localhost:5432/bank_app?sslmode=disable"
	serverAddress = "127.0.0.1:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to database", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server", err)
	}
}
