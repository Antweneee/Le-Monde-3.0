package articles

import (
	"github.com/gin-gonic/gin"
	req "main/http"
	"net/http"
)

type RegisterInput struct {
	Email    string `json:"email" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func ApplyAdminRoutes(public *gin.RouterGroup) {

	public.POST("/register", func(c *gin.Context) {
		var registerParams RegisterInput

		if err := c.ShouldBindJSON(&registerParams); err != nil {
			c.String(http.StatusBadRequest, "Invalid request body")
			return
		}

		responseBody, statusCode, err := req.MakeHTTPRequest(c, http.MethodPost, "http://admin-lemonde3-0:8081/register", registerParams)
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

		responseBody, statusCode, err := req.MakeHTTPRequest(c, http.MethodPost, "http://admin-lemonde3-0:8081/login", loginParams)
		if err != nil {
			c.String(statusCode, "Error making the request")
			return
		}

		c.Data(statusCode, "application/json", responseBody)
	})

}
