package controller

import (
	"Mock-API-Data/model"
	"Mock-API-Data/util"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Rule struct {
}

type ruleCreateParam struct {
	ProjectId int64  `form:"projectId" json:"projectId" binding:"required"`
	Path      string `form:"path" json:"path" binding:"required"`
	Method    string `form:"method" json:"method" binding:"required"`
	Enable    bool   `form:"enable" json:"enable"`
}

type ruleUpdateParam struct {
	RuleId int64  `form:"ruleId" json:"ruleId" binding:"required"`
	Path   string `form:"path" json:"path" binding:"required"`
	Method string `form:"method" json:"method" binding:"required"`
	Enable bool   `form:"enable" json:"enable" binding:"required"`
}

type ruleInfoParam struct {
	RuleId int64 `uri:"ruleId" form:"ruleId" json:"ruleId" binding:"required"`
}

type rulePageParam struct {
	ProjectId int64 `uri:"projectId" form:"projectId" json:"projectId" binding:"required"`
	PageNo    int64 `uri:"pageNo" form:"pageNo" json:"pageNo"`
	PageSize  int64 `uri:"pageNo" form:"pageSize" json:"pageSize"`
}

type ruleDeleteParam struct {
	RuleId int64 `form:"ruleId" json:"ruleId" binding:"required"`
}

// 创建规则 POST
func (r *Rule) Create(c *gin.Context) {
	loginUser, storageHelper, ok := ExtractLoginUserAndStorageHelper(c)
	if !ok {
		c.Writer.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	var param ruleCreateParam
	if err := c.ShouldBind(&param); err != nil {
		c.JSON(http.StatusBadRequest, util.GenerateErrorResponse(400, err.Error()))
		return
	}

	var tmp model.Rule
	err := storageHelper.DB().
		Where("project_id = ? and user_id = ? and path = ? and method = ?",
			param.ProjectId, loginUser.Id, param.Path, param.Method).
		Find(&tmp).Error

	if err != nil && !gorm.IsRecordNotFoundError(err) {
		c.JSON(http.StatusOK, util.GenerateErrorResponse(400, err.Error()))
		return
	}
	if tmp.Id > 0 {
		c.JSON(http.StatusBadRequest, util.GenerateErrorResponse(400, "已经存在,请勿重复提交"))
		return
	}

	tt := time.Now()
	rule := &model.Rule{
		ProjectId: param.ProjectId,
		UserId:    loginUser.Id,
		Path:      param.Path,
		Method:    param.Method,
		Enable:    param.Enable,
		CreatedAt: tt,
		UpdateAt:  tt,
	}
	err = storageHelper.DB().
		Create(rule).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, util.GenerateErrorResponse(400, err.Error()))
		return
	}
	c.JSON(http.StatusOK, util.GenerateSuccessResponse(rule))
}

// 获取规则 GET
func (r *Rule) Info(c *gin.Context) {
	loginUser, storageHelper, ok := ExtractLoginUserAndStorageHelper(c)
	if !ok {
		c.Writer.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	var param ruleInfoParam
	if err := c.ShouldBind(&param); err != nil {
		c.JSON(http.StatusBadRequest, util.GenerateErrorResponse(400, err.Error()))
		return
	}
	var rule model.Rule
	err := storageHelper.DB().
		Where("rule_id = ? and user_id = ?", param.RuleId, loginUser.Id).
		Find(&rule).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, util.GenerateErrorResponse(400, err.Error()))
		return
	}
	c.JSON(http.StatusOK, util.GenerateSuccessResponse(rule))
}

// 创建规则 POST
func (r *Rule) Update(c *gin.Context) {
	loginUser, storageHelper, ok := ExtractLoginUserAndStorageHelper(c)
	if !ok {
		c.Writer.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	var param ruleUpdateParam
	if err := c.ShouldBind(&param); err != nil {
		c.JSON(http.StatusBadRequest, util.GenerateErrorResponse(400, err.Error()))
		return
	}

	tt := time.Now()
	rule := &model.Rule{
		Id:     param.RuleId,
		UserId: loginUser.Id,
	}
	err := storageHelper.DB().
		Model(rule).
		Updates(&model.Rule{
			Path:     param.Path,
			Method:   param.Method,
			Enable:   param.Enable,
			UpdateAt: tt}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, util.GenerateErrorResponse(400, err.Error()))
		return
	}
	c.JSON(http.StatusOK, util.GenerateSuccessResponse(true))
}

// 删除规则 POST
func (r *Rule) Delete(c *gin.Context) {
	loginUser, storageHelper, ok := ExtractLoginUserAndStorageHelper(c)
	if !ok {
		c.Writer.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	var param ruleDeleteParam
	if err := c.ShouldBind(&param); err != nil {
		c.JSON(http.StatusBadRequest, util.GenerateErrorResponse(400, err.Error()))
		return
	}

	rule := &model.Rule{
		Id:     param.RuleId,
		UserId: loginUser.Id,
	}
	err := storageHelper.DB().
		Delete(rule).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, util.GenerateErrorResponse(400, err.Error()))
		return
	}

	c.JSON(http.StatusOK, util.GenerateSuccessResponse(rule))
}

// 创建规则 GET
func (r *Rule) List(c *gin.Context) {

	loginUser, storageHelper, ok := ExtractLoginUserAndStorageHelper(c)
	if !ok {
		c.Writer.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	var param rulePageParam
	if err := c.ShouldBind(&param); err != nil {
		c.JSON(http.StatusBadRequest, util.GenerateErrorResponse(400, err.Error()))
		return
	}

	if param.PageNo <= 0 {
		param.PageNo = 1
	}

	if param.PageSize <= 0 {
		param.PageSize = 10
	}

	var rules []*model.Rule
	err := storageHelper.DB().
		Where("user_id = ? and project_id = ?", loginUser.Id, param.ProjectId).
		Order("created_at desc").
		Offset(((param.PageNo - 1) * param.PageSize)).
		Limit(param.PageSize).
		Find(&rules).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, util.GenerateErrorResponse(400, err.Error()))
		return
	}
	c.JSON(http.StatusOK, util.GenerateSuccessResponse(rules))
}
