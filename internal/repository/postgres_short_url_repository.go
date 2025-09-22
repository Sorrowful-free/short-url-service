package repository

import (
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
	_ "github.com/jackc/pgx/v5/stdlib"

	"context"
	"database/sql"

	"errors"

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

func (r *PostgresShortURLRepository) Save(ctx context.Context, userID string, shortURL model.ShortURLDto) error {
	if ctx.Err() != nil {
		return ctx.Err()
	}
	_, err := r.db.ExecContext(ctx, "INSERT INTO short_urls (short_uid, original_url, user_id) VALUES ($1, $2, $3)", shortURL.ShortUID, shortURL.OriginalURL, userID)
	var pgxErr *pgconn.PgError
	if err != nil {
		if errors.As(err, &pgxErr) && pgxErr.Code == pgerrcode.UniqueViolation {
			return NewOriginalURLConflictRepositoryError(shortURL.OriginalURL)
		}
		return err
	}
	return nil
}

func (r *PostgresShortURLRepository) SaveBatch(ctx context.Context, userID string, shortURLs []model.ShortURLDto) error {
	if ctx.Err() != nil {
		return ctx.Err()
	}

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	stmt, err := tx.PrepareContext(ctx, "INSERT INTO short_urls (short_uid, original_url, user_id) VALUES ($1, $2, $3)")
	if err != nil {
		tx.Rollback()
		return err
	}
	defer stmt.Close()

	for _, shortURL := range shortURLs {
		_, err := stmt.ExecContext(ctx, shortURL.ShortUID, shortURL.OriginalURL, userID)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	err = tx.Commit()
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
	row := r.db.QueryRowContext(ctx, "SELECT original_url FROM short_urls WHERE short_uid = $1", shortUID)
	if row.Err() != nil {
		return model.ShortURLDto{}, row.Err()
	}
	if err := row.Scan(&originalURL); err != nil {
		return model.ShortURLDto{}, err
	}
	return model.ShortURLDto{ShortUID: shortUID, OriginalURL: originalURL}, nil
}

func (r *PostgresShortURLRepository) GetByOriginalURL(ctx context.Context, originalURL string) (model.ShortURLDto, error) {
	if ctx.Err() != nil {
		return model.ShortURLDto{}, ctx.Err()
	}
	var shortUID string
	row := r.db.QueryRowContext(ctx, "SELECT short_uid FROM short_urls WHERE original_url = $1", originalURL)
	if row.Err() != nil {
		return model.ShortURLDto{}, row.Err()
	}
	if err := row.Scan(&shortUID); err != nil {
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

func (r *PostgresShortURLRepository) GetUserUrls(ctx context.Context, userID string) ([]model.ShortURLDto, error) {
	if ctx.Err() != nil {
		return nil, ctx.Err()
	}

	rows, err := r.db.QueryContext(ctx, "SELECT short_uid, original_url FROM short_urls WHERE user_id = $1", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if err := rows.Err(); err != nil {
		return nil, err
	}

	shortURLs := make([]model.ShortURLDto, 0)
	for rows.Next() {
		var shortUID string
		var originalURL string
		err = rows.Scan(&shortUID, &originalURL)
		if err != nil {
			return nil, err
		}

		if err := rows.Err(); err != nil {
			return nil, err
		}

		shortURLs = append(shortURLs, model.NewShortURLDto(shortUID, originalURL))
	}
	return shortURLs, nil
}
