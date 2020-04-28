package controller

import (
	"github.com/gin-gonic/gin"
	"myblogs/dtos"
	"myblogs/service"
	"myblogs/util/err"
	"net/http"
	"net/url"
	"strconv"
)

type categoryController struct {

}


var CategoryControllerImpl =&categoryController{}


func (c categoryController) GetCategories(context *gin.Context) {
	result:=service.CategoryServiceImpl.GetCategories()
	context.JSON(http.StatusOK,result)
}

func (c categoryController) EditCategory(context *gin.Context) {
	idStr:=context.Param("id")
	id,e:=strconv.Atoi(idStr)
	if e!=nil{
		context.JSON(http.StatusOK,dtos.NotOk(err.NoVliedParames))
		return
	}
	name:=context.Param("name")
	result:=service.CategoryServiceImpl.EditCategory(id,name)
	context.JSON(http.StatusOK,result)
}

func (c categoryController) AddCategory(context *gin.Context) {
	name,_:=url.QueryUnescape(context.Param("name"))
	result:=service.CategoryServiceImpl.AddCategory(name)
	context.JSON(http.StatusOK,result)
}


