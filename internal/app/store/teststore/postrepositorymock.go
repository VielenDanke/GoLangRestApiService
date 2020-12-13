package teststore

import (
	"fmt"

	"github.com/vielendanke/restful-service/internal/app/model"
)

// PostRepository ...
type PostRepository struct {
	PostDB map[string]model.Post
}

// FindAll ...
func (pr *PostRepository) FindAll() ([]model.Post, error) {
	posts := make([]model.Post, len(pr.PostDB))
	for _, v := range pr.PostDB {
		posts = append(posts, v)
	}
	return posts, nil
}

// Find ...
func (pr *PostRepository) Find(id string) (*model.Post, error) {
	post, ok := pr.PostDB[id]
	if !ok {
		return nil, fmt.Errorf("Post not found")
	}
	return &post, nil
}

// Save ...
func (pr *PostRepository) Save(post *model.Post) error {
	pr.PostDB[post.ID] = *post
	return nil
}

// Delete ...
func (pr *PostRepository) Delete(id string) error {
	delete(pr.PostDB, id)
	return nil
}

// FindAllPostsByUserID ...
func (pr *PostRepository) FindAllPostsByUserID(id string) ([]model.Post, error) {
	posts := []model.Post{}
	for _, v := range pr.PostDB {
		if v.UserID == id {
			posts = append(posts, v)
		}
	}
	return posts, nil
}
