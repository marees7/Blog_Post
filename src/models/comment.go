package models

import (
	"time"
)

type Comment struct {
	ArticleID       uint      `gorm:"column:article_id;foreignkey:ArticleID" json:"articleId"`
	CommentID       uint      `gorm:"column:comment_id;primarykey" json:"commentId"`
	ParentCommentID uint      `gorm:"column:parent_comment_id" json:"parentCommentId"`
	Content         string    `gorm:"column:content;not null" json:"content"`
	Nickname        string    `gorm:"column:nickname;not null" json:"nickName"`
	CreationDate    time.Time `gorm:"column:creation_date;default:CURRENT_TIMESTAMP" json:"creationDate"`
	ParentComment   *Comment  `gorm:"foreignkey:ParentCommentID" json:"parentcomment"`
}

func (Comment) TableName() string {
	return "comment"
}
