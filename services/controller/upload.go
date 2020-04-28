package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	"myblogs/dtos"
	"myblogs/service"
	"myblogs/util"
	"myblogs/util/err"
	"net/http"
	"os"
	"path/filepath"
)

type uploadController struct {
}

var UploadControllerImpl = &uploadController{}

func (c *uploadController) Upload(context *gin.Context) {
	//计算文件唯一标识
	file, e := context.FormFile("file")
	//文件存在
	if file == nil || e != nil {
		context.JSON(http.StatusOK, dtos.NotOk(err.UploadFailed))
		return
	}
	fileHash, e := util.FileToken(file)
	if e != nil || fileHash == "" {
		context.JSON(http.StatusOK, dtos.NotOk(err.UploadFailed))
		return
	}
	//保存文件
	os.MkdirAll(viper.GetString("uploads"), os.ModeDir)
	tarPath := filepath.Join(viper.GetString("uploads"), fileHash+filepath.Ext(file.Filename))
	log.Println(tarPath)
	e = context.SaveUploadedFile(file, tarPath)
	if e != nil {
		log.Println("保存文件失败", e.Error())
		context.JSON(http.StatusOK, dtos.NotOk(err.UploadFailed))
		return
	}
	result:=service.UploadServiceImpl.AddRes(fileHash,file.Filename,file.Size)
	context.JSON(http.StatusOK,result)
}

func (c *uploadController) GetRes(context *gin.Context) {
	result:= service.UploadServiceImpl.GetRes()
	context.JSON(http.StatusOK,result)
}
