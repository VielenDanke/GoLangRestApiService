package sqlstore

import (
	"database/sql"

	"github.com/vielendanke/restful-service/internal/app/store"
)

// Store ...
type Store struct {
	db             *sql.DB
	userRepository *UserRepository
	postRepository *PostRepository
}

// NewStore ...
func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

// UserRepository ...
func (st *Store) UserRepository() store.UserRepository {
	if st.userRepository == nil {
		st.userRepository = &UserRepository{
			db: st.db,
		}
	}
	return st.userRepository
}

// PostRepository ...
func (st *Store) PostRepository() store.PostRepository {
	if st.postRepository == nil {
		st.postRepository = &PostRepository{
			db: st.db,
		}
	}
	return st.postRepository
}
