package database

import (
	"context"

	"github.com/glebbeliaev/purple_dz/config"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
)

func CreateDbPool(config *config.DataBaseConfig, logger *zerolog.Logger) *pgxpool.Pool {
	dbpool, err := pgxpool.New(context.Background(), config.Url)
	if err != nil {
		logger.Error().Msg("Не удалось подключится к базе данных")
		panic(err)
	}
	logger.Info().Msg("Подключение к базе данных установлено")

	return dbpool
}
