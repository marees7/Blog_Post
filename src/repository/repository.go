package repository

import (
	m "blog-post-service/src/models"

	"gorm.io/gorm"
)

type BlogPost interface {
	PostArticleRepo(req *m.Article, db *gorm.DB) (*m.Article, error)
	GetAllArticlesRepo(db *gorm.DB) ([]m.Article, error)
	GetArticleRepo(id string, db *gorm.DB) (m.Article, error)
	AddCommentRepo(req *m.Comment, db *gorm.DB) (*m.Comment, error)
	GetArticleCommentsRepo(articleId string, db *gorm.DB) ([]m.Comment, error)
	GetAllCommentsRepo(db *gorm.DB) ([]m.Comment, error)
	GetCommentOnCommentsRepo(id string, articleId string, db *gorm.DB) ([]m.Comment, error)
}
