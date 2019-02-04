package controller

import (
	"Mock-API-Data/model"
	"Mock-API-Data/util"
	"encoding/json"
	"encoding/xml"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var supportContentType map[string]int = map[string]int{
	"json": 1,
	"xml":  1,
	"html": 1,
	"text": 1,
}

type Data struct {
}

type dataCreateParam struct {
	RuleId       int64  `form:"ruleId" json:"ruleId" binding:"required"`
	ResponseCode int    `form:"responseCode" json:"responseCode" binding:"required"`
	ContentType  string `form:"contentType" json:"contentType" binding:"required"` // text html json xml
	Content      string `form:"content" json:"content" binding:"required"`
}

type dataUpdateParam struct {
	DataId       int64  `form:"dataId" json:"dataId" binding:"required"`
	ResponseCode int    `form:"responseCode" json:"responseCode" binding:"required"`
	ContentType  string `form:"contentType" json:"contentType" binding:"required"` // text html json xml
	Content      string `form:"content" json:"content" binding:"required"`
}

type dataDeleteParam struct {
	DataId int64 `form:"dataId" json:"dataId" binding:"required"`
}

type dataInfoParam struct {
	DataId int64 `uri:"dataId" form:"dataId" json:"dataId" binding:"required"`
}

// 创建数据
func (d *Data) Create(c *gin.Context) {
	loginUser, storageHelper, ok := ExtractLoginUserAndStorageHelper(c)
	if !ok {
		c.Writer.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	var param dataCreateParam
	if err := c.ShouldBind(&param); err != nil {
		c.JSON(http.StatusBadRequest, util.GenerateErrorResponse(400, err.Error()))
		return
	}

	var tmp model.Data
	err := storageHelper.DB().
		Where("rule_id = ? and user_id = ?", param.RuleId, loginUser.Id).
		Find(&tmp).Error

	if err != nil && !gorm.IsRecordNotFoundError(err) {
		c.JSON(http.StatusOK, util.GenerateErrorResponse(400, err.Error()))
		return
	}
	if tmp.Id > 0 {
		c.JSON(http.StatusBadRequest, util.GenerateErrorResponse(400, "已经存在,请勿重复提交"))
		return
	}

	if _, exsit := supportContentType[param.ContentType]; !exsit {
		c.JSON(http.StatusBadRequest, util.GenerateErrorResponse(400, "contentType 参数不合法"))
		return
	}

	if len(param.Content) == 0 {
		c.JSON(http.StatusBadRequest, util.GenerateErrorResponse(400, "content 参数内容为空"))
		return
	}

	switch param.ContentType {
	case "json":
		var jsonObj map[string]interface{}
		err := json.Unmarshal([]byte(param.Content), &jsonObj)
		if err != nil {
			c.JSON(http.StatusBadRequest, util.GenerateErrorResponse(400, "对应规则数据不是合法的JSON格式数据"))
			return
		}
	case "xml":
		var xmlObj map[string]interface{}
		err := xml.Unmarshal([]byte(param.Content), &xmlObj)
		if err != nil {
			c.JSON(http.StatusBadRequest, util.GenerateErrorResponse(400, "对应规则数据不是合法的XML格式数据"))
			return
		}
	default:
		// no thing
	}

	tt := time.Now()
	data := &model.Data{
		RuleId:       param.RuleId,
		UserId:       loginUser.Id,
		ResponseCode: param.ResponseCode,
		ContentType:  param.ContentType,
		Content:      param.Content,
		CreatedAt:    tt,
		UpdateAt:     tt,
	}
	err = storageHelper.DB().
		Create(data).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, util.GenerateErrorResponse(400, err.Error()))
		return
	}
	c.JSON(http.StatusOK, util.GenerateSuccessResponse(data))
}

// 更新数据
func (d *Data) Update(c *gin.Context) {
	loginUser, storageHelper, ok := ExtractLoginUserAndStorageHelper(c)
	if !ok {
		c.Writer.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	var param dataUpdateParam
	if err := c.ShouldBind(&param); err != nil {
		c.JSON(http.StatusBadRequest, util.GenerateErrorResponse(400, err.Error()))
		return
	}

	var tmp model.Data
	err := storageHelper.DB().
		Where("id = ? and user_id = ?",
			param.DataId,
			loginUser.Id).
		Find(&tmp).Error

	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			c.JSON(http.StatusBadRequest, util.GenerateErrorResponse(400, "没有对应的模拟数据"))
		} else {
			c.JSON(http.StatusBadRequest, util.GenerateErrorResponse(400, err.Error()))
		}
		return
	}

	tt := time.Now()
	tmp.ResponseCode = param.ResponseCode
	tmp.ContentType = param.ContentType
	tmp.Content = param.Content
	tmp.UpdateAt = tt
	err = storageHelper.DB().
		Save(&tmp).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, util.GenerateErrorResponse(400, err.Error()))
		return
	}

	c.JSON(http.StatusOK, util.GenerateSuccessResponse(true))
}

// 删除数据
func (d *Data) Delete(c *gin.Context) {
	loginUser, storageHelper, ok := ExtractLoginUserAndStorageHelper(c)
	if !ok {
		c.Writer.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	var param dataDeleteParam
	if err := c.ShouldBind(&param); err != nil {
		c.JSON(http.StatusBadRequest, util.GenerateErrorResponse(400, err.Error()))
		return
	}

	data := &model.Data{
		Id:     param.DataId,
		UserId: loginUser.Id,
	}

	err := storageHelper.DB().
		Delete(data).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, util.GenerateErrorResponse(400, err.Error()))
		return
	}
	c.JSON(http.StatusOK, util.GenerateSuccessResponse(true))
}

// 数据信息 GET
func (d *Data) Info(c *gin.Context) {
	loginUser, storageHelper, ok := ExtractLoginUserAndStorageHelper(c)
	if !ok {
		c.Writer.WriteHeader(http.StatusServiceUnavailable)
		return
	}
	var param dataInfoParam
	if err := c.ShouldBind(&param); err != nil {
		c.JSON(http.StatusBadRequest, util.GenerateErrorResponse(400, err.Error()))
		return
	}

	data := &model.Data{
		Id:     param.DataId,
		UserId: loginUser.Id,
	}

	err := storageHelper.DB().
		Find(data).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, util.GenerateErrorResponse(400, err.Error()))
		return
	}
	c.JSON(http.StatusOK, util.GenerateSuccessResponse(data))
}
