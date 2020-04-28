package service

import (
	"fmt"
	"myblogs/data"
	"myblogs/dtos"
	"myblogs/mapper"
	"myblogs/models"
	"myblogs/util/err"
	"path/filepath"
)

type uploadService struct {
}

var UploadServiceImpl = &uploadService{}

func (s *uploadService) AddRes(hash string, filename string, size int64) *dtos.ResultData {
	model := &models.UploadResModel{
		FileHash: hash,
		Name:     filename,
		Ext:      filepath.Ext(filename),
		Size:     size,
	}
	ok := data.NewUploadResRepository().AddUploadRes(model)
	if ok {
		return dtos.Ok(fmt.Sprintf("/res/%s%s",model.FileHash,model.Ext));
	}
	return dtos.NotOk(err.AddFailed)
}

func (s *uploadService) GetRes() *dtos.ResultData {
	resModels:=data.NewUploadResRepository().GetUploadRes()
	var resDtos []*dtos.UploadResDto
	for _, resModel := range resModels {
		resDtos = append(resDtos, mapper.UploadResModel2UploadResDto(&resModel))
	}
	return dtos.Ok(resDtos)
}
