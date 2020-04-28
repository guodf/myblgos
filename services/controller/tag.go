package controller

import (
	"github.com/gin-gonic/gin"
	"myblogs/dtos"
	"myblogs/service"
	"myblogs/util/err"
	"net/http"
	"strconv"
)

type tagController struct {

}


var TagControllerImpl =&tagController{}


	func (c tagController) GetTags(context *gin.Context) {
		result:=service.TagServiceImpl.GetTags()
		context.JSON(http.StatusOK,result)
	}

	func (c tagController) EditTagByID(context *gin.Context) {
		idStr:=context.Param("id")
		id,e:=strconv.Atoi(idStr)
		if e!=nil{
			context.JSON(http.StatusOK,dtos.NotOk(err.NoVliedParames))
			return
		}
		name:=context.Param("name")
		result:=service.TagServiceImpl.EditTagByID(id,name)
		context.JSON(http.StatusOK,result)
	}

	func (c tagController) AddTag(context *gin.Context) {
		name:=context.Param("name")
		result:=service.TagServiceImpl.AddTag(name)
		context.JSON(http.StatusOK,result)
	}