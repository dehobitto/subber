package main

import (
	"log"
	"subber/infra/database"
	"subber/routes"
)

func main() {
	r := routes.SetupRouter()
	db, err := database.Connect()

	if err != nil {
		log.Printf("Database is not initilazed: %v", err)
		return
	}
	defer db.Close()

	err = database.Migrate(db, "infra/database/schemas.sql")
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	log.Println("Server is starting on :8080")

	if err = r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
