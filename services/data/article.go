package data

import (
	"myblogs/models"
	"time"
)

type articleRepository struct {
	*repository
}

func NewArticleRepository(work ...*unitOfWork) *articleRepository {
	return &articleRepository{newRepository(work)}
}

func (r *articleRepository) GetArticles() []models.ArticleModel {
	var articles []models.ArticleModel
	r.work.db.Order("publish_time desc").Find(&articles)
	return articles
}

func (r *articleRepository) AddArticle(model *models.ArticleModel) *models.ArticleModel {
	model.CreateTime = time.Now().Unix()
	model.UpdateTime = time.Now().Unix()
	e := r.work.db.Where("id=?", model.ID).Assign(models.ArticleModel{
		Title:      model.Title,
		LogoUrl:    model.LogoUrl,
		Overview:   model.Overview,
		Content:    model.Content,
		CategoryID: model.CategoryID,
		UpdateTime: model.UpdateTime,
	}).FirstOrCreate(model)
	if e != nil {
		return nil
	}
	return model
}

func (r *articleRepository) AddTagsToArticle(id int, tagIds []int) []int {
	var ids []int
	for _, tId := range tagIds {
		e := r.work.db.FirstOrCreate(&models.ArticleTags{
			ArticleID: id,
			TagID:     tId,
		})
		if e.Error == nil {
			ids = append(ids, tId)
		}
	}
	return ids
}

func (r *articleRepository) GetArticleByID(id int) *models.ArticleModel {
	var model models.ArticleModel
	e := r.work.db.Where("id=?", id).First(&model)
	if e != nil {
		return &model
	}
	return nil
}

func (r *articleRepository) DelArticleByID(id int) bool {
	e := r.work.db.Model(&models.ArticleModel{}).UpdateColumn("status", 2).Where("id=?", id)
	return e == nil
}

func (r *articleRepository) PublishArticleID(id int) bool {
	e := r.work.db.Model(&models.ArticleModel{}).Where("id=?", id).UpdateColumns(map[string]interface{}{
		"status":       1,
		"publish_time": time.Now().Unix(),
	})
	return e == nil
}
