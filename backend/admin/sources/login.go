package sources

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"golang.org/x/crypto/bcrypt"
)

type LoginInput struct {
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}


func Login(c *gin.Context, db *gorm.DB) {
	var input LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := new(User);

	u.Email = input.Email
	u.Password = input.Password

	token, err := LoginCheck(u.Email, u.Password, db)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email or password is incorrect."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token":token})
}

func LoginCheck(username string, password string, db *gorm.DB) (string,error) {
	
	var err error

	u := new(User);

	err = db.Model(User{}).Where("email = ?", username).Take(&u).Error

	if err != nil {
		return "", err
	}

	err = VerifyPassword(password, u.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token,err := GenerateToken(u.ID)

	if err != nil {
		return "",err
	}

	return token,nil
	
}

func VerifyPassword(password,hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}