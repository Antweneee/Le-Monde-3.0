package sources

import (
	"github.com/gin-gonic/gin"
	"github.com/ipfs/go-ipfs-api"
	"gorm.io/gorm"
	"net/http"
)

func DeleteIPFS(c *gin.Context) {
	cid := c.Param("cid")

	sh := shell.NewShell("localhost:5001")

	err := sh.Unpin(cid)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{"message": "cid deleted successfully : " + cid})
}

func DeleteArticle(c *gin.Context, db *gorm.DB) {
	userId, err := getUserId(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	title := c.Param("title")

	if title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "title needed to delete an article"})
		return
	}

	result := db.Where(Article{UserId: userId, Title: title}).Delete(&Article{})
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error})
	} else {
		c.JSON(http.StatusOK, gin.H{"delete": "article has been deleted successfully"})
	}
}

func DeleteAllArticles(c *gin.Context, db *gorm.DB) {

	userId, err := getUserId(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	result := db.Where(Article{UserId: userId}).Delete(&Article{})
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error})
	} else {
		c.JSON(http.StatusOK, gin.H{"delete": "all articles have been successfully deleted"})
	}
}
