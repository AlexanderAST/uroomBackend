package sqlstore

import (
	"database/sql"
	"errors"
	"fmt"
	"uroomBackend/internal/model"
	"uroomBackend/internal/store"
)

type NewsRepository struct {
	store *Store
}

func (r *NewsRepository) CreateNews(n *model.News) error {
	return r.store.db.QueryRow(`INSERT INTO news(date,name,small_description, full_description,image_path) VALUES ($1,$2,$3,$4,$5) RETURNING id`, n.Date, n.Name, n.SmallDescription, n.FullDescription, n.ImagePath).Scan(&n.ID)
}

func (r *NewsRepository) DeleteNews(id int) error {
	n := &model.News{}

	if err := r.store.db.QueryRow("SELECT * FROM news WHERE id = $1", id).Scan(&n.ID, &n.Date, &n.Name, &n.SmallDescription, &n.FullDescription, &n.ImagePath); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return err
		}
		return err
	}

	if err := r.store.db.QueryRow("DELETE FROM  news WHERE id = $1", id); err != nil {
		return err.Err()
	}

	return nil
}

func (r *NewsRepository) FindById(id int) (*model.News, error) {
	n := &model.News{}

	if err := r.store.db.QueryRow("SELECT * FROM news WHERE id = $1", id).Scan(&n.ID, &n.Date, &n.Name, &n.SmallDescription, &n.FullDescription, &n.ImagePath); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}

	return n, nil
}

func (r *NewsRepository) GetAllNews() ([]*model.News, error) {
	rows, err := r.store.db.Query("SELECT * FROM news")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	pdk := make([]*model.News, 0)

	for rows.Next() {
		pd := new(model.News)
		err := rows.Scan(&pd.ID, &pd.Date, &pd.Name, &pd.SmallDescription, &pd.FullDescription, &pd.ImagePath)
		if err != nil {
			return nil, err
		}

		pdk = append(pdk, pd)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return pdk, err
}

func (r *NewsRepository) UpdateNews(n *model.News) (string, error) {

	if err := r.store.db.QueryRow("UPDATE news SET date = $2, name = $3, small_description = $4, full_description = $5,image_path = $6 WHERE id =$1", n.ID, n.Date, n.Name, n.SmallDescription, n.FullDescription, n.ImagePath).Scan(&n.ID, &n.Date, &n.Name, &n.SmallDescription, &n.FullDescription, &n.ImagePath); err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return "ok", nil
		case err == nil:
			fmt.Println(n.ID)
		default:
			return "ok", err
		}
	}

	return "success", nil
}
