package article

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

func ShowArticlesPage(c *gin.Context) {
	var articles []Article
	db := c.MustGet("db").(*gorm.DB)

	db.Find(&articles)

	c.HTML(http.StatusOK, "index.html", gin.H{"articles": articles})
}

func GetArticle(c *gin.Context) {
	var article Article
	id := c.Param("article_id")
	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("id = ?", id).First(&article).Error; err != nil {
		c.HTML(http.StatusBadRequest, "article.html", gin.H{"error": "Article not found"})
	} else {
		c.HTML(http.StatusOK, "article.html", gin.H{"article": article})
	}
}
