package models

type TagModel struct {
	ID   int    `gorm:"PRIMARY_KEY"`
	Name string `gorm:"unique;not null"`
}

func (TagModel) TableName() string {
	return "tags"
}
