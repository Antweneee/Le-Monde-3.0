package sources

import (
	"github.com/gin-gonic/gin"
)

func DeleteBookmarks(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Bookmark deleted successfully",})
}