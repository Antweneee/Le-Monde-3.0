package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"github.com/gin-contrib/cors"
	src "main/sources"
)

func Router(db *gorm.DB) *gin.Engine {
	
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: 	  []string{"http://localhost:8080"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length", "Content-Type"},
		AllowCredentials: true,
	}))
 
	r.POST("/topics", func(c *gin.Context) {
		src.AddTopics(c)
	})

	r.GET("/topics/:id", func(c *gin.Context) {
		src.GetTopics(c)
	})

	r.DELETE("/topics/:id", func(c *gin.Context) {
		src.DeleteTopics(c)
	})

	r.PUT("/topics/:id", func(c *gin.Context) {
		src.PutTopics(c)
	})

	return r
}