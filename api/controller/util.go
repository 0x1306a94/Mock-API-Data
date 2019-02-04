package controller

import (
	"Mock-API-Data/constant"
	"Mock-API-Data/model"
	"Mock-API-Data/storage"

	"github.com/gin-gonic/gin"
)

type pageParams struct {
	PageNo   int64 `uri:"pageNo" form:"pageNo" json:"pageNo" binding:"required"`
	PageSize int64 `uri:"pageSize" form:"pageSize" json:"pageSize" binding:"required"`
}

// 从上下文中 读取当前登录用户和存储实例
func ExtractLoginUserAndStorageHelper(c *gin.Context) (user *model.User, storageHelper *storage.Storage, ok bool) {

	val := c.MustGet(constant.MiddlewareLoginUserKey)
	if val == nil {
		ok = false
		return
	}

	user, ok = val.(*model.User)
	if !ok {
		return
	}

	val = c.MustGet(constant.MiddlewareStorageKey)
	if val == nil {
		ok = false
		return
	}

	storageHelper, ok = val.(*storage.Storage)
	return
}

// 从上下文中 读取当前登录用户
func ExtractLoginUser(c *gin.Context) (user *model.User, ok bool) {
	val := c.MustGet(constant.MiddlewareLoginUserKey)
	if val == nil {
		ok = false
		return
	}
	user, ok = val.(*model.User)
	return
}

// 从上下文中 读取存储实例
func ExtractStorageHelper(c *gin.Context) (storageHelper *storage.Storage, ok bool) {
	val := c.MustGet(constant.MiddlewareStorageKey)
	if val == nil {
		ok = false
		return
	}
	storageHelper, ok = val.(*storage.Storage)
	return
}
