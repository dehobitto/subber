package handlers

import (
	"subber/config"
	"subber/infra/database"
	"subber/workers"
)

type Handler struct {
	repo *database.Repository
	cfg  *config.Config
	jobs chan<- workers.NotificationJob
}

func NewHandler(repo *database.Repository, cfg *config.Config, jobs chan<- workers.NotificationJob) *Handler {
	return &Handler{
		repo: repo,
		cfg:  cfg,
		jobs: jobs,
	}
}
