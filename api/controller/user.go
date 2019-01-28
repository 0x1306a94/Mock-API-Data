package controller

import (
	"Mock-API-Data/model"
	"Mock-API-Data/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type User struct {
}

func (u *User) Info(c *gin.Context) {
	userId := c.Param("id")
	if userId == "" {
		loginUser, ok := ExtractLoginUser(c)
		if !ok {
			c.Writer.WriteHeader(http.StatusServiceUnavailable)
			return
		}
		c.JSON(http.StatusOK, util.GenerateSuccessResponse(loginUser))
		return
	}
	storageHelper, ok := ExtractStorageHelper(c)
	if !ok {
		c.Writer.WriteHeader(http.StatusServiceUnavailable)
		return
	}
	id, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, util.GenerateErrorResponse(400, "参数不合法"))
		return
	}
	user := model.User{
		Id: id,
	}
	err = storageHelper.DB().First(&user).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			c.JSON(http.StatusOK, util.GenerateErrorResponse(400, "没有该用户"))
		} else {
			c.JSON(http.StatusOK, util.GenerateErrorResponse(400, err.Error()))
		}
		return
	}
	c.JSON(http.StatusOK, util.GenerateSuccessResponse(user))
}
