package sqlstore

import (
	"database/sql"

	"github.com/vielendanke/restful-service/internal/app/model"
)

// PostRepository ...
type PostRepository struct {
	db *sql.DB
}

// FindAll ...
func (pr *PostRepository) FindAll() ([]model.Post, error) {
	row, err := pr.db.Query("SELECT id, name, content, user_id FROM posts")
	if err != nil {
		return nil, err
	}
	posts := []model.Post{}
	for row.Next() {
		post := &model.Post{}
		row.Scan(&post.ID, &post.Name, &post.Content, &post.UserID)
		posts = append(posts, *post)
	}
	return posts, nil
}

// FindByID ...
func (pr *PostRepository) FindByID(id string) *model.Post {
	return nil
}
