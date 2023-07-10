package router

import (
	"github.com/gin-gonic/gin"
	mw "main/middlewares"
	"github.com/gin-contrib/cors"
	art "main/router/articles"
	adm "main/router/admin"
)

func Router() *gin.Engine {
	
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: 	  []string{"http://localhost:3000"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length", "Content-Type"},
		AllowCredentials: true,
	}))

	public := r.Group("/")
	protected := r.Group("/")

	protected.Use(mw.JwtAuthMiddleware())

	adm.ApplyAdminRoutes(public)
	art.ApplyArticlesRoutes(protected)

	return r
}