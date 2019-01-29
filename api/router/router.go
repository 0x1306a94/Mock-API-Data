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
	//rootRouter.Use(middleware.ReverseProxyMiddleware())

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
	}

	return rootRouter
}

// 创建mock 路由
func InitMockRouter(storage *storage.Storage) *gin.Engine {

	rootRouter := gin.Default()
	rootRouter.HandleMethodNotAllowed = true
	rootRouter.Use(middleware.StorageMiddleware(storage))
	rootRouter.Use(middleware.AuthorizedMiddleware())
	rootRouter.Use(middleware.ReverseProxyMiddleware())

	// mock
	{
		mockController := &controller.Mock{}
		mockRouter := rootRouter.Group("/mock")
		mockRouter.Any("/:projectId/:ruleId", mockController.Handler)
	}

	return rootRouter
}
