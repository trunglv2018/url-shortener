package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/trunglen/g/x/rest"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				var errResponse = map[string]interface{}{
					"error":  err.(error).Error(),
					"status": "error",
				}
				if httpError, ok := err.(rest.IHttpError); ok {
					c.AbortWithStatusJSON(httpError.StatusCode(), errResponse)
				} else {
					c.AbortWithStatusJSON(500, errResponse)
				}
			}
		}()
		c.Next()
	}
}
