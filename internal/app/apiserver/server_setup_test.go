package apiserver_test

import (
	"os"
	"testing"

	"github.com/vielendanke/restful-service/internal/app/apiserver"
	"github.com/vielendanke/restful-service/internal/app/model"
	"github.com/vielendanke/restful-service/internal/app/store"
	"github.com/vielendanke/restful-service/internal/app/store/teststore"
)

var (
	testServer *apiserver.Server
	testStore  store.Store
	testUserDB map[string]model.User
	testPostDB map[string]model.Post
	testConfig *apiserver.Config
	err        error
)

func TestMain(m *testing.M) {
	testUserDB = make(map[string]model.User)
	testPostDB = make(map[string]model.Post)
	testConfig = apiserver.NewConfig()
	testConfig.TokenSecret = "test_token_secret"
	testConfig.TokenValidTime = 15
	testStore = teststore.NewTestStore(testUserDB, testPostDB)
	testServer, err = apiserver.NewServer(testStore, testConfig)
	if err != nil {
		os.Exit(1)
	}
	os.Exit(m.Run())
}
