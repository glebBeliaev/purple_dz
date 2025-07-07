package vacancy

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
)

type VacancyRepository struct {
	Dbpool       *pgxpool.Pool
	CustomLogger *zerolog.Logger
}

func NewVacancyRepository(dbpool *pgxpool.Pool, CustomLogger *zerolog.Logger) *VacancyRepository {
	return &VacancyRepository{
		Dbpool:       dbpool,
		CustomLogger: CustomLogger,
	}
}

func (r *VacancyRepository) CountAll() int {
	query := `SELECT COUNT(*) FROM vacancies`
	var count int
	r.Dbpool.QueryRow(context.Background(), query).Scan(&count)
	return count
}

func (r *VacancyRepository) GetAll(limit, offset int) ([]Vacancy, error) {
	query := `SELECT * from vacancies ORDER BY createdat LIMIT @limit OFFSET @offset`
	args := pgx.NamedArgs{
		"limit":  limit,
		"offset": offset,
	}
	rows, err := r.Dbpool.Query(context.Background(), query, args)
	if err != nil {
		return nil, err
	}
	vacancies, err := pgx.CollectRows(rows, pgx.RowToStructByName[Vacancy])
	if err != nil {
		return nil, err
	}
	return vacancies, nil
}

func (r *VacancyRepository) addVacancy(form VacancyCreateFrom) error {
	query := `INSERT INTO vacancies (email, role, company, type, salary, location, createdat) VALUES (@email, @role, @company, @type, @salary, @location, @createdat)`
	args := pgx.NamedArgs{
		"email":     form.Email,
		"role":      form.Role,
		"company":   form.Company,
		"type":      form.CompanyType,
		"salary":    form.Salary,
		"location":  form.Location,
		"createdat": time.Now(),
	}
	_, err := r.Dbpool.Exec(context.Background(), query, args)
	if err != nil {
		return fmt.Errorf("Невозможно создать вакансию: %v", err)
	}
	return nil
}
