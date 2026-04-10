package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) UnsubscribeByToken(c *gin.Context) {
	token := c.Param("token")

	err := h.Repo.Unsubscribe(c.Request.Context(), token)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Token not found."})
		return
	}

	c.Status(http.StatusOK)
}
