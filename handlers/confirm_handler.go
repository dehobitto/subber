package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) ConfirmByToken(c *gin.Context) {
	token := c.Param("token")

	err := h.Repo.ConfirmSubscriptionByToken(c.Request.Context(), token)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}
