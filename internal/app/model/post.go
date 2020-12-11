package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/google/uuid"
)

// Post ...
type Post struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Content string `json:"content"`
	UserID  string `json:"userID"`
}

// BeforeSaving ...
func (p *Post) BeforeSaving() {
	p.ID = uuid.New().String()
}

// Validate ...
func (p *Post) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.Name, validation.Required, validation.Length(3, 128)),
		validation.Field(&p.Content, validation.Required, validation.Length(20, 1024)),
		validation.Field(&p.UserID, validation.Required, validation.NotNil),
	)
}
