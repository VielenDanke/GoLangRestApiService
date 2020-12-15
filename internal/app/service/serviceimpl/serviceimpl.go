package serviceimpl

import (
	"github.com/vielendanke/restful-service/internal/app/service"
	"github.com/vielendanke/restful-service/internal/app/store"
)

// ServiceImpl ...
type ServiceImpl struct {
	store       store.Store
	userService service.UserService
	postService service.PostService
}

// NewService ...
func NewService(store store.Store) service.Service {
	return &ServiceImpl{
		store: store,
	}
}

// UserServiceImpl ...
func (us *ServiceImpl) UserService() service.UserService {
	if us.userService == nil {
		us.userService = &UserServiceImpl{
			store: us.store,
		}
	}
	return us.userService
}

// PostServiceImpl ...
func (us *ServiceImpl) PostService() service.PostService {
	if us.postService == nil {
		us.postService = &PostServiceImpl{
			store: us.store,
		}
	}
	return us.postService
}
