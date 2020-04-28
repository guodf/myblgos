package dtos

type CreateArticleDto struct {
	Title      string `json:"title"`
	LogoUrl    string `json:"logoUrl"`
	Overview   string `json:"overview"`
	Content    string `json:"content"`
	CategoryID int    `json:"categoryId"`
	Tags       []int  `json:"tags"`
	Status     int    `json:"status"`
}

type ArticleDto struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	LogoUrl     string `json:"logoUrl"`
	Overview    string `json:"overview"`
	Content     string `json:"content"`
	CategoryID  int    `json:"categoryId"`
	CreateTime  int64  `json:"createTime"`
	UpdateTime  int64  `json:"updateTime"`
	PublishTime int64  `json:"publishTime"`
	Tags        []int  `json:"tags"`
	Status      int    `json:"status"`
}
