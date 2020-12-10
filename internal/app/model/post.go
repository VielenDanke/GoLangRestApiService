package model

import "github.com/google/uuid"

// Post ...
type Post struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Content string `json:"content"`
	UserID  string `json:"userID"`
}

// Posts ...
type Posts []Post

// BeforeCreate ...
func (p *Post) BeforeCreate() {
	p.ID = uuid.New().String()
}
