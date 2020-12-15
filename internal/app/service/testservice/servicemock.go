package testservice

import "github.com/vielendanke/restful-service/internal/app/service"

type ServiceMock struct {
	userService service.UserService
	postService service.PostService
}

func NewServiceMock() service.Service {
	return &ServiceMock{}
}

func (sm *ServiceMock) UserService() service.UserService {
	if sm.userService == nil {
		return &UserServiceMock{}
	}
	return sm.userService
}

func (sm *ServiceMock) PostService() service.PostService {
	if sm.postService == nil {
		return &PostServiceMock{}
	}
	return sm.postService
}
