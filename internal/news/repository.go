package news

import (
	"context"

	"github.com/jackc/pgx/v5"
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

func (r *NewsRepository) CountAll() int {
	query := `SELECT COUNT(*) FROM news`
	var count int
	r.Dbpool.QueryRow(context.Background(), query).Scan(&count)
	return count
}

func (r *NewsRepository) createNews(form NewsForm) error {
	query := `INSERT INTO news (title, preview, content, cover, category, author) VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := r.Dbpool.Exec(context.Background(), query, form.Title, form.Preview, form.Content, form.Cover, form.Category, form.Author)
	if err != nil {
		return err
	}
	return nil
}

func (r *NewsRepository) GetAll(limit, offset int) ([]News, error) {
	query := `SELECT * FROM news ORDER BY datacreate LIMIT @limit OFFSET @offset`
	args := pgx.NamedArgs{
		"limit":  limit,
		"offset": offset,
	}
	rows, err := r.Dbpool.Query(context.Background(), query, args)
	if err != nil {
		return nil, err
	}
	news, err := pgx.CollectRows(rows, pgx.RowToStructByName[News])
	if err != nil {
		return nil, err
	}
	return news, nil
}
