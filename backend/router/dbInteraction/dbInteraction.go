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
	c.JSON(http.StatusCreated, gin.H{"Created" : "User created successfully"})
	database.AddUser(db, user)
}

func DeleteUser(c *gin.Context) {
	var user database.User;
	var db = database.DatabaseInteraction()

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := db.Delete(&user)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error})
        return
    }
	c.JSON(http.StatusOK, gin.H{"delete" : "User deleted successfully"})
    return
}

func UpdateUser(c *gin.Context) {
	var user database.User;
	var db = database.DatabaseInteraction()

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := db.Save(&user)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error})
        return
    }
	c.JSON(http.StatusOK, gin.H{"Update" : "User updated successfully"})
    return
}