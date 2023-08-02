package sources

import (
	"errors"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"html"
	"net/http"
	"strings"
)

type User struct {
	gorm.Model
	Id       uint
	Email    string
	Username string
	Password string
}

type ReceiveUser struct {
	Email string `json:"email"`
}

func AddUser(email string, username string, password string, c *gin.Context, db *gorm.DB) {
	user := new(User)

	user.Email = html.EscapeString(strings.TrimSpace(email))
	user.Username = html.EscapeString(strings.TrimSpace(username))
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error})
		return
	}

	user.Password = string(hashedPassword)

	result := db.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error})
	} else {
		c.JSON(http.StatusCreated, gin.H{"Created": "User created successfully"})
	}
}

func GetUser(c *gin.Context, db *gorm.DB) {
	user := new(User)

	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	email := c.Query("email")
	result := db.Where(User{Email: email}).Find(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": result.Error})
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error})
	} else if user.Id == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func DeleteUser(c *gin.Context, db *gorm.DB) {
	user := new(User)

	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	email := c.Query("email")
	res := db.Where(User{Email: email}).Find(&user)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": res.Error})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": res.Error})
		return
	}
	if user.Id == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	id := user.Id
	condition := User{Id: id}

	result := db.Delete(&condition)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error})
	} else {
		c.JSON(http.StatusOK, gin.H{"delete": "User deleted successfully"})
	}
}
