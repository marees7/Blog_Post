package mysql

import (
	"blog-post-service/src/models"
	"blog-post-service/src/utils/constants"
	"os"

	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"

	"gorm.io/gorm"
)

func DSN() string {
	dbHost := os.Getenv("DB_HOST")
	dbUsername := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASS")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dsn := dbUsername + "@(" + dbHost + ":" + dbPort + ")/" + dbName + "?parseTime=true"
	if dbPassword != "" {
		dsn = dbUsername + ":" + dbPassword + "@(" + dbHost + ":" + dbPort + ")"
	}
	return dsn
}

func NewDb(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.WithFields(log.Fields{
			"error":   err.Error(),
			"service": constants.ServiceName,
		}).Warn("failed to connect to database")
		return db, err

	}
	if err := db.AutoMigrate(&models.Article{}, &models.Comment{}); err != nil {
		log.WithFields(log.Fields{
			"error":   err.Error(),
			"service": constants.ServiceName,
		}).Warn("failed to migrate the database")
		return db, err

	}
	return db, nil
}
