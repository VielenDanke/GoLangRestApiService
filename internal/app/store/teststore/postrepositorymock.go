package teststore

import (
	"fmt"

	"github.com/vielendanke/restful-service/internal/app/model"
)

// PostRepository ...
type PostRepository struct {
	postDB map[string]model.Post
}

// FindAll ...
func (pr *PostRepository) FindAll() ([]model.Post, error) {
	posts := make([]model.Post, len(pr.postDB))
	for _, v := range pr.postDB {
		posts = append(posts, v)
	}
	return posts, nil
}

// Find ...
func (pr *PostRepository) Find(id string) (*model.Post, error) {
	post, ok := pr.postDB[id]
	if !ok {
		return nil, fmt.Errorf("Post not found")
	}
	return &post, nil
}

// Save ...
func (pr *PostRepository) Save(post *model.Post) error {
	pr.postDB[post.ID] = *post
	return nil
}

// Delete ...
func (pr *PostRepository) Delete(id string) error {
	delete(pr.postDB, id)
	return nil
}

// FindAllPostsByUserID ...
func (pr *PostRepository) FindAllPostsByUserID(id string) ([]model.Post, error) {
	posts := []model.Post{}
	for _, v := range pr.postDB {
		if v.UserID == id {
			posts = append(posts, v)
		}
	}
	return posts, nil
}
