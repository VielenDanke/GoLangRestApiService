package serviceimpl

import (
	"github.com/vielendanke/restful-service/internal/app/model"
	"github.com/vielendanke/restful-service/internal/app/store"
)

// PostServiceImpl ...
type PostServiceImpl struct {
	store store.Store
}

// SavePost ...
func (ps *PostServiceImpl) SavePost(post *model.Post) error {
	if err := post.Validate(); err != nil {
		return err
	}
	post.BeforeSaving()
	if err := ps.store.PostRepository().Save(post); err != nil {
		return err
	}
	return nil
}

// FindByID ...
func (ps *PostServiceImpl) FindByID(id string) (*model.Post, error) {
	return ps.store.PostRepository().Find(id)
}

// FindAllPosts ...
func (ps *PostServiceImpl) FindAllPosts() ([]model.Post, error) {
	return ps.store.PostRepository().FindAll()
}

// DeletePost ...
func (ps *PostServiceImpl) DeletePost(id string) error {
	return ps.store.PostRepository().Delete(id)
}

// FindAllPostsByUserID ...
func (ps *PostServiceImpl) FindAllPostsByUserID(id string) ([]model.Post, error) {
	return ps.store.PostRepository().FindAllPostsByUserID(id)
}
