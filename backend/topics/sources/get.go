package sources

import (
	"github.com/gin-gonic/gin"
)

func GetTopics(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Topic retreived successfully",})
}