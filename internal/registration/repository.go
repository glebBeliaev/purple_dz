package registration

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
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

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(form.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("не удалось захешировать пароль: %v", err)
	}

	query := `INSERT INTO users (email, pass, username) VALUES (@email, @pass, @username)`
	args := pgx.NamedArgs{
		"email":    form.Email,
		"pass":     string(hashedPassword),
		"username": form.Name,
	}

	_, err = r.Dbpool.Exec(context.Background(), query, args)
	if err != nil {
		// Проверка на ошибку уникальности email
		var pgErr *pgconn.PgError
		if ok := errors.As(err, &pgErr); ok && pgErr.Code == "23505" {
			return fmt.Errorf("email уже используется")
		}
		return fmt.Errorf("не удалось сохранить пользователя: %v", err)
	}
	return nil
}

func (r *UserRepository) loginUser(email, password string) (string, error) {
	var hashedPassword, username string

	query := `SELECT pass, username FROM users WHERE email=$1`
	err := r.Dbpool.QueryRow(context.Background(), query, email).Scan(&hashedPassword, &username)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return "", fmt.Errorf("неверный email или пароль")
		}
		return "", fmt.Errorf("ошибка при поиске пользователя: %v", err)
	}

	// Сравнение хешированного пароля
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return "", fmt.Errorf("неверный email или пароль")
	}

	return username, nil
}
