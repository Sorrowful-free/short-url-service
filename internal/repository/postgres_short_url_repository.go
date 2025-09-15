package repository

import (
	_ "github.com/jackc/pgx/v5/stdlib"

	"context"
	"database/sql"

	"github.com/Sorrowful-free/short-url-service/internal/model"
)

type PostgresShortURLRepository struct {
	db *sql.DB
}

func NewPostgresShortURLRepository(databaseDSN string) (ShortURLRepository, error) {

	db, err := sql.Open("pgx", databaseDSN)
	if err != nil {
		return nil, err
	}

	return &PostgresShortURLRepository{
		db: db,
	}, nil
}

func (r *PostgresShortURLRepository) Save(ctx context.Context, shortURL model.ShortURLDto) error {
	if ctx.Err() != nil {
		return ctx.Err()
	}
	_, err := r.db.ExecContext(ctx, "INSERT INTO short_urls (short_uid, original_url) VALUES ($1, $2)", shortURL.ShortUID, shortURL.OriginalURL)
	if err != nil {
		return err
	}
	return nil
}

func (r *PostgresShortURLRepository) ContainsUID(ctx context.Context, shortUID string) bool {
	if ctx.Err() != nil {
		return false
	}
	var count int
	err := r.db.QueryRowContext(ctx, "SELECT COUNT(*) FROM short_urls WHERE short_uid = $1", shortUID).Scan(&count)
	if err != nil {
		return false
	}
	return count > 0
}

func (r *PostgresShortURLRepository) GetByUID(ctx context.Context, shortUID string) (model.ShortURLDto, error) {
	if ctx.Err() != nil {
		return model.ShortURLDto{}, ctx.Err()
	}
	var originalURL string
	err := r.db.QueryRowContext(ctx, "SELECT original_url FROM short_urls WHERE short_uid = $1", shortUID).Scan(&originalURL)
	if err != nil {
		return model.ShortURLDto{}, err
	}
	return model.ShortURLDto{ShortUID: shortUID, OriginalURL: originalURL}, nil
}

func (r *PostgresShortURLRepository) Ping(ctx context.Context) error {
	if ctx.Err() != nil {
		return ctx.Err()
	}
	err := r.db.PingContext(ctx)
	if err != nil {
		return err
	}
	return nil
}
