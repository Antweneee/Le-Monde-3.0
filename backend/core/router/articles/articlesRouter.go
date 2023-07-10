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
	req "main/http"
	"net/http"
)

type ArticleInput struct {
	Email string `json:"email" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func ApplyArticlesRoutes(protected *gin.RouterGroup)  {
	

	protected.POST("/articles", func(c *gin.Context) {
		var articlesParams ArticleInput
	
		if err := c.ShouldBindJSON(&articlesParams); err != nil {
			c.String(http.StatusBadRequest, "Invalid request body")
			return
		}
	
		responseBody, statusCode, err := req.MakeHTTPRequest(c, http.MethodPost, "http://articles-lemonde3-0:8082/articles", articlesParams)
		if err != nil {
			c.String(statusCode, "Error making the request")
			return
		}
	
		c.Data(statusCode, "application/json", responseBody)
	})

	protected.GET("/articles/:id", func(c *gin.Context) {
		var articlesParams ArticleInput
	
		if err := c.ShouldBindJSON(&articlesParams); err != nil {
			c.String(http.StatusBadRequest, "Invalid request body")
			return
		}
	
		responseBody, statusCode, err := req.MakeHTTPRequest(c, http.MethodPost, "http://articles:8082/articles", articlesParams)
		if err != nil {
			c.String(statusCode, "Error making the request")
			return
		}
	
		c.Data(statusCode, "application/json", responseBody)
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