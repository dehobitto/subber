package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Connect() (*pgxpool.Pool, error) {
	dsn := getDSN()

	pool, err := pgxpool.New(context.Background(), dsn)

	if err != nil {
		return nil, fmt.Errorf("unable to create connection pool: %v", err)
	}

	if err := pool.Ping(context.Background()); err != nil {
		return nil, fmt.Errorf("unable to ping database: %v", err)
	}

	return pool, nil
}

func Migrate(pool *pgxpool.Pool, filePath string) error {
	schema, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("could not read schema file: %v", err)

	}

	_, err = pool.Exec(context.Background(), string(schema))
	if err != nil {
		return fmt.Errorf("could not execute schema: %v", err)
	}

	fmt.Println("Migrations applied successfully!")

	return nil
}

// Gets a Data Source Name for PostgreSQL, using env
func getDSN() string {
	host := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	pswd := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, pswd, dbname)
}
