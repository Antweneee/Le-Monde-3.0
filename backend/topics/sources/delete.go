package sources

import (
	"github.com/gin-gonic/gin"
)

func DeleteTopics(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Topic deleted successfully",})
}