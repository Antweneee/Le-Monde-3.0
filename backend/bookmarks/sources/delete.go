package sources

import "C"
import (
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func DeleteBookmark(c *gin.Context, db *gorm.DB) {
	userId, err := getUserId(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	bookmarkId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"get": "bookmark id could not be retrieved"})
		return
	}

	result := db.Where(Bookmark{Id: int32(bookmarkId), UserId: userId}).Delete(&Bookmark{})
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error})
		return
	} else if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "article not found or was not created by the current user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"delete": "bookmark has been deleted successfully"})
}

func DeleteAllBookmarks(c *gin.Context, db *gorm.DB) {
	userId, err := getUserId(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	result := db.Where(Bookmark{UserId: userId}).Delete(&Bookmark{})
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error})
	}
	c.JSON(http.StatusOK, gin.H{"delete": "all bookmarks have been successfully deleted"})
}

func DeleteAllArticlesBookmark(c *gin.Context, db *gorm.DB) {
	bookmark := new(Bookmark)

	userId, err := getUserId(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	bookmarkId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"get": "bookmark id could not be retrieved"})
		return
	}

	result := db.Where(Bookmark{Id: int32(bookmarkId), UserId: userId}).Find(&bookmark)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error})
		return
	} else if bookmark.Title == "" {
		c.JSON(http.StatusNotFound, gin.H{"error": "bookmark was not found"})
		return
	}
	bookmark.Articles = pq.Int32Array{}
	db.Save(&bookmark)
	c.JSON(http.StatusOK, bookmark)
}

func rmIfNotPresent(slice pq.Int32Array, key int32) pq.Int32Array {
	found := false
	for _, value := range slice {
		if value == key {
			found = true
			break
		}
	}

	if !found {
		return slice
	}

	result := []int32{}
	for _, value := range slice {
		if value != key {
			result = append(result, value)
		}
	}
	return result
}

func DeleteArticleBookmark(c *gin.Context, db *gorm.DB) {
	bookmark := new(Bookmark)

	userId, err := getUserId(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	bookmarkId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"get": "bookmark id could not be retrieved"})
		return
	}
	articleId, err := strconv.ParseInt(c.Param("id-article"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"get": "article id could not be retrieved"})
		return
	}

	result := db.Where(Bookmark{Id: int32(bookmarkId), UserId: userId}).Find(&bookmark)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error})
		return
	} else if bookmark.Title == "" {
		c.JSON(http.StatusNotFound, gin.H{"error": "bookmark was not found"})
		return
	}
	bookmark.Articles = rmIfNotPresent(bookmark.Articles, int32(articleId))
	db.Save(&bookmark)
	c.JSON(http.StatusOK, bookmark)
}
