package sources

import (
	"github.com/gin-gonic/gin"
)

func PutBookmarks(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Bookmark updated successfully",})
}