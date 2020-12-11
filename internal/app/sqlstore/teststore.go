package sqlstore_test

import (
	"database/sql"
	"os"

	"github.com/vielendanke/restful-service/internal/app/sqlstore"
)

var store testStore

type testStore struct {
	store *sqlstore.Store
}

func TestMain(m *testing.M) {
	if err := gotoenv.Load(); err != nil {

    }
	databaseTestUrl := os.Getenv("DATABASE_TEST_URL")
	databaseName := os.Getenv("DATABASE_TEST_NAME")

	db, err := sql.Open(databaseName, databaseTestUrl)
	if err != nil {
		
	}
	if err := db.Ping(); err != nil {
		
	}
	testStore = &testStore{
		store: sqlstore.NewStore(db)
	}
}
