package main

import (
	"subber/infra/db"
	"subber/routes"
)

func main() {
	r := routes.SetupRouter()
	db, err := db.Connect()

	r.Run(":8080")
}
