package apiserver_test

import (
	"github.com/vielendanke/restful-service/internal/app/service"
	"github.com/vielendanke/restful-service/internal/app/service/testservice"
	"os"
	"testing"

	"github.com/vielendanke/restful-service/internal/app/apiserver"
)

var (
	testServer *apiserver.Server
	testService service.Service
	testConfig *apiserver.Config
	err        error
)

func TestMain(m *testing.M) {
	testConfig = apiserver.NewConfig()
	testConfig.TokenSecret = "test_token_secret"
	testConfig.TokenValidTime = 15
	testService = testservice.NewServiceMock()
	testServer, err = apiserver.NewServer(testService, testConfig)
	if err != nil {
		os.Exit(1)
	}
	os.Exit(m.Run())
}
