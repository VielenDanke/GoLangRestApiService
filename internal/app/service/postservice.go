package service

import (
	"github.com/vielendanke/restful-service/internal/app/model"
	"github.com/vielendanke/restful-service/internal/app/sqlstore"
)

// PostService ...
type PostService struct {
	store *sqlstore.Store
}

// SavePost ...
func (ps *PostService) SavePost(post *model.Post) error {
	if err := post.Validate(); err != nil {
		return err
	}
	if err := ps.store.PostRepository().Save(post); err != nil {
		return err
	}
	return nil
}

// FindAllPosts ...
func (ps *PostService) FindAllPosts() ([]model.Post, error) {
	return ps.store.PostRepository().FindAll()
}

// DeletePost ...
func (ps *PostService) DeletePost(id string) error {
	return ps.store.PostRepository().Delete(id)
}

// FindAllPostsByUserID ...
func (ps *PostService) FindAllPostsByUserID(id string) ([]model.Post, error) {
	return ps.store.PostRepository().FindAllPostsByUserID(id)
}
