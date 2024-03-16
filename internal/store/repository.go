package store

import "uroomBackend/internal/model"

type NewsRepository interface {
	CreateNews(n *model.News) error
	DeleteNews(id int) error
	FindById(id int) (*model.News, error)
	GetAllNews() ([]*model.News, error)
	UpdateNews(n *model.News) (string, error)
}
