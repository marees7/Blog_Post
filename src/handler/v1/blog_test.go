package v1_test

import (
	v1 "blog-post-service/src/handler/v1"
	"blog-post-service/src/models"
	"blog-post-service/src/utils/constants"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = "3306"
	user     = "root"
	password = ""
	dbName   = "blog_post_testdb"
)

func setupTestDB() (*gorm.DB, error) {
	dsn := user + "@(" + host + ":" + port + ")/" + dbName + "?parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.WithFields(log.Fields{
			"error":   err.Error(),
			"service": constants.ServiceName,
		}).Warn("failed to connect to database")
		return db, err

	}
	db.AutoMigrate(&models.Article{})
	db.AutoMigrate(&models.Comment{})
	return db, nil
}

func TestGetAllArticles(t *testing.T) {
	db, err := setupTestDB()
	assert.NoError(t, err)
	d, _ := db.DB()
	defer d.Close()
	r := gin.Default()
	handler, err := v1.NewBlogPostHandler(db)
	assert.NoError(t, err)
	r.GET("/articles", handler.GetAllArticles)
	req, err := http.NewRequest("GET", "/articles", nil)
	assert.NoError(t, err)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	response := []models.Article{}
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	if reflect.DeepEqual([]models.Article{}, response) {
		t.Errorf("Error on response")
	}
}

func JSONStringify(data interface{}) string {
	byteData, _ := json.Marshal(data)
	stringData := string(byteData)
	return stringData
}

func TestPostArticle(t *testing.T) {
	db, err := setupTestDB()
	assert.NoError(t, err)
	d, _ := db.DB()
	defer d.Close()
	r := gin.Default()
	handler, err := v1.NewBlogPostHandler(db)
	assert.NoError(t, err)
	r.POST("/article", handler.PostArticle)
	PostArticle := models.Article{
		Title:        "Sample Article",
		Content:      "This is a sample article content.",
		Nickname:     "JohnDoe",
		CreationDate: time.Now(),
	}
	strData := JSONStringify(PostArticle)
	req, err := http.NewRequest("POST", "/article", strings.NewReader(strData))
	assert.NoError(t, err)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	var response models.Article
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.NotNil(t, response.ArticleID)
	assert.Equal(t, PostArticle.Title, response.Title)
	assert.Equal(t, PostArticle.Content, response.Content)
	assert.Equal(t, PostArticle.Nickname, response.Nickname)
}

func TestGetArticle(t *testing.T) {
	db, err := setupTestDB()
	assert.NoError(t, err)
	d, _ := db.DB()
	defer d.Close()
	r := gin.Default()
	handler, err := v1.NewBlogPostHandler(db)
	assert.NoError(t, err)
	r.GET("/articles/:id", handler.GetArticle)
	req, err := http.NewRequest("GET", "/articles/1", nil)
	assert.NoError(t, err)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	response := models.Article{}
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	if reflect.DeepEqual(models.Article{}, response) {
		t.Errorf("Error on response")
	}
}

func TestAddComment(t *testing.T) {
	db, err := setupTestDB()
	assert.NoError(t, err)
	d, _ := db.DB()
	defer d.Close()
	r := gin.Default()
	handler, err := v1.NewBlogPostHandler(db)
	assert.NoError(t, err)
	r.POST("/article/comment", handler.AddComment)
	addComment := models.Comment{
		ArticleID:       1,
		ParentCommentID: 1,
		Content:         "This is a sample article content.",
		Nickname:        "JohnDoe",
		CreationDate:    time.Now(),
	}
	strData := JSONStringify(addComment)
	req, err := http.NewRequest("POST", "/article/comment", strings.NewReader(strData))
	assert.NoError(t, err)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	var response models.Comment
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.NotNil(t, response.ArticleID)
	assert.Equal(t, addComment.ArticleID, response.ArticleID)
	assert.Equal(t, addComment.ParentCommentID, response.ParentCommentID)
	assert.Equal(t, addComment.Content, response.Content)
	assert.Equal(t, addComment.Nickname, response.Nickname)
}

func TestGetArticleComments(t *testing.T) {
	db, err := setupTestDB()
	assert.NoError(t, err)
	d, _ := db.DB()
	defer d.Close()
	r := gin.Default()
	handler, err := v1.NewBlogPostHandler(db)
	assert.NoError(t, err)
	r.GET("/comments/:article_id", handler.GetArticleComments)
	req, err := http.NewRequest("GET", "/comments/1", nil)
	assert.NoError(t, err)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	response := []models.Comment{}
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	if reflect.DeepEqual([]models.Comment{}, response) {
		t.Errorf("Error on response")
	}
}

func TestGetAllComments(t *testing.T) {
	db, err := setupTestDB()
	assert.NoError(t, err)
	d, _ := db.DB()
	defer d.Close()
	r := gin.Default()
	handler, err := v1.NewBlogPostHandler(db)
	assert.NoError(t, err)
	r.GET("/comments", handler.GetAllComments)
	req, err := http.NewRequest("GET", "/comments", nil)
	assert.NoError(t, err)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	response := []models.Comment{}
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	if reflect.DeepEqual([]models.Comment{}, response) {
		t.Errorf("Error on response")
	}
}

func TestGetComentOnComment(t *testing.T) {
	db, err := setupTestDB()
	assert.NoError(t, err)
	d, _ := db.DB()
	defer d.Close()
	r := gin.Default()
	handler, err := v1.NewBlogPostHandler(db)
	assert.NoError(t, err)
	r.GET("/comments/:article_id/:id", handler.GetComentOnComment)
	req, err := http.NewRequest("GET", "/comments/1/1", nil)
	assert.NoError(t, err)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	response := []models.Comment{}
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	if reflect.DeepEqual([]models.Comment{}, response) {
		t.Errorf("Error on response")
	}
}
