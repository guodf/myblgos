package data

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/viper"
	"myblogs/models"
	"log"
)

var db *gorm.DB

// init 初始化db
func init() {
	dbNew, err := gorm.Open("sqlite3", viper.GetString("db_file_path"))
	if err != nil {
		log.Fatal(err)
	}
	dbNew.AutoMigrate(&models.UploadResModel{})
	dbNew.AutoMigrate(&models.ArticleModel{})
	dbNew.AutoMigrate(&models.TagModel{})
	dbNew.AutoMigrate(&models.CategoryModel{})
	dbNew.AutoMigrate(&models.ArticleTags{})
	dbNew.Exec("INSERT INTO sqlite_sequence SELECT 'articles',2020000	WHERE NOT EXISTS (SELECT * FROM sqlite_sequence WHERE name='articles')");
	db = dbNew
}

type unitOfWork struct {
	db *gorm.DB
}

func NewUnitOfWork() *unitOfWork {
	return &unitOfWork{
		db: db.Begin(),
	}
}

//  Begin
func (work *unitOfWork) Completed() {
	if work.db == nil {
		return
	}
	e := work.db.Commit()
	if e != nil {
		work.db.Rollback()
	}
}
