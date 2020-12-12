package store

import "github.com/vielendanke/restful-service/internal/app/model"

// PostRepository ...
type PostRepository interface {
	FindAll() ([]model.Post, error)
	Find(string) (*model.Post, error)
	Save(*model.Post) error
	Delete(string) error
	FindAllPostsByUserID(string) ([]model.Post, error)
}
