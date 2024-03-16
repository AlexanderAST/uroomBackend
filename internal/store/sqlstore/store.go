package sqlstore

import (
	"database/sql"
	"uroomBackend/internal/store"
)

type Store struct {
	db             *sql.DB
	newsRepository *NewsRepository
}

func New(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) News() store.NewsRepository {
	if s.newsRepository != nil {
		return s.newsRepository
	}
	s.newsRepository = &NewsRepository{store: s}

	return s.newsRepository
}
