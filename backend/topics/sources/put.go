package sources

import (
	"github.com/gin-gonic/gin"
)

func PutTopics(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Topic updated successfully",})
}