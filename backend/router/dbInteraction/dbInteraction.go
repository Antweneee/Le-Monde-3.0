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
	err := database.AddUser(db, user)
	if (err != nil) {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"Created" : "User created successfully"})
}

func DeleteUser(c *gin.Context) {
	var user database.User;
	var db = database.DatabaseInteraction()

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := database.DeleteUser(db, user.Id);
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
        return
    }
	c.JSON(http.StatusOK, gin.H{"delete" : "User deleted successfully"})
}

func UpdateUser(c *gin.Context) {
	var user database.User;
	var db = database.DatabaseInteraction()

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := database.UpdateUser(db, user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
        return
    }
	c.JSON(http.StatusOK, gin.H{"Update" : "User updated successfully"})
}

func AddArticle(c *gin.Context) {
	var article database.Article
	var db = database.DatabaseInteraction()

	if err := c.ShouldBindJSON(&article); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := database.AddArticle(db, article)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
        return
	}
	c.JSON(http.StatusCreated, gin.H{"Created" : "Article created successfully"})
}

func DeleteArticle(c *gin.Context) {
	var article database.Article;
	var db = database.DatabaseInteraction()

	if err := c.ShouldBindJSON(&article); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := database.DeleteArticle(db, article.Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
        return
    }
	c.JSON(http.StatusOK, gin.H{"delete" : "Article deleted successfully"})
    return
}

func UpdateArticle(c *gin.Context) {
	var article database.Article;
	var db = database.DatabaseInteraction()

	if err := c.ShouldBindJSON(&article); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := database.UpdateArticle(db, article)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
        return
    }
	c.JSON(http.StatusOK, gin.H{"Update" : "Article updated successfully"})
    return
}

func AddLike(c *gin.Context) {
	var like database.Like
	var db = database.DatabaseInteraction()

	if err := c.ShouldBindJSON(&like); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := database.AddLike(db, like)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
        return
	}
	c.JSON(http.StatusCreated, gin.H{"Created" : "Like created successfully"})
}

func DeleteLike(c *gin.Context) {
	var like database.Like;
	var db = database.DatabaseInteraction()

	if err := c.ShouldBindJSON(&like); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := database.DeleteLike(db, like.Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
        return
    }
	c.JSON(http.StatusOK, gin.H{"delete" : "Like deleted successfully"})
    return
}

func UpdateLike(c *gin.Context) {
	var like database.Like;
	var db = database.DatabaseInteraction()

	if err := c.ShouldBindJSON(&like); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := database.UpdateLike(db, like)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
        return
    }
	c.JSON(http.StatusOK, gin.H{"Update" : "Like updated successfully"})
    return
}

func AddBookmark(c *gin.Context) {
	var bookmark database.Bookmark
	var db = database.DatabaseInteraction()

	if err := c.ShouldBindJSON(&bookmark); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := database.AddBookmark(db, bookmark)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
        return
	}
	c.JSON(http.StatusCreated, gin.H{"Created" : "Bookmark created successfully"})
}

func DeleteBookmark(c *gin.Context) {
	var bookmark database.Bookmark
	var db = database.DatabaseInteraction()

	if err := c.ShouldBindJSON(&bookmark); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := database.DeleteBookmark(db, bookmark.Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
        return
    }
	c.JSON(http.StatusOK, gin.H{"delete" : "Bookmark deleted successfully"})
    return
}

func UpdateBookmark(c *gin.Context) {
	var bookmark database.Bookmark
	var db = database.DatabaseInteraction()

	if err := c.ShouldBindJSON(&bookmark); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := database.UpdateBookmark(db, bookmark)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
        return
    }
	c.JSON(http.StatusOK, gin.H{"Update" : "Bookmark updated successfully"})
    return
}

func AddTopic(c *gin.Context) {
	var topic database.Topic
	var db = database.DatabaseInteraction()

	if err := c.ShouldBindJSON(&topic); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := database.AddTopic(db, topic)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
        return
	}
	c.JSON(http.StatusCreated, gin.H{"Created" : "Topic created successfully"})
}

func DeleteTopic(c *gin.Context) {
	var topic database.Topic
	var db = database.DatabaseInteraction()

	if err := c.ShouldBindJSON(&topic); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := database.DeleteTopic(db, topic.Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
        return
    }
	c.JSON(http.StatusOK, gin.H{"delete" : "Topix deleted successfully"})
    return
}

func UpdateTopic(c *gin.Context) {
	var topic database.Topic
	var db = database.DatabaseInteraction()

	if err := c.ShouldBindJSON(&topic); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := database.UpdateTopic(db, topic)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
        return
    }
	c.JSON(http.StatusOK, gin.H{"Update" : "Topic updated successfully"})
    return
}