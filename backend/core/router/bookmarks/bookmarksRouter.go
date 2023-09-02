package articles

import (
	"github.com/gin-gonic/gin"
)

func ApplyBookmarksRoutes(protected *gin.RouterGroup) {

	protected.POST("/bookmarks", AddBookmark)
	protected.POST("/bookmarks/:id/articles/:id-article", AddArticleInBookmark)

	protected.GET("/bookmarks", GetAllBookmarks)
	protected.GET("/bookmarks/:id", GetBookmark)
	protected.GET("/bookmarks/:id/articles", GetAllArticlesBookmark)

	protected.PUT("/bookmarks/:id", EditBookmark)

	protected.DELETE("/bookmarks", DeleteAllBookmarks)
	protected.DELETE("/bookmarks/:id", DeleteBookmark)
	protected.DELETE("/bookmarks/:id/articles", DeleteAllArticlesBookmark)
	protected.DELETE("/bookmarks/:id/articles/:id-article", DeleteArticleBookmark)
}
