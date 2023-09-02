package sources

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
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

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"put": "article id could not be retrieved"})
		return
	}

	result := db.Where(Article{UserId: userId, Id: int32(id)}).Find(&article)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error})
		return
	} else if article.Title == "" {
		c.JSON(http.StatusNotFound, gin.H{"error": "article not found or was not created by the current user"})
		return
	}
	article.Content = editedArticle.Content
	article.Title = editedArticle.Title

	db.Save(&article)
	c.JSON(http.StatusOK, article)
}
