package handlers

import (
	"subber/infra/database"
)

type Handler struct {
	Repo *database.Repository
}

func NewHandler(repo *database.Repository) *Handler {
	return &Handler{
		Repo: repo,
	}
}
