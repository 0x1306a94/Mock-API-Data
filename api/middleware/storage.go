package middleware

import (
	"Mock-API-Data/constant"
	"Mock-API-Data/storage"

	"github.com/gin-gonic/gin"
)

func StorageMiddleware(storage *storage.Storage) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(constant.MiddlewareStorageKey, storage)
		c.Next()
	}
}
