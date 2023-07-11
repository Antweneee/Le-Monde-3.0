package sources

import (
	"github.com/gin-gonic/gin"
)

func AddBookmarks(c *gin.Context) {
	c.JSON(201, gin.H{"message": "Bookmarks added successfully",})
}