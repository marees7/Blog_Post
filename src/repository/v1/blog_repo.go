package repository_v1

import (
	m "blog-post-service/src/models"
	r "blog-post-service/src/repository"
	"blog-post-service/src/utils/constants"
	"errors"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type BlogRepo struct {
}

// Returns new task repository
func NewBlogPostRepo() r.BlogPost {
	return &BlogRepo{}
}

func (BlogRepo) GetAllArticlesRepo(db *gorm.DB) ([]m.Article, error) {
	log.WithFields(log.Fields{"service": constants.ServiceName, "ended_at": time.Now()}).Info("GetAllArticlesRepo - started")
	articles := []m.Article{}
	err := db.Find(&articles).Error
	fmt.Println("error ::", err)
	if err != nil {
		err = errors.New("Got error while fetch all the article " + err.Error())
		log.WithFields(log.Fields{
			"service": constants.ServiceName,
			"err":     err,
		}).Error(err.Error())
		return nil, err
	}
	return articles, nil
}

func (BlogRepo) PostArticleRepo(req *m.Article, db *gorm.DB) (*m.Article, error) {
	log.WithFields(log.Fields{"service": constants.ServiceName, "ended_at": time.Now()}).Info("PostAarticleRepo - started")
	err := db.Create(&req).Error
	fmt.Println("error ::", err)
	if err != nil {
		err = errors.New("Got error while post the article " + err.Error())
		log.WithFields(log.Fields{
			"service": constants.ServiceName,
			"err":     err,
		}).Error(err.Error())
		return nil, err
	}
	return req, nil
}

func (BlogRepo) GetArticleRepo(id string, db *gorm.DB) (m.Article, error) {
	log.WithFields(log.Fields{"service": constants.ServiceName, "ended_at": time.Now()}).Info("GetArticleRepo - started")
	articles := m.Article{}
	err := db.Where("article_id=?", id).Find(&articles).Error
	fmt.Println("error ::", err)
	if err != nil {
		err = errors.New("Got error while fetch an article " + err.Error())
		log.WithFields(log.Fields{
			"service": constants.ServiceName,
			"err":     err,
		}).Error(err.Error())
		return articles, err
	}
	return articles, nil
}

func (BlogRepo) AddCommentRepo(req *m.Comment, db *gorm.DB) (*m.Comment, error) {
	log.WithFields(log.Fields{"service": constants.ServiceName, "ended_at": time.Now()}).Info("AddCommentRepo - started")
	err := db.Debug().Create(&req).Error
	fmt.Println("error ::", err)
	if err != nil {
		err = errors.New("Got error while post the article " + err.Error())
		log.WithFields(log.Fields{
			"service": constants.ServiceName,
			"err":     err,
		}).Error(err.Error())
		return nil, err
	}
	return req, nil
}

func (BlogRepo) GetArticleCommentsRepo(articleId string, db *gorm.DB) ([]m.Comment, error) {
	log.WithFields(log.Fields{"service": constants.ServiceName, "ended_at": time.Now()}).Info("GetArticleCommentsRepo - started")
	comments := []m.Comment{}
	err := db.Where("article_id=?", articleId).Find(&comments).Error
	fmt.Println("error ::", err)
	if err != nil {
		err = errors.New("Got error while post the article " + err.Error())
		log.WithFields(log.Fields{
			"service": constants.ServiceName,
			"err":     err,
		}).Error(err.Error())
		return nil, err
	}
	return comments, nil
}

func (BlogRepo) GetAllCommentsRepo(db *gorm.DB) ([]m.Comment, error) {
	log.WithFields(log.Fields{"service": constants.ServiceName, "ended_at": time.Now()}).Info("GetAllCommentsRepo - started")
	comments := []m.Comment{}
	err := db.Find(&comments).Error
	fmt.Println("error ::", err)
	if err != nil {
		err = errors.New("Got error while post the article " + err.Error())
		log.WithFields(log.Fields{
			"service": constants.ServiceName,
			"err":     err,
		}).Error(err.Error())
		return nil, err
	}
	return comments, nil
}

func (BlogRepo) GetCommentOnCommentsRepo(id string, articleId string, db *gorm.DB) ([]m.Comment, error) {
	log.WithFields(log.Fields{"service": constants.ServiceName, "ended_at": time.Now()}).Info("GetCommentOnCommentsRepo - started")
	comments := []m.Comment{}
	err := db.Debug().Where("parent_comment_id=? AND article_id = ?", id, articleId).Find(&comments).Error
	fmt.Println("error ::", err)
	if err != nil {
		err = errors.New("Got error while post the article " + err.Error())
		log.WithFields(log.Fields{
			"service": constants.ServiceName,
			"err":     err,
		}).Error(err.Error())
		return nil, err
	}
	return comments, nil
}
