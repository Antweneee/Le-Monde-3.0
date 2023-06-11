package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	a "backend/router/addArticle"
	d "backend/router/deleteArticle"
	g "backend/router/getArticle"
	dbService "backend/router/dbInteraction"
)

func Router(db *gorm.DB) *gin.Engine {
	
	r := gin.Default()

	r.POST("/articles", a.Add)

	r.GET("/articles/:cid", g.Get)

	r.DELETE("/articles/:cid", d.Delete)

	// DB User routes

	r.POST("/database/User", func(c *gin.Context) {
		dbService.AddUser(c, db)
	})

	r.GET("/database/User", func(c *gin.Context) {
		dbService.GetUser(c, db)
	})

	r.DELETE("/database/User", func(c *gin.Context) {
		dbService.DeleteUser(c, db)
	})

	// DB Article routes

	r.POST("/database/Article", func(c *gin.Context) {
		dbService.AddArticle(c, db)
	})

	r.GET("/database/Article", func(c *gin.Context) {
		dbService.GetArticle(c, db)
	})

	r.DELETE("/database/Article", func(c *gin.Context) {
		dbService.DeleteArticle(c, db)
	})

	// DB Like routes

	r.POST("/database/Like", func(c *gin.Context) {
		dbService.AddLike(c, db)
	})


	// DB Bookmark routes

	r.POST("/database/Bookmark", func(c *gin.Context) {
		dbService.AddBookmark(c, db)
	})


	// DB Topic routes

	r.POST("/database/Topic", func(c *gin.Context) {
		dbService.AddTopic(c, db)
	})

	return r
}