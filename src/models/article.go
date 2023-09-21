package models

import "time"

type Article struct {
	ArticleID    int       `gorm:"column:article_id;primaryKey;autoIncrement" json:"articleId"`
	Title        string    `gorm:"column:title;not null" json:"title"`
	Content      string    `gorm:"column:content;not null" json:"content"`
	Nickname     string    `gorm:"column:nickname;not null" json:"nickname"`
	CreationDate time.Time `gorm:"column:creation_date;type:timestamp;default:CURRENT_TIMESTAMP" json:"creationDate"`
}

func (Article) TableName() string {
	return "article"
}
