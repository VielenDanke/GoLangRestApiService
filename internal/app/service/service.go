package service

type Service interface {
	UserService() UserService
	PostService() PostService
}
