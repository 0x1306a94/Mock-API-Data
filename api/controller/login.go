package controller

import (
	"Mock-API-Data/model"
	"Mock-API-Data/util"
	"net/http"

	"github.com/jinzhu/gorm"

	"github.com/gin-gonic/gin"
)

type loginParam struct {
	UserName string `form:"userName" json:"userName" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	storageHelper, ok := ExtractStorageHelper(c)
	if !ok {
		c.Writer.WriteHeader(http.StatusServiceUnavailable)
		return
	}
	var login loginParam
	if err := c.ShouldBind(&login); err != nil {
		c.JSON(http.StatusOK, util.GenerateErrorResponse(400, err.Error()))
		return
	}
	var user model.User
	err := storageHelper.DB().Where("name = ? OR email = ?", login.UserName, login.UserName).First(&user).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			c.JSON(http.StatusOK, util.GenerateErrorResponse(400, "没有该用户"))
		} else {
			c.JSON(http.StatusOK, util.GenerateErrorResponse(400, err.Error()))
		}
		return
	}

	if user.Id == 0 {
		c.JSON(http.StatusOK, util.GenerateErrorResponse(400, "用户名或者邮箱错误"))
		return
	}

	if !util.ValidationPassword(login.Password, user.Password) {
		c.JSON(http.StatusOK, util.GenerateErrorResponse(400, "密码错误"))
		return
	}

	tokenStr, err := util.GenerateAuthorizationToken(user)
	if err != nil {
		c.JSON(http.StatusOK, util.GenerateErrorResponse(400, err.Error()))
		return
	}

	c.JSON(http.StatusOK, util.GenerateSuccessResponse(gin.H{
		"user":  user,
		"token": tokenStr,
	}))
}
