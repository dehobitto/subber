package routes

import (
	"subber/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.SetTrustedProxies(nil)

	api := r.Group("/api")
	{
		api.POST("/subscribe", handlers.Subscribe)
		api.GET("/confirm/", handlers.ConfirmByToken)
		api.GET("/unsubscribe/", handlers.UnsubscribeByToken)
		api.GET("/subscriptions/", handlers.GetSubscriptions)
	}

	return r
}
