package router

import (
	"github.com/spf13/viper"
	"myblogs/controller"
	"myblogs/middleware"

	"github.com/gin-gonic/gin"
)

// InitRouter loading router
func InitRouter() *gin.Engine {
	r := gin.Default()

	// cors 跨域调用
	r.Use(middleware.Cors())

	r.Static("/res",viper.GetString("uploads"))

	r.POST("/login", Login)
	// 验证
	rJwt := r.Use(middleware.JwtAuth())
	// 文章管理
	rJwt.GET("/articles", controller.ArticeControllerImpl.GetArticles)
	rJwt.GET("/articles/:id", controller.ArticeControllerImpl.GetArticle)
	rJwt.POST("/articles", controller.ArticeControllerImpl.AddArticle)
	rJwt.DELETE("/articles/:id", controller.ArticeControllerImpl.DelArticle)
	rJwt.PUT("/articles/:id", controller.ArticeControllerImpl.PublishArticle)

	// 分类管理
	rJwt.GET("/categories", controller.CategoryControllerImpl.GetCategories)
	rJwt.POST("/categories/:name", controller.CategoryControllerImpl.AddCategory)
	rJwt.PUT("/categories/:id/:name", controller.CategoryControllerImpl.EditCategory)
	//rJwt.DELETE("/categories/:id", controller.CategoryControllerImpl.DelCategory)

	// 标签管理
	rJwt.GET("/tags", controller.TagControllerImpl.GetTags)
	rJwt.POST("/tags/:name", controller.TagControllerImpl.AddTag)
	rJwt.PUT("/tags/:id/:name", controller.TagControllerImpl.EditTagByID)
	//rJwt.DELETE("/tags/:id", controller.TagControllerImpl.DelTag)

	// 图片库
	rJwt.GET("/res",controller.UploadControllerImpl.GetRes)
	rJwt.POST("/upload",controller.UploadControllerImpl.Upload)
	return r
}
