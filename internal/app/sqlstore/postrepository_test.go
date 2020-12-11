package sqlstore_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/vielendanke/restful-service/internal/app/model"
)

func TestPostRepository_SavePost(t *testing.T) {
	defer teardownTables("posts")
	post := model.TestPost(t)
	post.ID = uuid.New().String()
	post.UserID = uuid.New().String()
	err := testStore.PostRepository().Save(post)
	assert.NoError(t, err)
}

func TestPostRepository_DeletePost(t *testing.T) {
	defer teardownTables("posts")
	post := model.TestPost(t)
	post.ID = uuid.New().String()
	post.UserID = uuid.New().String()
	err := testStore.PostRepository().Save(post)
	assert.NoError(t, err)
	err = testStore.PostRepository().Delete(post.ID)
	assert.NoError(t, err)
	_, err = testStore.PostRepository().Find(post.ID)
	assert.Error(t, err)
}

func TestRepository_FindAllPostsByUserID(t *testing.T) {
	defer teardownTables("posts")
	post := model.TestPost(t)
	post.ID = uuid.New().String()
	post.UserID = uuid.New().String()
	posts, err := testStore.PostRepository().FindAllPostsByUserID(post.UserID)
	assert.Equal(t, 0, len(posts))
	assert.NoError(t, err)
	testStore.PostRepository().Save(post)
	posts, err = testStore.PostRepository().FindAllPostsByUserID(post.UserID)
	assert.NoError(t, err)
	assert.NotEqual(t, 0, len(posts))
}

func TestRepository_FindAll(t *testing.T) {
	defer teardownTables("posts")
	posts, err := testStore.PostRepository().FindAll()
	assert.NoError(t, err)
	assert.Equal(t, 0, len(posts))
	post := model.TestPost(t)
	post.ID = uuid.New().String()
	post.UserID = uuid.New().String()
	err = testStore.PostRepository().Save(post)
	assert.NoError(t, err)
	posts, err = testStore.PostRepository().FindAll()
	assert.NoError(t, err)
	assert.NotEqual(t, 0, len(posts))
}
