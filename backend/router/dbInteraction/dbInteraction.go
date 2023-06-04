package dbInteraction

import (
    "github.com/gin-gonic/gin"
	"backend/database"
	"net/http"
)

func AddUser(c *gin.Context) {
	var user database.User
	var db = database.DatabaseInteraction()

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.AddUser(db, user)
}