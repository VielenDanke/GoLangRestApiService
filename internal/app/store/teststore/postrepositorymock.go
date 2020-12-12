package teststore

import "github.com/vielendanke/restful-service/internal/app/model"

// PostRepository ...
type PostRepository struct {
	postDB map[string]*model.Post
}

// FindAll ...
func (pr *PostRepository) FindAll() ([]model.Post, error) {
	return nil, nil
}

// Find ...
func (pr *PostRepository) Find(string) (*model.Post, error) {
	return nil, nil
}

// Save ...
func (pr *PostRepository) Save(*model.Post) error {
	return nil
}

// Delete ...
func (pr *PostRepository) Delete(string) error {
	return nil
}

// FindAllPostsByUserID ...
func (pr *PostRepository) FindAllPostsByUserID(string) ([]model.Post, error) {
	return nil, nil
}
