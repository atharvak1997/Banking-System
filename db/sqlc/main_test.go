package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:99220Avk97*@localhost:5432/simple_bank?sslmode=disable"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatal("Cannot connect to db:", err)
	} else {
		testQueries = New(conn)
		log.Println("Connection to db successful!")
		os.Exit(m.Run())
	}

}
