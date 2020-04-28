package service

import (
	"myblogs/data"
	"myblogs/dtos"
	"myblogs/mapper"
	"myblogs/models"
	"myblogs/util/err"
)

type articleService struct {
}

var ArticleServiceImpl = &articleService{}

func (s *articleService) AddArticle(model *models.ArticleModel) *dtos.ResultData {
	work := data.NewUnitOfWork()
	defer work.Completed()

	articleRepo := data.NewArticleRepository(work)

	articleRepo.AddArticle(model)

	if nil != model && model.ID > 0 {
		tIds := articleRepo.AddTagsToArticle(model.ID, model.TagIDs)
		model.TagIDs = tIds
		return dtos.Ok(mapper.ArticleModel2ArticleDto(model))
	}

	return dtos.NotOk(err.AddFailed)
}

func (s *articleService) GetArticles() *dtos.ResultData {
	articleModels := data.NewArticleRepository().GetArticles()
	var articleDtos []*dtos.ArticleDto
	for _, model := range articleModels {
		dto := mapper.ArticleModel2ArticleDto(&model)
		articleDtos = append(articleDtos, dto)
	}
	return dtos.Ok(articleDtos)
}

func (s *articleService) GetArticleByID(id int) *dtos.ResultData {
	model := data.NewArticleRepository().GetArticleByID(id)
	model.TagIDs = data.NewTagRepository().GetTagsById(id)
	dto := mapper.ArticleModel2ArticleDto(model)
	return dtos.Ok(dto)
}

func (s *articleService) DelArticleByID(id int) *dtos.ResultData {
	articleRepo := data.NewArticleRepository()
	articleRepo.DelArticleByID(id)
	return dtos.Ok()
}

func (s *articleService) PublishArticleByID(id int) *dtos.ResultData {
	articleRepo := data.NewArticleRepository()
	articleRepo.PublishArticleID(id)
	return dtos.Ok()
}
