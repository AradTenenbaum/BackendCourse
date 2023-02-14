package db

import (
	"database/sql"
	"log"
	"testing"

	"github.com/AradTenenbaum/BackendCourse/util"
	_ "github.com/lib/pq"
)

var testQuery *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect db:", err)
	}

	testQuery = New(testDB)

	m.Run()
}