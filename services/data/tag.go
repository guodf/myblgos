package data

import (
	"log"
	"myblogs/models"
)

type tagRepository struct {
	*repository
}

func NewTagRepository(work ...*unitOfWork) *tagRepository {
	return &tagRepository{newRepository(work)}
}

func (r *tagRepository) GetTags() []models.TagModel {
	var tagModels []models.TagModel
	r.work.db.Order("id desc").Find(&tagModels)
	return tagModels
}

func (r *tagRepository) EditTag(id int, name string) bool {
	e := r.work.db.Model(&models.TagModel{}).Where("id=?", id).UpdateColumn("name=?", name)
	return e == nil
}

func (r *tagRepository) AddTag(name string) *models.TagModel {
	model := &models.TagModel{
		Name: name,
	}
	e := r.work.db.Where(model).FirstOrCreate(model)
	if e.Error != nil {
		log.Println(e)
		return nil
	}
	return model
}

func (r *tagRepository) GetTagsById(id int) []int {
	var articleTags []models.ArticleTags
	e := r.work.db.Where("article_id=?", id).Find(&articleTags)
	if e.Error != nil {
		log.Println(e.Error)
	}
	var ids []int
	for _, tag := range articleTags {
		ids = append(ids, tag.TagID)
	}
	return ids
}
