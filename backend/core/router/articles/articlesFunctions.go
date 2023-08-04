package articles

import (
	"github.com/gin-gonic/gin"
	req "main/http"
	"net/http"
)

type ArticleInput struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

func AddArticle(c *gin.Context) {
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
}

func AddLike(c *gin.Context) {

	responseBody, statusCode, err := req.MakeHTTPRequest(c, http.MethodPost, "http://articles-lemonde3-0:8082/articles/"+c.Param("id")+"/likes", nil)
	if err != nil {
		c.String(statusCode, "Error making the request")
		return
	}
	c.Data(statusCode, "application/json", responseBody)
}

func GetAllArticles(c *gin.Context) {

	responseBody, statusCode, err := req.MakeHTTPRequest(c, http.MethodGet, "http://articles-lemonde3-0:8082/articles", nil)
	if err != nil {
		c.String(statusCode, "Error making the request")
		return
	}
	c.Data(statusCode, "application/json", responseBody)
}

func GetArticle(c *gin.Context) {

	responseBody, statusCode, err := req.MakeHTTPRequest(c, http.MethodGet, "http://articles-lemonde3-0:8082/articles/"+c.Param("id"), nil)
	if err != nil {
		c.String(statusCode, "Error making the request")
		return
	}
	c.Data(statusCode, "application/json", responseBody)
}

func GetMyArticles(c *gin.Context) {

	responseBody, statusCode, err := req.MakeHTTPRequest(c, http.MethodGet, "http://articles-lemonde3-0:8082/articles/me", nil)
	if err != nil {
		c.String(statusCode, "Error making the request")
		return
	}
	c.Data(statusCode, "application/json", responseBody)
}

func GetLikesInfo(c *gin.Context) {

	responseBody, statusCode, err := req.MakeHTTPRequest(c, http.MethodGet, "http://articles-lemonde3-0:8082/articles/"+c.Param("id")+"/likes", nil)
	if err != nil {
		c.String(statusCode, "Error making the request")
		return
	}
	c.Data(statusCode, "application/json", responseBody)
}

func EditArticle(c *gin.Context) {
	var articlesParams ArticleInput

	if err := c.ShouldBindJSON(&articlesParams); err != nil {
		c.String(http.StatusBadRequest, "Invalid request body")
		return
	}

	responseBody, statusCode, err := req.MakeHTTPRequest(c, http.MethodPut, "http://articles-lemonde3-0:8082/articles/"+c.Param("id"), articlesParams)
	if err != nil {
		c.String(statusCode, "Error making the request")
		return
	}
	c.Data(statusCode, "application/json", responseBody)
}

func DeleteAllArticles(c *gin.Context) {

	responseBody, statusCode, err := req.MakeHTTPRequest(c, http.MethodDelete, "http://articles-lemonde3-0:8082/articles/me", nil)
	if err != nil {
		c.String(statusCode, "Error making the request")
		return
	}
	c.Data(statusCode, "application/json", responseBody)
}

func DeleteArticle(c *gin.Context) {

	responseBody, statusCode, err := req.MakeHTTPRequest(c, http.MethodDelete, "http://articles-lemonde3-0:8082/articles/"+c.Param("id"), nil)
	if err != nil {
		c.String(statusCode, "Error making the request")
		return
	}
	c.Data(statusCode, "application/json", responseBody)
}

func RemoveLike(c *gin.Context) {

	responseBody, statusCode, err := req.MakeHTTPRequest(c, http.MethodDelete, "http://articles-lemonde3-0:8082/articles/"+c.Param("id")+"/likes", nil)
	if err != nil {
		c.String(statusCode, "Error making the request")
		return
	}
	c.Data(statusCode, "application/json", responseBody)
}
