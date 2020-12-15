package testservice

import "github.com/vielendanke/restful-service/internal/app/model"

type PostServiceMock struct {}

func (pm *PostServiceMock) SavePost(post *model.Post) error {
	return nil
}

// FindByID ...
func (pm *PostServiceMock) FindByID(id string) (*model.Post, error) {
	return model.MockPost(), nil
}

// FindAllPosts ...
func (pm *PostServiceMock) FindAllPosts() ([]model.Post, error) {
	return []model.Post{}, nil
}

// DeletePost ...
func (pm *PostServiceMock) DeletePost(id string) error {
	return nil
}

// FindAllPostsByUserID ...
func (pm *PostServiceMock) FindAllPostsByUserID(id string) ([]model.Post, error) {
	return []model.Post{}, nil
}