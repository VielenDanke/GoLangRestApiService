package service_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/vielendanke/restful-service/internal/app/model"
	"github.com/vielendanke/restful-service/internal/app/service"
	"github.com/vielendanke/restful-service/internal/app/store"
	"github.com/vielendanke/restful-service/internal/app/store/teststore"
)

var st store.Store
var testService *service.Service

func init() {
	st = teststore.NewTestStore()
	testService = service.NewService(st)
}

func TestPostService_SavePost(t *testing.T) {
	post := model.TestPost(t)
	post.UserID = uuid.New().String()
	err := testService.PostService().SavePost(post)
	assert.NoError(t, err)
	assert.NotEmpty(t, post.ID)
}

func TestPostService_FindByID(t *testing.T) {
	post := model.TestPost(t)
	post.UserID = uuid.New().String()
	_, err := testService.PostService().FindByID(post.ID)
	err = testService.PostService().SavePost(post)
	_, err = testService.PostService().FindByID(post.ID)
	assert.NoError(t, err)
}

func TestPostService_FindAll(t *testing.T) {
	_, err := testService.PostService().FindAllPosts()
	assert.NoError(t, err)
}

func TestPostService_DeletePost(t *testing.T) {
	post := model.TestPost(t)
	post.ID = uuid.New().String()
	testService.PostService().SavePost(post)
	err := testService.PostService().DeletePost(post.ID)
	assert.NoError(t, err)
}

func TestPostService_FindAllPostsByUserID(t *testing.T) {
	_, err := testService.PostService().FindAllPostsByUserID(uuid.New().String())
	assert.NoError(t, err)
}
