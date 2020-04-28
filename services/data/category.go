package data

import (
	"log"
	"myblogs/models"
)

type categoryRepository struct {
	*repository
}

func NewCategoryRepository(work ...*unitOfWork) *categoryRepository {
	return &categoryRepository{newRepository(work)}
}

func (r categoryRepository) GetCategories() []models.CategoryModel {
	var categoryModels []models.CategoryModel
	r.work.db.Order("id desc").Find(&categoryModels)
	return categoryModels
}

func (r categoryRepository) EditCategory(id int, name string) bool {
	e := r.work.db.Model(&models.CategoryModel{}).Where("id=?", id).UpdateColumn("name=?", name)
	return e == nil
}

func (r categoryRepository) AddCategory(name string) *models.CategoryModel {
	model := &models.CategoryModel{
		Name: name,
	}
	e := r.work.db.Where(model).FirstOrCreate(model)
	if e.Error != nil {
		log.Println(e)
		return nil
	}
	return model
}
