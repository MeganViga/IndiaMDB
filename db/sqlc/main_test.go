package db

import (
	"database/sql"
	"log"
	"os"
	"testing"
	_ "github.com/lib/pq"
)

var testQueries *Queries
const (
	sqlEngine= "postgres"
	sqlSource = "postgres://root:secret@localhost:5432/indiamdb?sslmode=disable"
)
func TestMain(m *testing.M){
	sqldb, err := sql.Open(sqlEngine,sqlSource)
	if err != nil{
		log.Fatal(err)
	}
	testQueries= New(sqldb)
	os.Exit((m.Run()))
}