package service

import (
	"myblogs/data"
	"myblogs/dtos"
	"myblogs/mapper"
	"myblogs/util/err"
)

type tagService struct {
	
}

var TagServiceImpl=&tagService{}


func (s tagService) GetTags() *dtos.ResultData {
	tagModels := data.NewTagRepository().GetTags()
	var tagDtos []*dtos.TagDto
	for _, model := range tagModels {
		tagDtos = append(tagDtos, mapper.TagModel2TagDto(&model))
	}

	return dtos.Ok(tagDtos)
}

func (s tagService) EditTagByID(id int, name string) *dtos.ResultData {
	data.NewTagRepository().EditTag(id, name)
	return dtos.Ok()
}

func (s tagService) AddTag(name string) *dtos.ResultData {
	model:=data.NewTagRepository().AddTag(name)
	if model==nil{
		return dtos.NotOk(err.AddFailed)
	}
	return dtos.Ok(mapper.TagModel2TagDto(model))
}
