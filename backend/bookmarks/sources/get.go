package sources

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type Article struct {
	Id      int32
	UserId  int32
	Title   string
	Content string
	Likes   pq.Int32Array `gorm:"type:integer[]"`
}

func GetBookmark(c *gin.Context, db *gorm.DB) {
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
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": result.Error})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error})
		return
	} else if bookmark.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bookmark not found"})
		return
	}
	c.JSON(http.StatusOK, bookmark)
}

func GetAllBookmarks(c *gin.Context, db *gorm.DB) {
	bookmarks := new([]Bookmark)

	userId, err := getUserId(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	result := db.Where(Bookmark{UserId: userId}).Find(&bookmarks)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": result.Error})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error})
		return
	}
	c.JSON(http.StatusOK, bookmarks)
}

func GetAllArticlesBookmark(c *gin.Context, db *gorm.DB) {
	bookmark := new(Bookmark)
	var articlesBookmark []Article

	userId, err := getUserId(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	bookmarkId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"get": "bookmark id could not be retrieved"})
		return
	}

	result := db.Where(Bookmark{Id: int32(bookmarkId), UserId: userId}).Find(&bookmark)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": result.Error})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error})
		return
	} else if bookmark.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bookmark not found"})
		return
	}

	articlesBookmark, err = getArticlesFromBookmark(c, bookmark.Articles)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, articlesBookmark)
}
