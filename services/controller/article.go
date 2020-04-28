package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"myblogs/dtos"
	"myblogs/mapper"
	"myblogs/service"
	"myblogs/util/err"
	"net/http"
	"strconv"
)

type articleController struct {
}

var ArticeControllerImpl = &articleController{}

func (*articleController) GetArticles(context *gin.Context) {
	result := service.ArticleServiceImpl.GetArticles()
	context.JSON(http.StatusOK, result)
}

func (*articleController) GetArticle(context *gin.Context) {
	idStr:=context.Param("id")
	id,e:=strconv.Atoi(idStr)
	if e!=nil{
		context.JSON(http.StatusOK,dtos.NotOk(err.NoVliedParames))
		return
	}
	result:= service.ArticleServiceImpl.GetArticleByID(id)
	context.JSON(http.StatusOK,result)
}

func (*articleController) DelArticle(context *gin.Context) {
	idStr:=context.Param("id")
	id,e:=strconv.Atoi(idStr)
	if e!=nil{
		context.JSON(http.StatusOK,dtos.NotOk(err.NoVliedParames))
		return
	}
	result:=service.ArticleServiceImpl.DelArticleByID(id)
	context.JSON(http.StatusOK,result)
}

func (*articleController) PublishArticle(context *gin.Context) {
	idStr:=context.Param("id")
	id,e:=strconv.Atoi(idStr)
	if e!=nil{
		context.JSON(http.StatusOK,dtos.NotOk(err.NoVliedParames))
		return
	}
	result:=service.ArticleServiceImpl.PublishArticleByID(id)
	context.JSON(http.StatusOK,result)
}

func (c *articleController) AddArticle(context *gin.Context) {
	var createArticle dtos.ArticleDto
	e := context.ShouldBindJSON(&createArticle)
	if e != nil {
		log.Println(e)
		dtos.NotOk(err.NoVliedParames)
		return
	}
	articleModel := mapper.CreateArticleDto2ArticleModel(&createArticle)
	result := service.ArticleServiceImpl.AddArticle(articleModel)
	context.JSON(http.StatusOK, result)
}
