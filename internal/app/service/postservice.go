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
	return nil
}

// FindAllPosts ...
func (ps *PostService) FindAllPosts() ([]model.Post, error) {
	return ps.store.PostRepository().FindAll()
}
