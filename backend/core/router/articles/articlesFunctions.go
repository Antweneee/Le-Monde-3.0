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

type Article struct {
	Id      int32
	UserId  int32
	Title   string
	Content string
	Likes   []int32
}

type LikesResponse struct {
	Amount   int
	Accounts []int32
}

type DeletedResponse struct {
	Delete string `json:"delete" example:"all articles have been successfully deleted"`
}

// AddArticle godoc
// @Schemes
// @Description Add an article
// @Tags articles
// @Accept json
// @Produce json
// @Param ArticleInput body ArticleInput true "Params to create an article"
// @Success 200 {object} Article
// @Failure      400  {object}  req.HTTPError
// @Failure      500  {object}  req.HTTPError
// @Router /articles [post]
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

// AddLike godoc
// @Schemes
// @Description Add a like to an article
// @Tags articles
// @Accept json
// @Produce json
// @Success 200 {object} Article
// @Failure      400  {object}  req.HTTPError
// @Failure      500  {object}  req.HTTPError
// @Router /articles/:id/likes [post]
func AddLike(c *gin.Context) {

	responseBody, statusCode, err := req.MakeHTTPRequest(c, http.MethodPost, "http://articles-lemonde3-0:8082/articles/"+c.Param("id")+"/likes", nil)
	if err != nil {
		c.String(statusCode, "Error making the request")
		return
	}
	c.Data(statusCode, "application/json", responseBody)
}

// GetAllArticles godoc
// @Schemes
// @Description Retrieve all articles
// @Tags articles
// @Accept json
// @Produce json
// @Success 200 {object} []Article
// @Failure      400  {object}  req.HTTPError
// @Failure      500  {object}  req.HTTPError
// @Router /articles [get]
func GetAllArticles(c *gin.Context) {

	responseBody, statusCode, err := req.MakeHTTPRequest(c, http.MethodGet, "http://articles-lemonde3-0:8082/articles", nil)
	if err != nil {
		c.String(statusCode, "Error making the request")
		return
	}
	c.Data(statusCode, "application/json", responseBody)
}

// GetArticle godoc
// @Schemes
// @Description Retrieve an article
// @Tags articles
// @Accept json
// @Produce json
// @Success 200 {object} Article
// @Failure      400  {object}  req.HTTPError
// @Failure      500  {object}  req.HTTPError
// @Router /articles/:id [get]
func GetArticle(c *gin.Context) {

	responseBody, statusCode, err := req.MakeHTTPRequest(c, http.MethodGet, "http://articles-lemonde3-0:8082/articles/"+c.Param("id"), nil)
	if err != nil {
		c.String(statusCode, "Error making the request")
		return
	}
	c.Data(statusCode, "application/json", responseBody)
}

// GetMyArticles godoc
// @Schemes
// @Description Retrieve connected user articles
// @Tags articles
// @Accept json
// @Produce json
// @Success 200 {object} []Article
// @Failure      400  {object}  req.HTTPError
// @Failure      500  {object}  req.HTTPError
// @Router /articles/me [get]
func GetMyArticles(c *gin.Context) {

	responseBody, statusCode, err := req.MakeHTTPRequest(c, http.MethodGet, "http://articles-lemonde3-0:8082/articles/me", nil)
	if err != nil {
		c.String(statusCode, "Error making the request")
		return
	}
	c.Data(statusCode, "application/json", responseBody)
}

// GetLikesInfo godoc
// @Schemes
// @Description Retrieve likes information related to an article
// @Tags articles
// @Accept json
// @Produce json
// @Success 200 {object} LikesResponse
// @Failure      400  {object}  req.HTTPError
// @Failure      500  {object}  req.HTTPError
// @Router /articles/:id/likes [get]
func GetLikesInfo(c *gin.Context) {

	responseBody, statusCode, err := req.MakeHTTPRequest(c, http.MethodGet, "http://articles-lemonde3-0:8082/articles/"+c.Param("id")+"/likes", nil)
	if err != nil {
		c.String(statusCode, "Error making the request")
		return
	}
	c.Data(statusCode, "application/json", responseBody)
}

// EditArticle godoc
// @Schemes
// @Description Edit an article
// @Tags articles
// @Accept json
// @Produce json
// @Param ArticleInput body ArticleInput true "Params to edit an article"
// @Success 200 {object} Article
// @Failure      400  {object}  req.HTTPError
// @Failure      500  {object}  req.HTTPError
// @Router /articles [put]
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

// DeleteAllArticles godoc
// @Schemes
// @Description Delete all the connected user articles
// @Tags articles
// @Accept json
// @Produce json
// @Success 200 {object} DeletedResponse
// @Failure      400  {object}  req.HTTPError
// @Failure      500  {object}  req.HTTPError
// @Router /articles [delete]
func DeleteAllArticles(c *gin.Context) {

	responseBody, statusCode, err := req.MakeHTTPRequest(c, http.MethodDelete, "http://articles-lemonde3-0:8082/articles/me", nil)
	if err != nil {
		c.String(statusCode, "Error making the request")
		return
	}
	c.Data(statusCode, "application/json", responseBody)
}

// DeleteArticle godoc
// @Schemes
// @Description Delete an article
// @Tags articles
// @Accept json
// @Produce json
// @Success 200 {object} DeletedResponse
// @Failure      400  {object}  req.HTTPError
// @Failure      500  {object}  req.HTTPError
// @Router /articles/:id [delete]
func DeleteArticle(c *gin.Context) {

	responseBody, statusCode, err := req.MakeHTTPRequest(c, http.MethodDelete, "http://articles-lemonde3-0:8082/articles/"+c.Param("id"), nil)
	if err != nil {
		c.String(statusCode, "Error making the request")
		return
	}
	c.Data(statusCode, "application/json", responseBody)
}

// RemoveLike godoc
// @Schemes
// @Description Remove a like from an article
// @Tags articles
// @Accept json
// @Produce json
// @Success 200 {object} Article
// @Failure      400  {object}  req.HTTPError
// @Failure      500  {object}  req.HTTPError
// @Router /articles/:id/likes [delete]
func RemoveLike(c *gin.Context) {

	responseBody, statusCode, err := req.MakeHTTPRequest(c, http.MethodDelete, "http://articles-lemonde3-0:8082/articles/"+c.Param("id")+"/likes", nil)
	if err != nil {
		c.String(statusCode, "Error making the request")
		return
	}
	c.Data(statusCode, "application/json", responseBody)
}
