package articles

import (
	"github.com/gin-gonic/gin"
	req "main/http"
	"net/http"
)

type BookmarkInput struct {
	Title string `json:"title" binding:"required"`
}

func AddBookmark(c *gin.Context) {
	var bookmarksParams BookmarkInput

	if err := c.ShouldBindJSON(&bookmarksParams); err != nil {
		c.String(http.StatusBadRequest, "Invalid request body")
		return
	}

	responseBody, statusCode, err := req.MakeHTTPRequest(c, http.MethodPost, "http://bookmarks-lemonde3-0:8084/bookmarks", bookmarksParams)
	if err != nil {
		c.String(statusCode, "Error making the request")
		return
	}

	c.Data(statusCode, "application/json", responseBody)
}

func AddArticleInBookmark(c *gin.Context) {

	responseBody, statusCode, err := req.MakeHTTPRequest(c, http.MethodPost, "http://bookmarks-lemonde3-0:8084/bookmarks/"+c.Param("id")+"/articles/"+c.Param("id-article"), nil)
	if err != nil {
		c.String(statusCode, "Error making the request")
		return
	}

	c.Data(statusCode, "application/json", responseBody)
}

func GetAllBookmarks(c *gin.Context) {

	responseBody, statusCode, err := req.MakeHTTPRequest(c, http.MethodGet, "http://bookmarks-lemonde3-0:8084/bookmarks", nil)
	if err != nil {
		c.String(statusCode, "Error making the request")
		return
	}

	c.Data(statusCode, "application/json", responseBody)
}

func GetBookmark(c *gin.Context) {

	responseBody, statusCode, err := req.MakeHTTPRequest(c, http.MethodGet, "http://bookmarks-lemonde3-0:8084/bookmarks/"+c.Param("id"), nil)
	if err != nil {
		c.String(statusCode, "Error making the request")
		return
	}

	c.Data(statusCode, "application/json", responseBody)
}

func GetAllArticlesBookmark(c *gin.Context) {

	responseBody, statusCode, err := req.MakeHTTPRequest(c, http.MethodGet, "http://bookmarks-lemonde3-0:8084/bookmarks/"+c.Param("id")+"/articles", nil)
	if err != nil {
		c.String(statusCode, "Error making the request")
		return
	}

	c.Data(statusCode, "application/json", responseBody)
}

func EditBookmark(c *gin.Context) {
	var bookmarksParams BookmarkInput

	if err := c.ShouldBindJSON(&bookmarksParams); err != nil {
		c.String(http.StatusBadRequest, "Invalid request body")
		return
	}

	responseBody, statusCode, err := req.MakeHTTPRequest(c, http.MethodPut, "http://bookmarks-lemonde3-0:8084/bookmarks/"+c.Param("id"), bookmarksParams)
	if err != nil {
		c.String(statusCode, "Error making the request")
		return
	}

	c.Data(statusCode, "application/json", responseBody)
}

func DeleteAllBookmarks(c *gin.Context) {

	responseBody, statusCode, err := req.MakeHTTPRequest(c, http.MethodDelete, "http://bookmarks-lemonde3-0:8084/bookmarks", nil)
	if err != nil {
		c.String(statusCode, "Error making the request")
		return
	}

	c.Data(statusCode, "application/json", responseBody)
}

func DeleteBookmark(c *gin.Context) {

	responseBody, statusCode, err := req.MakeHTTPRequest(c, http.MethodDelete, "http://bookmarks-lemonde3-0:8084/bookmarks/"+c.Param("id"), nil)
	if err != nil {
		c.String(statusCode, "Error making the request")
		return
	}

	c.Data(statusCode, "application/json", responseBody)
}

func DeleteAllArticlesBookmark(c *gin.Context) {

	responseBody, statusCode, err := req.MakeHTTPRequest(c, http.MethodDelete, "http://bookmarks-lemonde3-0:8084/bookmarks/"+c.Param("id")+"/articles", nil)
	if err != nil {
		c.String(statusCode, "Error making the request")
		return
	}

	c.Data(statusCode, "application/json", responseBody)
}

func DeleteArticleBookmark(c *gin.Context) {

	responseBody, statusCode, err := req.MakeHTTPRequest(c, http.MethodDelete, "http://bookmarks-lemonde3-0:8084/bookmarks/"+c.Param("id")+"/"+c.Param("id-article"), nil)
	if err != nil {
		c.String(statusCode, "Error making the request")
		return
	}

	c.Data(statusCode, "application/json", responseBody)
}
