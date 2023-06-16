package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	a "backend/router/addArticle"
	d "backend/router/deleteArticle"
	g "backend/router/getArticle"
	auth "backend/router/authentification"
	dbService "backend/router/dbInteraction"
	"backend/middlewares"
	"github.com/gin-contrib/cors"
)

func Router(db *gorm.DB) *gin.Engine {
	
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: 	  []string{"http://localhost:3000"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length", "Content-Type"},
		AllowCredentials: true,
	   }))

	public := r.Group("/api")

	public.POST("/login", func(c *gin.Context) {
		auth.Login(c, db)
	})

	public.POST("/register", func(c *gin.Context) {
		auth.Register(c, db)
	})

	protected := r.Group("/api/admin")
	protected.Use(middlewares.JwtAuthMiddleware())

	protected.POST("/articles", a.Add)

	protected.GET("/articles/:cid", g.Get)

	protected.DELETE("/articles/:cid", d.Delete)

	// Account routes



	// DB User routes


	protected.GET("/database/User", func(c *gin.Context) {
		dbService.GetUser(c, db)
	})

	protected.DELETE("/database/User", func(c *gin.Context) {
		dbService.DeleteUser(c, db)
	})

	// DB Article routes

	protected.POST("/database/Article", func(c *gin.Context) {
		dbService.AddArticle(c, db)
	})

	protected.GET("/database/Article", func(c *gin.Context) {
		dbService.GetArticle(c, db)
	})

	protected.DELETE("/database/Article", func(c *gin.Context) {
		dbService.DeleteArticle(c, db)
	})

	// DB Like routes

	protected.POST("/database/Like", func(c *gin.Context) {
		dbService.AddLike(c, db)
	})


	// DB Bookmark routes

	protected.POST("/database/Bookmark", func(c *gin.Context) {
		dbService.AddBookmark(c, db)
	})


	// DB Topic routes

	protected.POST("/database/Topic", func(c *gin.Context) {
		dbService.AddTopic(c, db)
	})

	return r
}