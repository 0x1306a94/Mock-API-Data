package router

import (
	"Mock-API-Data/api/controller"
	"Mock-API-Data/api/middleware"
	"Mock-API-Data/storage"

	"github.com/gin-gonic/gin"
)

// 创建dashboard 路由
func InitDashboardRouter(storage *storage.Storage) *gin.Engine {
	rootRouter := gin.Default()
	rootRouter.HandleMethodNotAllowed = true
	rootRouter.Use(middleware.StorageMiddleware(storage))

	{
		rootRouter.POST("/login", controller.Login)
		rootRouter.POST("/registered", controller.Registered)
	}

	{
		authorizedRouter := rootRouter.Group("/admin")
		authorizedRouter.Use(middleware.AuthorizedMiddleware())
		{
			userController := &controller.User{}
			userRouter := authorizedRouter.Group("/user")
			userRouter.GET("/:userId", userController.Info)
			userRouter.GET("/", userController.Info)
		}

		{
			projectController := &controller.Project{}
			projectRouter := authorizedRouter.Group("/project")
			projectRouter.POST("/create", projectController.Create)
			projectRouter.GET("/info", projectController.Info)
			projectRouter.POST("/update", projectController.Update)
			projectRouter.POST("/delete", projectController.Delete)
			projectRouter.GET("/list", projectController.List)
		}

		{
			ruleController := &controller.Rule{}
			ruleRouter := authorizedRouter.Group("/rule")
			ruleRouter.POST("/create", ruleController.Create)
			ruleRouter.GET("/info", ruleController.Info)
			ruleRouter.POST("/update", ruleController.Update)
			ruleRouter.POST("/delete", ruleController.Delete)
			ruleRouter.GET("/list", ruleController.List)
		}

		{
			dataController := &controller.Data{}
			dataRouter := authorizedRouter.Group("/data")
			dataRouter.POST("/create", dataController.Create)
			dataRouter.GET("/info", dataController.Info)
			dataRouter.POST("/update", dataController.Update)
			dataRouter.POST("/delete", dataController.Delete)
		}
	}

	return rootRouter
}

// 创建mock 路由
func InitMockRouter(storage *storage.Storage) *gin.Engine {

	rootRouter := gin.Default()
	rootRouter.HandleMethodNotAllowed = true
	// 注入数据库实例
	rootRouter.Use(middleware.StorageMiddleware(storage))
	// token 授权检查
	rootRouter.Use(middleware.AuthorizedMiddleware())
	// 代理请求
	rootRouter.Use(middleware.ReverseProxyMiddleware())

	// mock
	{
		mockController := &controller.Mock{}
		mockRouter := rootRouter.Group("/mock")
		mockRouter.GET("/:projectId/:ruleId", mockController.Handler)
		mockRouter.POST("/:projectId/:ruleId", mockController.Handler)
		mockRouter.PUT("/:projectId/:ruleId", mockController.Handler)
		mockRouter.DELETE("/:projectId/:ruleId", mockController.Handler)
		mockRouter.HEAD("/:projectId/:ruleId", mockController.Handler)
		mockRouter.OPTIONS("/:projectId/:ruleId", mockController.Handler)
		mockRouter.PATCH("/:projectId/:ruleId", mockController.Handler)
	}

	return rootRouter
}
