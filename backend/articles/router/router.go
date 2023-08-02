package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	src "main/sources"
)

func Router(db *gorm.DB) *gin.Engine {

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8080"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length", "Content-Type"},
		AllowCredentials: true,
	}))

	r.POST("/articles", func(c *gin.Context) {
		src.AddArticle(c, db)
	})

	r.GET("/articles", func(c *gin.Context) {
		src.GetAllArticles(c, db)
	})

	r.GET("/articles/:title", func(c *gin.Context) {
		src.GetArticle(c, db)
	})

	r.DELETE("/articles", func(c *gin.Context) {
		src.DeleteAllArticles(c, db)
	})

	r.DELETE("/articles/:title", func(c *gin.Context) {
		src.DeleteArticle(c, db)
	})

	r.PUT("/articles/:title", func(c *gin.Context) {
		src.EditArticle(c, db)
	})

	return r
}
