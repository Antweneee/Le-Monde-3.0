package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	src "main/sources"
)

func DatabaseInit() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("database/articles.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}

	db.AutoMigrate(&src.Article{})

	return db
}
