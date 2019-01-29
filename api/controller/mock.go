package controller

import (
	"Mock-API-Data/model"
	"Mock-API-Data/util"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Mock struct {
}

type mockParam struct {
	ProjectId int64 `uri:"projectId" form:"projectId" json:"projectId" binding:"required"`
	RuleId    int64 `uri:"ruleId" form:"ruleId" json:"ruleId" binding:"required"`
}

func (m *Mock) Handler(c *gin.Context) {
	loginUser, storageHelper, ok := ExtractLoginUserAndStorageHelper(c)
	if !ok {
		c.Writer.WriteHeader(http.StatusServiceUnavailable)
		return
	}
	var param mockParam
	if err := c.ShouldBindUri(&param); err != nil {
		c.JSON(http.StatusBadRequest, util.GenerateErrorResponse(400, err.Error()))
		return
	}

	var rule model.Rule
	err := storageHelper.DB().
		Where("project_id = ? and id = ? and user_id = ?",
			param.ProjectId, param.RuleId, loginUser.Id).
		Find(&rule).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, util.GenerateErrorResponse(400, err.Error()))
		return
	}

	if rule.Method != c.Request.Method {
		c.Writer.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	if !rule.Enable {
		c.JSON(http.StatusBadRequest, util.GenerateErrorResponse(400, "改规则未启用"))
		return
	}

	var data model.Data
	err = storageHelper.DB().
		Where("rule_id = ? and user_id = ?", param.RuleId, loginUser.Id).
		Find(&data).Error

	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			c.JSON(http.StatusBadRequest, util.GenerateErrorResponse(400, "没有该规则数据,请先创建对应的数据"))
		} else {
			c.JSON(http.StatusBadRequest, util.GenerateErrorResponse(400, err.Error()))
		}
		return
	}
	switch data.ContentType {
	case "json":
		var jsonObj map[string]interface{}
		err := json.Unmarshal([]byte(data.Content), &jsonObj)
		if err != nil {
			c.JSON(http.StatusBadRequest, util.GenerateErrorResponse(400, "对应规则数据不是合法的json数据"))
			return
		}
		c.JSON(data.ResponseCode, jsonObj)
	case "xml":
	}
}
