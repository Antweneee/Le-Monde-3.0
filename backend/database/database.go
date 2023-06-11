package database

import (
	// "fmt"
	"log"
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
	"time"
  )

// Like represent a like
type Like struct {
	gorm.Model
	Id uint
	UserId uint
	User User
	ArticleId uint
	Article Article
}

// Topic represent a topic
type Topic struct {
	gorm.Model
	Id uint
	Name string
}

// Article represent a article
type Article struct {
	gorm.Model
	Id uint
	UserId uint
	User User
	Title string
	Cid string
	CreatedAt time.Time
	TopicId uint
	Topic Topic
	ArticleLikes []Like
}

// Bookmark represent a bookmark
type Bookmark struct {
	gorm.Model
	Id uint
	Name string
	Description string
	UserId uint
	User User
	BookmarkArticles []Article
}

// User represent a user
type User struct {
	gorm.Model
	Id uint64
	Email string
	Username string
	Password string
	WrittenArticles []Article
	LikedArticles []Like
	Bookmarks []Bookmark
}

func DatabaseInteraction() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}

	db.AutoMigrate(&User{}, &Article{}, &Like{}, &Bookmark{}, &Topic{})

	return db;
}