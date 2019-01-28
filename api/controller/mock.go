package controller

import (
	"github.com/gin-gonic/gin"
)

type Mock struct {
}

func (m *Mock) Handler(c *gin.Context) {
	// user, ok := c.MustGet("_user").(*model.User)
	// if !ok {
	// 	c.Writer.WriteHeader(http.StatusServiceUnavailable)
	// 	return
	// }

	// storageHelper, ok := c.MustGet("_storage").(*storage.Storage)
	// if !ok {
	// 	c.Writer.WriteHeader(http.StatusServiceUnavailable)
	// 	return
	// }
}
