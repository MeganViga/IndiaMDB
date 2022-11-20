package main

import (
	"database/sql"
	"log"

	"github.com/MeganViga/IndiaMDB/api"
	db "github.com/MeganViga/IndiaMDB/db/sqlc"
	_ "github.com/lib/pq"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgres://root:secret@localhost:5432/indiamdb?sslmode=disable"
	sourceAddress = "0.0.0.0:8082"
)
func main(){
	conn, err := sql.Open(dbDriver,dbSource)
	if err != nil{
		log.Fatal(err)
	}

	store := db.NewStore(conn)
	server , err:= api.NewServer(store)

	server.Start(sourceAddress)
}