package main

import (
	"log"
	"subber/infra/database"
	"subber/routes"
)

func main() {
	pool, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	err = database.Migrate(pool, "infra/database/schemas.sql")
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	repo := database.NewRepository(pool)

	r := routes.SetupRouter(repo)

	if err = r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
