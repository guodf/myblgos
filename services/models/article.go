package models

type ArticleModel struct {
	ID          int    `gorm:"primary_key"`
	Title       string `gorm:"not null"`
	LogoUrl     string
	Overview    string
	Content     string `gorm:"not null"`
	CategoryID  int    `gorm:"DEFAULT:(0)"`
	TagIDs      []int  `gorm:"-"`
	CreateTime  int64  `gorm:"DEFAULT:(strftime('%s', 'now'))"`
	UpdateTime  int64
	PublishTime int64
	Status      int `gorm:"DEFAULT:(0)"`
}

func (ArticleModel) TableName() string {
	return "articles"
}

type ArticleTags struct {
	ArticleID int `gorm:"primary_key;auto_increment:false"`
	TagID     int `gorm:"primary_key;auto_increment:false"`
}

func (ArticleTags) TableName() string {
	return "article_tags"
}
