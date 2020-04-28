package service

import (
	"myblogs/data"
	"myblogs/dtos"
	"myblogs/mapper"
)

type categoryService struct {
}

var CategoryServiceImpl = &categoryService{}

func (s categoryService) GetCategories() *dtos.ResultData {

	categoryModels := data.NewCategoryRepository().GetCategories()
	var categoryDtos []*dtos.CategoryDto
	for _, model := range categoryModels {
		categoryDtos = append(categoryDtos, mapper.CategoryModel2CategoryDto(&model))
	}

	return dtos.Ok(categoryDtos)
}

func (s categoryService) EditCategory(id int, name string) *dtos.ResultData {
	data.NewCategoryRepository().EditCategory(id, name)
	return dtos.Ok()
}

func (s categoryService) AddCategory(name string) *dtos.ResultData {
	model:=data.NewCategoryRepository().AddCategory(name)
	return dtos.Ok(mapper.CategoryModel2CategoryDto(model))
}
