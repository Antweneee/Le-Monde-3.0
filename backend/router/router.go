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

	r.POST("/database/addUser", func(c *gin.Context) {
		dbService.AddUser(c, db)
	})

	r.POST("/database/updateUser", func(c *gin.Context) {
		dbService.UpdateUser(c, db)
	})

	r.DELETE("/database/deleteUser", func(c *gin.Context) {
		dbService.DeleteUser(c, db)
	})

	// DB Article routes

	r.POST("/database/addArticle", func(c *gin.Context) {
		dbService.AddArticle(c, db)
	})

	r.POST("/database/updateArticle", func(c *gin.Context) {
		dbService.UpdateArticle(c, db)
	})

	r.DELETE("/database/deleteArticle", func(c *gin.Context) {
		dbService.DeleteArticle(c, db)
	})

	// DB Like routes

	r.POST("/database/addLike", func(c *gin.Context) {
		dbService.AddLike(c, db)
	})

	r.POST("/database/UpdateLike", func(c *gin.Context) {
		dbService.UpdateLike(c, db)
	})

	r.DELETE("/database/deleteLike", func(c *gin.Context) {
		dbService.DeleteLike(c, db)
	})

	// DB Bookmark routes

	r.POST("/database/addBookmark", func(c *gin.Context) {
		dbService.AddBookmark(c, db)
	})

	r.POST("/database/updateBookmark", func(c *gin.Context) {
		dbService.UpdateBookmark(c, db)
	})

	r.DELETE("/database/deleteBookmark", func(c *gin.Context) {
		dbService.DeleteBookmark(c, db)
	})

	// DB Topic routes

	r.POST("/database/addTopic", func(c *gin.Context) {
		dbService.AddTopic(c, db)
	})

	r.POST("/database/updateTopic", func(c *gin.Context) {
		dbService.UpdateTopic(c, db)
	})

	r.DELETE("/database/deleteTopic", func(c *gin.Context) {
		dbService.DeleteTopic(c, db)
	})
	return r
}