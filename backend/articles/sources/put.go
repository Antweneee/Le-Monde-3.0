package sources

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type EditedArticle struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func EditArticle(c *gin.Context, db *gorm.DB) {
	article := new(Article)

	userId, err := getUserId(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	editedArticle := EditedArticle{}
	if err := c.ShouldBindJSON(&editedArticle); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	title := c.Param("title")

	if title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "title needed to retrieve an article"})
		return
	}

	result := db.Where(Article{UserId: userId, Title: title}).Find(&article)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": result.Error})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error})
		return
	} else if article.Id == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "article not found"})
		return
	}
	article.Content = editedArticle.Content
	article.Title = editedArticle.Title

	db.Save(&article)

	c.JSON(http.StatusOK, article)
}
