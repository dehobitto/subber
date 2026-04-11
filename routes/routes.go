package routes

import (
	"subber/config"
	"subber/handlers"
	"subber/infra/database"
	"subber/workers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(repo *database.Repository, cfg *config.Config, jobs chan<- workers.NotificationJob) *gin.Engine {
	r := gin.Default()
	r.SetTrustedProxies(nil)

	h := handlers.NewHandler(repo, cfg, jobs)

	api := r.Group("/api")
	{
		api.POST("/subscribe", h.Subscribe)
		api.GET("/confirm/:token", h.ConfirmByToken)
		api.GET("/unsubscribe/:token", h.UnsubscribeByToken)
		api.GET("/subscriptions/", h.GetSubscriptions)
	}

	return r
}
