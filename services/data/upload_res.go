package data

import (
	"myblogs/models"
)

type uploadResRepository struct {
	*repository
}

func NewUploadResRepository(work ...*unitOfWork) *uploadResRepository {
	return &uploadResRepository{newRepository(work)}
}

func (r uploadResRepository) GetUploadRes() []models.UploadResModel {
	var uploadResModels []models.UploadResModel
	r.work.db.Order("upload_time desc").Find(&uploadResModels)
	return uploadResModels
}

func (r uploadResRepository) AddUploadRes(model *models.UploadResModel) bool {
	e := r.work.db.Where(models.UploadResModel{
		FileHash: model.FileHash,
	}).Assign(models.UploadResModel{
		Name: model.Name,
		Ext:  model.Ext,
		Size: model.Size,
	}).FirstOrCreate(model)
	return e.Error == nil
}
