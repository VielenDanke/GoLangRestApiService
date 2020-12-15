package serviceimpl_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vielendanke/restful-service/internal/app/model"
)

func TestUserService_SaveUser(t *testing.T) {
	defer teardownTestDB()
	user := model.TestUser(t)
	err := testService.UserService().SaveUser(user)
	assert.NoError(t, err)
	assert.NotEmpty(t, user.EncryptedPassword)
	assert.Empty(t, user.Password)
}

func TestUserService_FindByUsername(t *testing.T) {
	defer teardownTestDB()
	user := model.TestUser(t)
	u, err := testService.UserService().FindByUsername(user.Username)
	assert.Nil(t, u)
	assert.Error(t, err)
	testService.UserService().SaveUser(user)
	u, err = testService.UserService().FindByUsername(user.Username)
	assert.NotNil(t, u)
	assert.NoError(t, err)
}

func TestUserService_Login(t *testing.T) {
	defer teardownTestDB()
	user := model.TestUser(t)
	userPassword := user.Password
	u, err := testService.UserService().Login(user.Username, user.Password)
	assert.Nil(t, u)
	assert.Error(t, err)
	testService.UserService().SaveUser(user)
	u, err = testService.UserService().Login(user.Username, userPassword)
	assert.NotNil(t, u)
	assert.NoError(t, err)
}

func TestUserService_FindAllUsers(t *testing.T) {
	defer teardownTestDB()
	_, err := testService.UserService().FindAllUsers()
	assert.NoError(t, err)
}
