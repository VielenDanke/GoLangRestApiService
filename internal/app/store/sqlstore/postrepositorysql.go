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

// Find ...
func (pr *PostRepository) Find(id string) (*model.Post, error) {
	post := &model.Post{}
	if err := pr.db.QueryRow(
		"SELECT id, name, content, user_id FROM posts WHERE id=$1",
		id,
	).Scan(&post.ID, &post.Name, &post.Content, &post.UserID); err != nil {
		return nil, err
	}
	return post, nil
}

// Save ...
func (pr *PostRepository) Save(post *model.Post) error {
	row := pr.db.QueryRow(
		"INSERT INTO posts (id, name, content, user_id) VALUES ($1, $2, $3, $4)",
		&post.ID,
		&post.Name,
		&post.Content,
		&post.UserID,
	)
	if row.Err() != nil {
		return row.Err()
	}
	return nil
}

// Delete ...
func (pr *PostRepository) Delete(id string) error {
	_, err := pr.db.Exec("DELETE FROM posts WHERE id=$1", id)
	if err != nil {
		return err
	}
	return nil
}

// FindAllPostsByUserID ...
func (pr *PostRepository) FindAllPostsByUserID(id string) ([]model.Post, error) {
	rows, err := pr.db.Query("SELECT id, name, content FROM posts WHERE user_id=$1", id)
	if err != nil {
		return nil, err
	}
	posts := []model.Post{}
	for rows.Next() {
		post := &model.Post{}
		rows.Scan(&post.ID, &post.Name, &post.Content)
		posts = append(posts, *post)
	}
	return posts, nil
}
