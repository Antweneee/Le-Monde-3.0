package articles

import (
	"github.com/gin-gonic/gin"
	// "gorm.io/gorm"
	// a "backend/router/addArticle"
	// d "backend/router/deleteArticle"
	// g "backend/router/getArticle"
	// auth "backend/router/authentification"
	// dbService "backend/router/dbInteraction"
	// "backend/middlewares"
	// "github.com/gin-contrib/cors"
)

func ApplyArticlesRoutes(protected *gin.RouterGroup)  {
	

	protected.POST("/articles", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "POST ARTICLES",
		})
	})

	protected.GET("/articles/:id", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "GET ARTICLES",
		})
	})

	protected.DELETE("/articles/:id", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "DELETE ARTICLES",
		})
	})

	protected.PUT("/articles/:id", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "PUT ARTICLES",
		})
	})

}