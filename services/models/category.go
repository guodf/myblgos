package models

type CategoryModel struct {
	ID   int    `gorm:"PRIMARY_KEY"`
	Name string `gorm:"unique;not null"`
}

func (CategoryModel) TableName()string  {
	return "categories"
}