package databaseInteraction

import (
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
	"github.com/gin-gonic/gin"
	"time"
  )

type User struct {
	gorm.Model
	id uint
	email string
	username string
	password string
	written_articles []Article
	liked_articles []Like
	bookmarks []Bookmark
}

type Article struct {
	gorm.Model
	id uint
	user_id uint
	title string
	cid string
	data time.Time
	topic Topic
	likes []Like
}

type Like struct {
	gorm.Model
	id uint
	user_id uint
	articles_id uint
}

type Bookmark struct {
	gorm.Model
	id uint
	name string
	description string
	user_id uint
	articles []Article
}

type Topic struct {
	gorm.Model
	name string
}

// * AddUser adds a user to the database.
// * It takes a gorm.DB instance and a user parameter representing the User object to be added.
// * It returns an error if the operation fails, or nil if successful.
func AddUser(db *gorm.DB, user User) error {
	result := db.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// func UpdateUser(db *gorm.DB, user User) error {}

// * DeleteUser deletes a user from the database based on the given userID.
// * It takes a gorm.DB instance and a userID parameter representing the ID of the user to be deleted.
// * It returns an error if the deletion fails, or nil if successful.
func DeleteUser(db *gorm.DB, userID uint) error {
	var user User
	result := db.First(&user, userID)
	if result.Error != nil {
		return result.Error
	}

	result = db.Delete(&user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func databaseInteraction(c *gin.Context) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
  	if err != nil {
    	panic("failed to connect database")
  	}
	db.AutoMigrate(&User{}, &Article{}, &Like{}, &Bookmark{}, &Topic{})
	
}