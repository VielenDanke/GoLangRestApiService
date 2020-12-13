package teststore

import (
	"github.com/vielendanke/restful-service/internal/app/model"
	"github.com/vielendanke/restful-service/internal/app/store"
)

// TestStore ...
type TestStore struct {
	UserRepo *UserRepository
	PostRepo *PostRepository
}

// NewTestStore ...
func NewTestStore() store.Store {
	return &TestStore{}
}

// UserRepository ...
func (ts *TestStore) UserRepository() store.UserRepository {
	if ts.UserRepo == nil {
		ts.UserRepo = &UserRepository{
			UserDB: make(map[string]model.User),
		}
	}
	return ts.UserRepo
}

// PostRepository ...
func (ts *TestStore) PostRepository() store.PostRepository {
	if ts.PostRepo == nil {
		ts.PostRepo = &PostRepository{
			PostDB: make(map[string]model.Post),
		}
	}
	return ts.PostRepo
}
