package router

import (
	"github.com/gin-gonic/gin"
	a "backend/router/addArticle"
	d "backend/router/deleteArticle"
	g "backend/router/getArticle"

)

func Router() *gin.Engine {
	
	r := gin.Default()

	r.POST("/articles", a.Add)

	r.GET("/articles/:cid", g.Get)

	r.DELETE("/articles/:cid", d.Delete)

	return r
}