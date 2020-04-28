package mapper

import (
	"myblogs/dtos"
	"myblogs/models"
)

func CreateArticleDto2ArticleModel(dto *dtos.ArticleDto) *models.ArticleModel {
	return &models.ArticleModel{
		ID:         dto.ID,
		Title:      dto.Title,
		LogoUrl:    dto.LogoUrl,
		Overview:   dto.Overview,
		Content:    dto.Content,
		CategoryID: dto.CategoryID,
		TagIDs:     dto.Tags,
		Status:     dto.Status,
	}
}

func ArticleModel2ArticleDto(model *models.ArticleModel) *dtos.ArticleDto {
	return &dtos.ArticleDto{
		ID:          model.ID,
		Title:       model.Title,
		LogoUrl:     model.LogoUrl,
		Overview:    model.Overview,
		Content:     model.Content,
		CategoryID:  model.CategoryID,
		CreateTime:  model.CreateTime,
		UpdateTime:  model.UpdateTime,
		PublishTime: model.PublishTime,
		Tags:        model.TagIDs,
		Status:      model.Status,
	}
}

func CategoryModel2CategoryDto(model *models.CategoryModel) *dtos.CategoryDto {
	return &dtos.CategoryDto{
		ID:   model.ID,
		Name: model.Name,
	}
}

func TagModel2TagDto(model *models.TagModel) *dtos.TagDto {
	return &dtos.TagDto{
		ID:   model.ID,
		Name: model.Name,
	}
}

func UploadResModel2UploadResDto(model *models.UploadResModel) *dtos.UploadResDto {
	return &dtos.UploadResDto{
		FileHash:   model.FileHash,
		Name:       model.Name,
		Ext:        model.Ext,
		Size:       model.Size,
		UploadTime: model.UploadTime,
	}
}
