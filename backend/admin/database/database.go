package database

import (
	"log"
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
	src "main/sources"
)

func DatabaseInit() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("database/admin.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}

	db.AutoMigrate(&src.User{})

	return db;
}