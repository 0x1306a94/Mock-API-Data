package controller

import (
	"Mock-API-Data/model"
	"Mock-API-Data/util"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type regParam struct {
	UserName        string `form:"userName" json:"userName" binding:"required"`
	Email           string `form:"email" json:"email" binding:"required"`
	Password        string `form:"password" json:"password" binding:"required"`
	ConfirmPassword string `form:"confirmPassword" json:"confirmPassword" binding:"required"`
}

func Registered(c *gin.Context) {
	storageHelper, ok := ExtractStorageHelper(c)
	if !ok {
		c.Writer.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	var reg regParam
	if err := c.Bind(&reg); err != nil {
		c.JSON(http.StatusOK, util.GenerateErrorResponse(400, err.Error()))
		return
	}

	if !util.CheckEmail(reg.Email) {
		c.JSON(http.StatusOK, util.GenerateErrorResponse(400, "邮箱格式错误"))
		return
	}

	if reg.ConfirmPassword != reg.Password {
		c.JSON(http.StatusOK, util.GenerateErrorResponse(400, "密码与确认密码不符"))
		return
	}

	encryptionPassword := util.EncryptionPassword(reg.Password)
	t := time.Now()
	user := &model.User{
		Name:     reg.UserName,
		Email:    reg.Email,
		Password: encryptionPassword,
		CreateAt: t,
		UpdateAt: t,
	}

	err := storageHelper.DB().Create(user).Error
	if err != nil {
		c.JSON(http.StatusOK, util.GenerateErrorResponse(400, err.Error()))
		return
	}

	c.JSON(http.StatusOK, util.GenerateSuccessResponse(user))
}
