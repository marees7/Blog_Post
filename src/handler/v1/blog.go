package v1

import (
	"blog-post-service/src/models"
	"blog-post-service/src/repository"
	r "blog-post-service/src/repository/v1"
	"blog-post-service/src/utils/constants"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type BlogPostHandler struct {
	repo repository.BlogPost
	db   *gorm.DB
}

func NewBlogPostHandler(db *gorm.DB) (*BlogPostHandler, error) {
	return &BlogPostHandler{
		repo: r.NewBlogPostRepo(),
		db:   db,
	}, nil
}

func (b *BlogPostHandler) GetAllArticles(c *gin.Context) {
	log.WithFields(log.Fields{"service": constants.ServiceName, "ended_at": time.Now()}).Info("GetAllArticles - started")
	res, err := b.repo.GetAllArticlesRepo(b.db)
	if err != nil {
		log.WithFields(log.Fields{"service": constants.ServiceName, "error": err}).Error("failed to fetch the all articles")
		c.JSON(http.StatusBadRequest, map[string]interface{}{"err": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (b *BlogPostHandler) GetArticle(c *gin.Context) {
	log.WithFields(log.Fields{"service": constants.ServiceName, "ended_at": time.Now()}).Info("GetAllArticles - started")
	id := c.Param("id")
	res, err := b.repo.GetArticleRepo(id, b.db)
	if err != nil {
		log.WithFields(log.Fields{"service": constants.ServiceName, "error": err}).Error("failed to fetch the all articles")
		c.JSON(http.StatusBadRequest, map[string]interface{}{"err": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (b *BlogPostHandler) PostArticle(c *gin.Context) {
	log.WithFields(log.Fields{"service": constants.ServiceName, "ended_at": time.Now()}).Info("PostAarticle - started")
	article := &models.Article{}
	articleData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.WithFields(log.Fields{"service": constants.ServiceName, "error": err}).Error("failed to read request body")
		c.JSON(http.StatusBadRequest, map[string]interface{}{"err": err.Error()})
		return
	}
	err = json.Unmarshal(articleData, &article)
	if err != nil {
		log.WithFields(log.Fields{"service": constants.ServiceName, "error": err}).Error("failed to unmarshal request body")
		c.JSON(http.StatusBadRequest, map[string]interface{}{"err": err.Error()})
		return
	}
	res, err := b.repo.PostArticleRepo(article, b.db)
	if err != nil {
		log.WithFields(log.Fields{"service": constants.ServiceName, "error": err}).Error("failed to post the article")
		c.JSON(http.StatusBadRequest, map[string]interface{}{"err": err.Error()})
		return
	}
	result := map[string]interface{}{
		"articleId": res.ArticleID,
		"title":     res.Title,
		"content":   res.Content,
		"nickname":  res.Nickname,
	}
	c.JSON(http.StatusOK, result)
}

func (b *BlogPostHandler) AddComment(c *gin.Context) {
	log.WithFields(log.Fields{"service": constants.ServiceName, "ended_at": time.Now()}).Info("AddComment - started")
	comment := &models.Comment{}
	commentDetails, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.WithFields(log.Fields{"service": constants.ServiceName, "error": err}).Error("failed to read request body")
		c.JSON(http.StatusBadRequest, map[string]interface{}{"err": err.Error()})
		return
	}
	err = json.Unmarshal(commentDetails, &comment)
	if err != nil {
		log.WithFields(log.Fields{"service": constants.ServiceName, "error": err}).Error("failed to unmarshal request body")
		c.JSON(http.StatusBadRequest, map[string]interface{}{"err": err.Error()})
		return
	}
	res, err := b.repo.AddCommentRepo(comment, b.db)
	if err != nil {
		log.WithFields(log.Fields{"service": constants.ServiceName, "error": err}).Error("failed to add the comment")
		c.JSON(http.StatusBadRequest, map[string]interface{}{"err": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (b *BlogPostHandler) GetArticleComments(c *gin.Context) {
	log.WithFields(log.Fields{"service": constants.ServiceName, "ended_at": time.Now()}).Info("GetArticleComments - started")
	id := c.Param("article_id")
	res, err := b.repo.GetArticleCommentsRepo(id, b.db)
	if err != nil {
		log.WithFields(log.Fields{"service": constants.ServiceName, "error": err}).Error("failed to fetch the article comments")
		c.JSON(http.StatusBadRequest, map[string]interface{}{"err": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (b *BlogPostHandler) GetAllComments(c *gin.Context) {
	log.WithFields(log.Fields{"service": constants.ServiceName, "ended_at": time.Now()}).Info("GetAllArticles - started")
	res, err := b.repo.GetAllCommentsRepo(b.db)
	if err != nil {
		log.WithFields(log.Fields{"service": constants.ServiceName, "error": err}).Error("failed to fetch the all comments")
		c.JSON(http.StatusBadRequest, map[string]interface{}{"err": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (b *BlogPostHandler) GetComentOnComment(c *gin.Context) {
	log.WithFields(log.Fields{"service": constants.ServiceName, "ended_at": time.Now()}).Info("GetComentOnComment - started")
	parentCommentId := c.Param("id")
	articleId := c.Param("article_id")
	res, err := b.repo.GetCommentOnCommentsRepo(parentCommentId, articleId, b.db)
	if err != nil {
		log.WithFields(log.Fields{"service": constants.ServiceName, "error": err}).Error("failed to fetch the Get ComentOnComment")
		c.JSON(http.StatusBadRequest, map[string]interface{}{"err": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
