package controller

import (
	"Mock-API-Data/model"
	"Mock-API-Data/util"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Project struct {
}

type projectParam struct {
	ProjectId int64  `form:"projectId" json:"projectId"`
	Name      string `form:"name" json:"name" binding:"required"`
	Host      string `form:"host" json:"host" binding:"required"`
}

// 创建项目 POST
func (p *Project) Create(c *gin.Context) {
	loginUser, storageHelper, ok := ExtractLoginUserAndStorageHelper(c)
	if !ok {
		c.Writer.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	var param projectParam
	if err := c.Bind(&param); err != nil {
		c.JSON(http.StatusOK, util.GenerateErrorResponse(400, err.Error()))
		return
	}

	tt := time.Now()
	project := &model.Project{
		Key:       util.NewProjectKey(loginUser.Id),
		Name:      param.Name,
		Host:      param.Host,
		UserId:    loginUser.Id,
		CreatedAt: tt,
		UpdateAt:  tt,
	}
	err := storageHelper.DB().Create(project).Error
	if err != nil {
		c.JSON(http.StatusOK, util.GenerateErrorResponse(400, err.Error()))
		return
	}

	c.JSON(http.StatusOK, util.GenerateSuccessResponse(project))
}

// 更新项目 POST
func (p *Project) Update(c *gin.Context) {

	storageHelper, ok := ExtractStorageHelper(c)
	if !ok {
		c.Writer.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	var param projectParam
	if err := c.Bind(&param); err != nil {
		c.JSON(http.StatusOK, util.GenerateErrorResponse(400, err.Error()))
		return
	}
	if param.ProjectId == 0 {
		c.JSON(http.StatusBadRequest, util.GenerateErrorResponse(400, "projectId 参数为空"))
		return
	}

	tt := time.Now()
	project := &model.Project{
		Id: param.ProjectId,
	}
	err := storageHelper.DB().Model(project).Updates(&model.Project{Name: param.Name, Host: param.Host, UpdateAt: tt}).Error
	if err != nil {
		c.JSON(http.StatusOK, util.GenerateErrorResponse(400, err.Error()))
		return
	}
	c.JSON(http.StatusOK, util.GenerateSuccessResponse(true))
}

// 删除项目 POST
func (p *Project) Delete(c *gin.Context) {
	storageHelper, ok := ExtractStorageHelper(c)
	if !ok {
		c.Writer.WriteHeader(http.StatusServiceUnavailable)
		return
	}
	projectId := c.PostForm("projectId")
	if projectId == "" {
		c.JSON(http.StatusOK, util.GenerateErrorResponse(400, "参数为空"))
		return
	}
	id, err := strconv.ParseInt(projectId, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, util.GenerateErrorResponse(400, "参数不合法"))
		return
	}

	project := &model.Project{
		Id: id,
	}

	err = storageHelper.DB().Delete(project).Error
	if err != nil {
		c.JSON(http.StatusOK, util.GenerateErrorResponse(400, err.Error()))
		return
	}
	c.JSON(http.StatusOK, util.GenerateSuccessResponse(true))
}

// 项目信息 GET
func (p *Project) Info(c *gin.Context) {
	storageHelper, ok := ExtractStorageHelper(c)
	if !ok {
		c.Writer.WriteHeader(http.StatusServiceUnavailable)
		return
	}
	projectId := c.Query("projectId")
	if projectId == "" {
		c.JSON(http.StatusOK, util.GenerateErrorResponse(400, "参数为空"))
		return
	}
	id, err := strconv.ParseInt(projectId, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, util.GenerateErrorResponse(400, "参数不合法"))
		return
	}
	project := &model.Project{
		Id: id,
	}

	err = storageHelper.DB().Find(project).Error
	if err != nil {
		c.JSON(http.StatusOK, util.GenerateErrorResponse(400, err.Error()))
		return
	}
	c.JSON(http.StatusOK, util.GenerateSuccessResponse(project))
}

func (p *Project) List(c *gin.Context) {
	loginUser, storageHelper, ok := ExtractLoginUserAndStorageHelper(c)
	if !ok {
		c.Writer.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	var param pageParams
	if err := c.Bind(&param); err != nil {
		c.JSON(http.StatusOK, util.GenerateErrorResponse(400, err.Error()))
		return
	}

	if param.PageNo <= 0 {
		param.PageNo = 1
	}

	if param.PageSize <= 0 {
		param.PageSize = 10
	}

	var projects []*model.Project
	err := storageHelper.DB().Where("user_id = ?", loginUser.Id).Order("created_at desc").Offset(((param.PageNo - 1) * param.PageSize)).Limit(param.PageSize).Find(&projects).Error
	if err != nil {
		c.JSON(http.StatusOK, util.GenerateErrorResponse(400, err.Error()))
		return
	}
	c.JSON(http.StatusOK, util.GenerateSuccessResponse(projects))

}
