package teststore

import (
	"github.com/vielendanke/restful-service/internal/app/model"
	"github.com/vielendanke/restful-service/internal/app/store"
)

// TestStore ...
type TestStore struct {
	userRepository *UserRepository
	postRepository *PostRepository
}

// NewTestStore ...
func NewTestStore() store.Store {
	return &TestStore{}
}

// UserRepository ...
func (ts *TestStore) UserRepository() store.UserRepository {
	if ts.userRepository == nil {
		ts.userRepository = &UserRepository{
			userDB: make(map[string]*model.User),
		}
	}
	return ts.userRepository
}

// PostRepository ...
func (ts *TestStore) PostRepository() store.PostRepository {
	if ts.postRepository == nil {
		ts.postRepository = &PostRepository{
			postDB: make(map[string]*model.Post),
		}
	}
	return ts.postRepository
}
