package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HealthAdapter struct {
}

func NewHealthAdapter() *HealthAdapter {
	return &HealthAdapter{}
}

func (a *HealthAdapter) RegisterRoutes(rg *gin.RouterGroup) {
	health := rg.Group("/ping")
	{
		health.GET("/", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{"message": "pong"})
		})
	}
}
