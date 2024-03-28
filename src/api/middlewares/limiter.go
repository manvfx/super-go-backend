package middlewares

import (
	"net/http"

	"github.com/didip/tollbooth/v7"
	"github.com/gin-gonic/gin"
	"github.com/manvfx/super-go-backend/api/helper"
)

func LimitByRequest() gin.HandlerFunc {
	lmt := tollbooth.NewLimiter(1, nil)
	return func(c *gin.Context) {
		err := tollbooth.LimitByRequest(lmt, c.Writer, c.Request)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusTooManyRequests,
				helper.GenerateBaseResponseWithError(nil, false, helper.LimiterError, err))
			return
		} else {
			c.Next()
		}
	}
}
