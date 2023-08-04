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

	r.POST("/articles/:id/likes", func(c *gin.Context) {
		src.AddLike(c, db)
	})

	r.GET("/articles", func(c *gin.Context) {
		src.GetAllArticles(c, db)
	})

	r.GET("/articles/:id", func(c *gin.Context) {
		src.GetArticle(c, db)
	})

	r.GET("/articles/me", func(c *gin.Context) {
		src.GetMyArticles(c, db)
	})

	r.GET("/articles/:id/likes", func(c *gin.Context) {
		src.GetLikesInfo(c, db)
	})

	r.PUT("/articles/:id", func(c *gin.Context) {
		src.EditArticle(c, db)
	})

	r.DELETE("/articles", func(c *gin.Context) {
		src.DeleteAllArticles(c, db)
	})

	r.DELETE("/articles/:id", func(c *gin.Context) {
		src.DeleteArticle(c, db)
	})

	r.DELETE("/articles/:id/likes", func(c *gin.Context) {
		src.RemoveLike(c, db)
	})

	return r
}
