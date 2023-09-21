package handler

import (
	h "blog-post-service/src/handler/v1"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type HttpHandler struct {
	Blog *h.BlogPostHandler
}

func GetHttpHander(db *gorm.DB) (*HttpHandler, error) {
	blogHandler, err := h.NewBlogPostHandler(db)
	if err != nil {
		return nil, err
	}

	return &HttpHandler{
		Blog: blogHandler,
	}, nil
}

func GetRoutes(r *gin.Engine, db *gorm.DB) (*HttpHandler, error) {
	handler, err := GetHttpHander(db)
	if err != nil {
		return nil, err
	}
	r.GET("/articles", handler.Blog.GetAllArticles)
	r.GET("/articles/:id", handler.Blog.GetArticle)
	r.POST("/article", handler.Blog.PostArticle)
	r.POST("/article/comment", handler.Blog.AddComment)
	r.GET("/comments/:article_id", handler.Blog.GetArticleComments)
	r.GET("/comments", handler.Blog.GetAllComments)
	r.GET("/comments/:article_id/:id", handler.Blog.GetComentOnComment)
	return nil, nil
}
