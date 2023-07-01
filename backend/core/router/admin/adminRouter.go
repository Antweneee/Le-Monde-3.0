package articles

import (
	"github.com/gin-gonic/gin"
	"net/http"
	// "gorm.io/gorm"
	// a "backend/router/addArticle"
	// d "backend/router/deleteArticle"
	// g "backend/router/getArticle"
	// auth "backend/router/authentification"
	// dbService "backend/router/dbInteraction"
	// "backend/middlewares"
	// "github.com/gin-contrib/cors"
	"encoding/json"
	"bytes"
	"io/ioutil"
)

type RegisterInput struct {
	Email string `json:"email" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginInput struct {
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func makeHTTPRequest(c *gin.Context, method string, url string, requestBody interface{}) ([]byte, int, error) {
	jsonParams, err := json.Marshal(requestBody)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	request, err := http.NewRequest(method, url, bytes.NewBuffer(jsonParams))
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	defer response.Body.Close()

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return responseBody, response.StatusCode, nil
}


func ApplyAdminRoutes(protected *gin.RouterGroup, public *gin.RouterGroup)  {

	public.POST("/register", func(c *gin.Context) {
		var registerParams RegisterInput
	
		if err := c.ShouldBindJSON(&registerParams); err != nil {
			c.String(http.StatusBadRequest, "Invalid request body")
			return
		}
	
		responseBody, statusCode, err := makeHTTPRequest(c, http.MethodPost, "http://localhost:8081/register", registerParams)
		if err != nil {
			c.String(statusCode, "Error making the request")
			return
		}
	
		c.Data(statusCode, "application/json", responseBody)
	})
	
	public.POST("/login", func(c *gin.Context) {
		var loginParams LoginInput
	
		if err := c.ShouldBindJSON(&loginParams); err != nil {
			c.String(http.StatusBadRequest, "Invalid request body")
			return
		}
	
		responseBody, statusCode, err := makeHTTPRequest(c, http.MethodPost, "http://localhost:8081/login", loginParams)
		if err != nil {
			c.String(statusCode, "Error making the request")
			return
		}
	
		c.Data(statusCode, "application/json", responseBody)
	})
	
}