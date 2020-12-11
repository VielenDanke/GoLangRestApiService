package sqlstore_test

import (
	"database/sql"
	"fmt"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/vielendanke/restful-service/internal/app/sqlstore"
)

var (
	testDB    *sql.DB
	testStore *sqlstore.Store
)

func teardownTables(tables ...string) {
	for _, v := range tables {
		if v != "" {
			query := fmt.Sprintf("DELETE FROM %s", v)
			testDB.Exec(query)
		}
	}
}

func TestMain(m *testing.M) {
	var databaseURL string = os.Getenv("DATABASE_TEST_URL")
	var databaseName string = os.Getenv("DATABASE_TEST_NAME")
	if databaseURL == "" {
		databaseURL = "host=localhost dbname=restful_api_test sslmode=disable user=user password=userpassword"
	}
	if databaseName == "" {
		databaseName = "postgres"
	}
	db, err := sql.Open(databaseName, databaseURL)
	if err != nil {
		os.Exit(1)
	}
	if err := db.Ping(); err != nil {
		os.Exit(1)
	}
	testDB = db
	testStore = sqlstore.NewStore(db)
	os.Exit(m.Run())
}
