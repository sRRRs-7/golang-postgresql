package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/sRRRs-7/golang-postgresql/api"
	db "github.com/sRRRs-7/golang-postgresql/db/sqlc"
	"github.com/sRRRs-7/golang-postgresql/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBdriver, config.DBsource)
	if err != nil {
		log.Fatal("cannot connect to database", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server", err)
	}
}
