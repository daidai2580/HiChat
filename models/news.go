package models

type News struct {
	Model
	Title       string `json:"title" json:"title,omitempty" binding:"required"`       //标题
	Content     string `json:"content" json:"content,omitempty" binding:"required"`   //内容
	AuthorId    int64  `json:"authorId" json:"authorId,omitempty" binding:"required"` //作者ID
	ReviewCount int    `json:"reviewCount,omitempty"`                                 //阅读数量

}

func (m *News) NewsTable() string {
	return "news"
}
