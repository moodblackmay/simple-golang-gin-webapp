package core

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/viper"
	"practice-golang-gin-webapp/article"
)

func SetupDb() *gorm.DB {
	viper.AutomaticEnv()

	viperUser := viper.Get("POSTGRES_USER").(string)
	viperPassword := viper.Get("POSTGRES_PASSWORD").(string)
	viperHost := viper.Get("POSTGRES_HOST").(string)
	viperPort := viper.Get("POSTGRES_PORT").(string)
	viperDb := viper.Get("POSTGRES_DB").(string)

	dbArgs := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s",
		viperHost, viperPort, viperUser, viperDb, viperPassword)

	db, err := gorm.Open("postgres", dbArgs)
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&article.Article{})

	return db
}

func SetupRouters(router *gin.Engine) {
	router.GET("/", article.ShowArticlePage)
	router.GET("/article/:article_id", article.GetArticle)
}
