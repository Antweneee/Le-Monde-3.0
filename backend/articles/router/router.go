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
 
	r.POST("/articles", func(c *gin.Context) {
		src.AddIpfs(c)
		// src.AddDB(c, db)
	})

	r.GET("/articles/:id", func(c *gin.Context) {
		src.GetIPFS(c)
		// src.GetDB(c, db)
	})

	r.DELETE("/articles/:id", func(c *gin.Context) {
		src.DeleteIPFS(c)
		// src.DeleteDB(c, db)
	})

	// r.PUT("/articles/:id", func(c *gin.Context) {
	// 	src.PutIPFS(c, db)
	// 	src.PutDB(c, db)
	// })

	return r
}