package service

import (
	"github.com/vielendanke/restful-service/internal/app/store"
)

// Service ...
type Service struct {
	store       store.Store
	userService *UserService
	postService *PostService
}

// NewService ...
func NewService(store store.Store) *Service {
	return &Service{
		store: store,
	}
}

// UserService ...
func (us *Service) UserService() *UserService {
	if us.userService == nil {
		us.userService = &UserService{
			store: us.store,
		}
	}
	return us.userService
}

// PostService ...
func (us *Service) PostService() *PostService {
	if us.postService == nil {
		us.postService = &PostService{
			store: us.store,
		}
	}
	return us.postService
}
