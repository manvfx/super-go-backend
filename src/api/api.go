package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/manvfx/super-go-backend/api/routers"
	"github.com/manvfx/super-go-backend/config"
)

func InitServer() {
	cfg := config.GetConfig()
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	api := r.Group("/api")

	v1 := api.Group("/v1")
	{
		health := v1.Group("/health")
		routers.Health(health)

		test := v1.Group("/test")
		routers.TestRouter(test)
	}

	r.Run(fmt.Sprintf(":%s", cfg.Server.ExternalPort))
}
