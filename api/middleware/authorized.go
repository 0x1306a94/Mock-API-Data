package middleware

import (
	"Mock-API-Data/constant"
	"Mock-API-Data/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthorizedMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.Request.Header.Get(constant.HTTPHeaderMockTokenKey)
		if tokenStr == "" {
			c.Writer.WriteHeader(http.StatusUnauthorized)
			c.Abort()
			return
		}
		user, err := util.ParseUserWithToken(tokenStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, util.GenerateErrorResponse(http.StatusBadRequest, err.Error()))
			c.Abort()
			return
		}
		c.Set("_user", user)
		c.Next()
	}
}
