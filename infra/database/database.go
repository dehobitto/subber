package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"subber/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Connect() (*pgxpool.Pool, error) {
	dsn := getDSN()

	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		return nil, err
	}

	if err := pool.Ping(context.Background()); err != nil {
		return nil, err
	}

	log.Println("Database connection established")
	return pool, nil
}

func Migrate(pool *pgxpool.Pool, filePath string) error {
	schema, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	_, err = pool.Exec(context.Background(), string(schema))
	if err != nil {
		return err
	}

	log.Println("Migrations applied successfully!")
	return nil
}

type Repository struct {
	pool *pgxpool.Pool
}

func NewRepository(pool *pgxpool.Pool) *Repository {
	return &Repository{pool: pool}
}

func (r *Repository) SaveSubscription(ctx context.Context, sub models.Subscription) error {
	query := `
        INSERT INTO subscriptions (email, repo, confirmed, last_seen_tag, token)
        VALUES ($1, $2, $3, $4, $5)
        ON CONFLICT (email, repo) DO UPDATE 
        SET last_seen_tag = EXCLUDED.last_seen_tag;
    `

	_, err := r.pool.Exec(ctx, query, sub.Email, sub.Repo, sub.Confirmed, sub.LastSeenTag, sub.Token)

	if err != nil {
		log.Printf("Failed to save subscription for %s: %v", sub.Email, err)
		return err
	}

	log.Printf("Subscription saved for %s on %s", sub.Email, sub.Repo)
	return nil
}

func getDSN() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
}

func (r *Repository) ConfirmSubscriptionByToken(ctx context.Context, token string) error {
	query := `
	UPDATE subscriptions
	SET confirmed = true
	WHERE token = $1
	`

	_, err := r.pool.Exec(ctx, query, token)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) Unsubscribe(ctx context.Context, token string) error {
	query := `
	DELETE FROM subscriptions
	WHERE token = $1
	`

	result, err := r.pool.Exec(ctx, query, token)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("Token not found.")
	}

	return nil
}

func (r *Repository) GetSubscriptions(ctx context.Context, email string) ([]models.Subscription, error) {
	query := `
        SELECT email, repo, confirmed, last_seen_tag 
        FROM subscriptions
        WHERE email = $1 AND confirmed = true
    `

	rows, err := r.pool.Query(ctx, query, email)
	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}
	defer rows.Close()

	var subs []models.Subscription
	for rows.Next() {
		var s models.Subscription

		err := rows.Scan(
			&s.Email,
			&s.Repo,
			&s.Confirmed,
			&s.LastSeenTag,
		)
		if err != nil {
			return nil, fmt.Errorf("scan failed: %w", err)
		}

		subs = append(subs, s)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return subs, nil
}
