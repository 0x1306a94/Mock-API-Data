package util

import (
	"github.com/gin-gonic/gin"
)

// 生成错误响应数据
func GenerateErrorResponse(errorCode int, errorMsg string) gin.H {
	return gin.H{
		"state":     1,
		"errorCode": errorCode,
		"errorMsg":  errorMsg,
		"data":      nil,
	}
}

// 生成成功响应数据
func GenerateSuccessResponse(data interface{}) gin.H {
	return gin.H{
		"state":     1,
		"errorCode": 0,
		"errorMsg":  "",
		"data":      data,
	}
}
