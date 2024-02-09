package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthHandler struct {
}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) Health(c *gin.Context) {
	c.JSON(http.StatusOK, "Working route...!")
	return
}

func (h *HealthHandler) HealthPost(c *gin.Context) {
	c.JSON(http.StatusOK, "Working created...!")
	return
}

func (h *HealthHandler) HealthById(c *gin.Context) {
	id := c.Params.ByName("id")
	c.JSON(http.StatusOK, fmt.Sprintf("Working show by id: %s!", id))
	return
}
