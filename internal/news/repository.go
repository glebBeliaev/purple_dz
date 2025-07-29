package news

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type NewsRepository struct {
	Dbpool *pgxpool.Pool
}

func NewNewsRepository(dbpool *pgxpool.Pool) *NewsRepository {
	return &NewsRepository{
		Dbpool: dbpool,
	}
}

func (r *NewsRepository) createNews(form NewsForm) error {
	query := `INSERT INTO news (title, preview, content, cover, category, author) VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := r.Dbpool.Exec(context.Background(), query, form.Title, form.Preview, form.Content, form.Cover, form.Category, form.Author)
	if err != nil {
		return err
	}
	return nil
}
