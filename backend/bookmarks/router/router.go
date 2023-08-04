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

	r.POST("/bookmarks", func(c *gin.Context) {
		src.AddBookmark(c, db)
	})

	r.POST("/bookmarks/:id/articles/:id-article", func(c *gin.Context) {
		src.AddArticleInBookmark(c, db)
	})

	r.GET("/bookmarks", func(c *gin.Context) {
		src.GetAllBookmarks(c, db)
	})

	r.GET("/bookmarks/:id", func(c *gin.Context) {
		src.GetBookmark(c, db)
	})

	r.GET("/bookmarks/:id/articles", func(c *gin.Context) {
		src.GetAllArticlesBookmark(c, db)
	})

	r.PUT("/bookmarks/:id", func(c *gin.Context) {
		src.EditBookmark(c, db)
	})

	r.DELETE("/bookmarks", func(c *gin.Context) {
		src.DeleteAllBookmarks(c, db)
	})

	r.DELETE("/bookmarks/:id", func(c *gin.Context) {
		src.DeleteBookmark(c, db)
	})

	r.DELETE("/bookmarks/:id/articles", func(c *gin.Context) {
		src.DeleteAllArticlesBookmark(c, db)
	})

	r.DELETE("/bookmarks/:id/articles/:id-article", func(c *gin.Context) {
		src.DeleteArticleBookmark(c, db)
	})

	return r
}
