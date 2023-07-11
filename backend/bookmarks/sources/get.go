package sources

import (
	"github.com/gin-gonic/gin"
)

func GetBookmarks(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Bookmark retreived successfully",})
}