package sources

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type EditedBookmark struct {
	Title string `json:"title"`
}

func EditBookmark(c *gin.Context, db *gorm.DB) {
	bookmark := new(Bookmark)

	userId, err := getUserId(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	editedBookmark := EditedBookmark{}
	if err := c.ShouldBindJSON(&editedBookmark); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
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
	}
	bookmark.Title = editedBookmark.Title

	db.Save(&bookmark)
	c.JSON(http.StatusOK, bookmark)
}
