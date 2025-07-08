package registration

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	Dbpool *pgxpool.Pool
}

func NewUserRepository(dbpool *pgxpool.Pool) *UserRepository {
	return &UserRepository{
		Dbpool: dbpool,
	}
}

func (r *UserRepository) addUser(form RegistrationForm) error {
	query := `INSERT INTO users (email, pass, username) VALUES (@email, @pass, @username)`
	args := pgx.NamedArgs{
		"email":    form.Email,
		"pass":     form.Password,
		"username": form.Name,
	}
	_, err := r.Dbpool.Exec(context.Background(), query, args)
	if err != nil {
		return fmt.Errorf("Невозможно сохранить пользователя: %v", err)
	}
	return nil
}
