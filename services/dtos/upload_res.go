package dtos

type UploadResDto struct {
	FileHash   string `json:"fileHash"`
	Name       string `json:"name"`
	Ext        string `json:"ext"`
	Size       int64 `json:"size"`
	UploadTime int64 `json:"uploadTime"`
}
