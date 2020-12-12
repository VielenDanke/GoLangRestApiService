package store

// Store ...
type Store interface {
	UserRepository() UserRepository
	PostRepository() PostRepository
}
