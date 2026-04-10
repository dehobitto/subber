package routes

import (
	"subber/handlers"
	"subber/infra/database"

	"github.com/gin-gonic/gin"
)

func SetupRouter(repo *database.Repository) *gin.Engine {
	r := gin.Default()
	r.SetTrustedProxies(nil)

	h := handlers.NewHandler(repo)

	api := r.Group("/api")
	{
		api.POST("/subscribe", h.Subscribe)
		api.GET("/confirm/:token", h.ConfirmByToken)
		api.GET("/unsubscribe/:token", h.UnsubscribeByToken)
		api.GET("/subscriptions/", h.GetSubscriptions)
	}

	return r
}
