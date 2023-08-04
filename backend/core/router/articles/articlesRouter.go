package articles

import (
	"github.com/gin-gonic/gin"
)

func ApplyArticlesRoutes(protected *gin.RouterGroup) {

	protected.POST("/articles", AddArticle)
	protected.POST("/articles/:id/likes", AddLike)

	protected.GET("/articles", GetAllArticles)
	protected.GET("/articles/:id", GetArticle)
	protected.GET("/articles/me", GetMyArticles)
	protected.GET("/articles/:id/likes", GetLikesInfo)

	protected.PUT("/articles/:id", EditArticle)

	protected.DELETE("/articles", DeleteAllArticles)
	protected.DELETE("/articles/id", DeleteArticle)
	protected.DELETE("/articles/:id/likes", RemoveLike)
}
