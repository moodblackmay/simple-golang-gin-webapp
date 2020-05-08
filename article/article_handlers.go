package article

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var articles = []Article{
	Article{ID: 1, Title: "Article 1", Content: "Article 1 body"},
	Article{ID: 2, Title: "Article 2", Content: "Article 2 body"},
}

func ShowArticlePage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{"title": "Test", "articles": articles})
}

func GetArticle(c *gin.Context) {
	id := c.Param("article_id")
	c.HTML(http.StatusOK, "article.html", gin.H{"data": id})
}
