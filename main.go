package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/nikit34/golang_backend/api"
	db "github.com/nikit34/golang_backend/db/sqlc"
	"github.com/nikit34/golang_backend/util"
)


func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server: ", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
}