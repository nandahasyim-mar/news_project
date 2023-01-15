package models

type News struct {
	Id        int64  `gorm:"primary_key" json:"news_id"`
	TitleNews string `gorm:"type:varchar(300)" json:"title_news"`
	Author    string `gorm:"type:varchar(300)" json:"author"`
	Deskripsi string `gorm:"type:text" json:"description"`
}
