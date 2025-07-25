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

func (r *NewsRepository) CreateNews(news News) error {
	query := `INSERT INTO news (title, content, author_email) VALUES ($1, $2, $3)`
	_, err := r.Dbpool.Exec(context.Background(), query, news.Title, news.Content, news.AuthorEmail)
	if err != nil {
		return err
	}
	return nil
}
