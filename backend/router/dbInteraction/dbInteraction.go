package dbInteraction

import (
	"gorm.io/gorm"
	"net/http"
    "github.com/gin-gonic/gin"
	"backend/database"
	"errors"
)

type ReceiveUser struct {
    Email string `json:"email"`
}

func AddUser(c *gin.Context, db *gorm.DB) {
	user := new(database.User);

	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result := db.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error})
	} else {
		c.JSON(http.StatusCreated, gin.H{"Created" : "User created successfully"})
	}
}

func GetUser(c *gin.Context, db *gorm.DB) {
	user := new(database.User);

	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	email := c.Query("email")
	result := db.Where(database.User{Email: email}).Find(&user)
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
	user := new(database.User);

	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	email := c.Query("email")
	res := db.Where(database.User{Email: email}).Find(&user)
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
	condition := database.User{Id: id}

    result := db.Delete(&condition)
    if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error})
    } else {
		c.JSON(http.StatusOK, gin.H{"delete" : "User deleted successfully"})
	}
}

func AddArticle(c *gin.Context, db *gorm.DB) {
	article := new(database.Article)

	if err := c.Bind(&article); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result := db.Create(&article)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"Created" : "Article created successfully"})
}

func GetArticle(c *gin.Context, db *gorm.DB) {
	article := new(database.Article);

	if err := c.Bind(&article); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cid := c.Query("cid")
	result := db.Where(database.Article{Cid: cid}).Find(&article)
	if result.Error != nil {
        if errors.Is(result.Error, gorm.ErrRecordNotFound) {
            c.JSON(http.StatusNotFound, gin.H{"error": result.Error})
			return
        }
        c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error})
		return
    } else if article.Id == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "article not found"})
		return
	} else {
		c.JSON(http.StatusOK, article)
		return
	}
}

func DeleteArticle(c *gin.Context, db *gorm.DB) {
	article := new(database.Article);

	if err := c.Bind(&article); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cid := c.Query("cid")
	res := db.Where(database.Article{Cid: cid}).Find(&article)
	if res.Error != nil {
        if errors.Is(res.Error, gorm.ErrRecordNotFound) {
            c.JSON(http.StatusNotFound, gin.H{"error": res.Error})
			return
        }
        c.JSON(http.StatusInternalServerError, gin.H{"error": res.Error})
		return
    }
	if article.Id == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "article not found"})
		return
	}

	id := article.Id
	condition := database.Article{Id: id}

    result := db.Delete(&condition)
    if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error})
    } else {
		c.JSON(http.StatusOK, gin.H{"delete" : "Article deleted successfully"})
	}
}

func AddLike(c *gin.Context, db *gorm.DB) {
	like := new(database.Like)

	if err := c.Bind(&like); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result := db.Create(&like)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"Created" : "Like created successfully"})
}

func AddBookmark(c *gin.Context, db *gorm.DB) {
	bookmark := new(database.Bookmark)

	if err := c.Bind(&bookmark); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result := db.Create(&bookmark)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"Created" : "Bookmark created successfully"})
}

func AddTopic(c *gin.Context, db *gorm.DB) {
	topic := new(database.Topic)

	if err := c.Bind(&topic); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result := db.Create(&topic)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"Created" : "Topic created successfully"})
}
