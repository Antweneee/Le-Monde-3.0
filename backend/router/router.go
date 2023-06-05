package router

import (
	"github.com/gin-gonic/gin"
	a "backend/router/addArticle"
	d "backend/router/deleteArticle"
	g "backend/router/getArticle"
	db "backend/router/dbInteraction"
)

func Router() *gin.Engine {
	
	r := gin.Default()

	r.POST("/articles", a.Add)

	r.GET("/articles/:cid", g.Get)

	r.DELETE("/articles/:cid", d.Delete)

	// DB User routes

	r.POST("/database/addUser", db.AddUser)

	r.POST("/database/updateUser", db.UpdateUser)

	r.DELETE("/database/deleteUser", db.DeleteUser)

	// DB Article routes

	r.POST("/database/addArticle", db.AddArticle)

	r.POST("/database/updateArticle", db.UpdateArticle)

	r.DELETE("/database/deleteArticle", db.DeleteArticle)

	// DB Like routes

	r.POST("/database/addLike", db.AddLike)

	r.POST("/database/UpdateLike", db.UpdateLike)

	r.DELETE("/database/deleteLike", db.DeleteLike)

	// DB Bookmark routes

	r.POST("/database/addBookmark", db.AddBookmark)

	r.POST("/database/updateBookmark", db.UpdateBookmark)

	r.DELETE("/database/deleteBookmark", db.DeleteBookmark)

	// DB Topic routes

	r.POST("/database/addTopic", db.AddTopic)

	r.POST("/database/updateTopic", db.UpdateTopic)

	r.DELETE("/database/deleteTopic", db.DeleteTopic)
	return r
}