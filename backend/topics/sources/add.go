package sources

import (
	"github.com/gin-gonic/gin"
)

func AddTopics(c *gin.Context) {
	c.JSON(201, gin.H{"message": "Topic added successfully",})
}