package models

type UploadResModel struct {
	FileHash   string `gorm:"primary_key"`
	Name       string
	Ext        string
	Size       int64
	UploadTime int64 `gorm:"DEFAULT:(strftime('%s', 'now'))"`
}


func (UploadResModel) TableName()string{
	return "upload_res"
}
