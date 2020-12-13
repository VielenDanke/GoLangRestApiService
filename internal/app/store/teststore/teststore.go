package teststore

import (
	"github.com/vielendanke/restful-service/internal/app/model"
	"github.com/vielendanke/restful-service/internal/app/store"
)

// TestStore ...
type TestStore struct {
	TestUserDB map[string]model.User
	TestPostDB map[string]model.Post
	UserRepo   *UserRepository
	PostRepo   *PostRepository
}

// NewTestStore ...
func NewTestStore(testUserDB map[string]model.User, testPostDB map[string]model.Post) store.Store {
	return &TestStore{
		TestUserDB: testUserDB,
		TestPostDB: testPostDB,
	}
}

// UserRepository ...
func (ts *TestStore) UserRepository() store.UserRepository {
	if ts.UserRepo == nil {
		ts.UserRepo = &UserRepository{
			UserDB: ts.TestUserDB,
		}
	}
	return ts.UserRepo
}

// PostRepository ...
func (ts *TestStore) PostRepository() store.PostRepository {
	if ts.PostRepo == nil {
		ts.PostRepo = &PostRepository{
			PostDB: ts.TestPostDB,
		}
	}
	return ts.PostRepo
}
