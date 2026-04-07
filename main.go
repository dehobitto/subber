package main

import (
	"context"
	"subber/routes"

	"github.com/jackc/pgx/v5"
)

func main() {
	r := routes.SetupRouter()
	r.Run(":8080")
}

func connect() (*pgx.Conn, error) {
	// TODO:
	conn, err := pgx.Connect(context.Background(), "")

	if err != nil {
		return nil, err
	}
	return conn, nil
}
