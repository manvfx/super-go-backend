package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/manvfx/super-go-backend/api/middlewares"
	"github.com/manvfx/super-go-backend/api/routers"
	validation "github.com/manvfx/super-go-backend/api/validations"
	"github.com/manvfx/super-go-backend/config"
	"github.com/manvfx/super-go-backend/pkg/logging"
)

var logger = logging.NewLogger(config.GetConfig())

func InitServer() {
	cfg := config.GetConfig()
	r := gin.New()

	// RegisterValidators()
	val, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		err := val.RegisterValidation("mobile", validation.IranianMobileNumberValidator, true)
		if err != nil {
			logger.Error(logging.Validation, logging.Startup, err.Error(), nil)
		}
		err = val.RegisterValidation("password", validation.PasswordValidator, true)
		if err != nil {
			logger.Error(logging.Validation, logging.Startup, err.Error(), nil)
		}
	}

	r.Use(gin.Logger(), gin.Recovery(), middlewares.LimitByRequest(), middlewares.AuthMiddleware())

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

// func RegisterValidators() {
// 	val, ok := binding.Validator.Engine().(*validator.Validate)
// 	if ok {
// 		err := val.RegisterValidation("mobile", validation.IranianMobileNumberValidator, true)
// 		if err != nil {
// 			logger.Error(logging.Validation, logging.Startup, err.Error(), nil)
// 		}
// 		err = val.RegisterValidation("password", validation.PasswordValidator, true)
// 		if err != nil {
// 			logger.Error(logging.Validation, logging.Startup, err.Error(), nil)
// 		}
// 	}
// }
