package service

import "github.com/vielendanke/restful-service/internal/app/model"

type PostService interface {
	SavePost(post *model.Post) error
	FindByID(string) (*model.Post, error)
	FindAllPosts() ([]model.Post, error)
	DeletePost(string) error
	FindAllPostsByUserID(string) ([]model.Post, error)
}
