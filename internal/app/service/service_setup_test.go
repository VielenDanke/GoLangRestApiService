package service_test

import (
	"os"
	"testing"

	"github.com/vielendanke/restful-service/internal/app/model"
	"github.com/vielendanke/restful-service/internal/app/service"
	"github.com/vielendanke/restful-service/internal/app/store"
	"github.com/vielendanke/restful-service/internal/app/store/teststore"
)

var st store.Store
var testService *service.Service
var testPostDB map[string]model.Post
var testUserDB map[string]model.User

func teardownTestDB() {
	if testPostDB != nil && len(testPostDB) > 0 {
		for k := range testPostDB {
			delete(testPostDB, k)
		}
	}
	if testUserDB != nil && len(testUserDB) > 0 {
		for k := range testUserDB {
			delete(testUserDB, k)
		}
	}
}

func TestMain(m *testing.M) {
	testUserDB = make(map[string]model.User)
	testPostDB = make(map[string]model.Post)
	st = teststore.NewTestStore(testUserDB, testPostDB)
	testService = service.NewService(st)
	os.Exit(m.Run())
}
