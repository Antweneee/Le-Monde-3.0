package database

import (
	// "fmt"
	"log"
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
	"time"
  )

type User struct {
	gorm.Model
	Id uint
	Email string
	Username string
	Password string
	Written_articles []Article
	Liked_articles []Like
	Bookmarks []Bookmark
}

type Article struct {
	gorm.Model
	Id uint
	User_id uint
	Title string
	Cid string
	Data time.Time
	Topic Topic
	Likes []Like
}

type Like struct {
	gorm.Model
	Id uint
	User_id uint
	Articles_id uint
}

type Bookmark struct {
	gorm.Model
	Id uint
	Name string
	Description string
	User_id uint
	Articles []Article
}

type Topic struct {
	gorm.Model
	Id uint
	Name string
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

// * UpdateUser updates a user in the database.
// * It takes a gorm.DB instance and a user parameter representing the updated User object.
// * It returns an error if the operation fails, or nil if successful.
func UpdateUser(db *gorm.DB, user User) error {
	result := db.Save(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// * AddArticle adds a article to the database.
// * It takes a gorm.DB instance and a article parameter representing the Article object to be added.
// * It returns an error if the operation fails, or nil if successful.
func AddArticle(db *gorm.DB, article Article) error {
	result := db.Create(&article)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// * DeleteArticle deletes a article from the database based on the given articleID.
// * It takes a gorm.DB instance and a articleID parameter representing the ID of the article to be deleted.
// * It returns an error if the deletion fails, or nil if successful.
func DeleteArticle(db *gorm.DB, articleID uint) error {
	var article Article
	result := db.First(&article, articleID)
	if result.Error != nil {
		return result.Error
	}

	result = db.Delete(&article)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// * UpdateArticle updates a article in the database.
// * It takes a gorm.DB instance and a article parameter representing the updated Article object.
// * It returns an error if the operation fails, or nil if successful.
func UpdateArticle(db *gorm.DB, article Article) error {
	result := db.Save(&article)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// * AddLike adds a like to the database.
// * It takes a gorm.DB instance and a like parameter representing the Like object to be added.
// * It returns an error if the operation fails, or nil if successful.
func AddLike(db *gorm.DB, like Like) error {
	result := db.Create(&like)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// * DeleteLike deletes a like from the database based on the given likeID.
// * It takes a gorm.DB instance and a likeID parameter representing the ID of the like to be deleted.
// * It returns an error if the deletion fails, or nil if successful.
func DeleteLike(db *gorm.DB, likeID uint) error {
	var like Like
	result := db.First(&like, likeID)
	if result.Error != nil {
		return result.Error
	}

	result = db.Delete(&like)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// * UpdateLike updates a Like in the database.
// * It takes a gorm.DB instance and a Like parameter representing the updated Like object.
// * It returns an error if the operation fails, or nil if successful.
func UpdateLike(db *gorm.DB, like Like) error {
	result := db.Save(&like)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// * AddBookmark adds a bookmark to the database.
// * It takes a gorm.DB instance and a bookmark parameter representing the Bookmark object to be added.
// * It returns an error if the operation fails, or nil if successful.
func AddBookmark(db *gorm.DB, bookmark Bookmark) error {
	result := db.Create(&bookmark)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// * UpdateBookmark updates a bookmark in the database.
// * It takes a gorm.DB instance and a bookmark parameter representing the updated Bookmark object.
// * It returns an error if the operation fails, or nil if successful.
func UpdateBookmark(db *gorm.DB, bookmark Bookmark) error {
	result := db.Save(&bookmark)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// * DeleteBookmark deletes a bookmark from the database based on the given bookmarkID.
// * It takes a gorm.DB instance and a bookmarkID parameter representing the ID of the bookmark to be deleted.
// * It returns an error if the deletion fails, or nil if successful.
func DeleteBookmark(db *gorm.DB, bookmarkID uint) error {
	var bookmark Bookmark
	result := db.First(&bookmark, bookmarkID)
	if result.Error != nil {
		return result.Error
	}

	result = db.Delete(&bookmark)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// * AddTopic adds a topic to the database.
// * It takes a gorm.DB instance and a topic parameter representing the Topic object to be added.
// * It returns an error if the operation fails, or nil if successful.
func AddTopic(db *gorm.DB, topic Topic) error {
	result := db.Create(&topic)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// * UpdateTopic updates a topic in the database.
// * It takes a gorm.DB instance and a topic parameter representing the updated Topic object.
// * It returns an error if the operation fails, or nil if successful.
func UpdateTopic(db *gorm.DB, topic Topic) error {
	result := db.Save(&topic)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// * DeleteTopic deletes a topic from the database based on the given topicID.
// * It takes a gorm.DB instance and a topicID parameter representing the ID of the topic to be deleted.
// * It returns an error if the deletion fails, or nil if successful.
func DeleteTopic(db *gorm.DB, topicID uint) error {
	var topic Topic
	result := db.First(&topic, topicID)
	if result.Error != nil {
		return result.Error
	}

	result = db.Delete(&topic)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func DatabaseInteraction() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}

	db.AutoMigrate(&User{}, &Article{}, &Like{}, &Bookmark{}, &Topic{})

	return db;
}