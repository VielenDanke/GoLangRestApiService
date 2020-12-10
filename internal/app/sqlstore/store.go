package sqlstore

import (
	"database/sql"
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
func (st *Store) UserRepository() *UserRepository {
	if st.userRepository == nil {
		st.userRepository = &UserRepository{
			db: st.db,
		}
	}
	return st.userRepository
}

// PostRepository ...
func (st *Store) PostRepository() *PostRepository {
	if st.postRepository == nil {
		st.postRepository = &PostRepository{
			db: st.db,
		}
	}
	return st.postRepository
}
