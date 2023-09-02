package router

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	src "main/sources"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var db *gorm.DB
var user src.User
var router *gin.Engine

func fakeDB() *gorm.DB {
	if _, err := os.Stat("fakeAdmin.db"); err == nil {
		if err := os.Remove("fakeAdmin.db"); err != nil {
			log.Fatal("failed to remove existing database:", err)
		}
	}

	db, err := gorm.Open(sqlite.Open("fakeAdmin.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}

	db.AutoMigrate(&src.User{})

	return db
}

func setUp() {
	db = fakeDB()

	user = src.User{
		Id:       1,
		Email:    "test@test.com",
		Username: "Bob",
		Password: "$2a$10$CZ4GTkQUTzM32wYykH5msuHuh5tZtnRpRZGEsBlk2Vhu4tsiSn1B2",
	}
	router = Router(db)
}

func TestLoginAdmin(t *testing.T) {

	setUp()

	result := db.Create(&user)
	if result.Error != nil {
		panic(result.Error)
	}
	requestBody := src.LoginInput{
		Email:    "test@test.com",
		Password: "pass",
	}

	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatalf("Failed to marshal request body: %s", err)
	}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonBody))
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestRegisterAdmin(t *testing.T) {

	setUp()

	requestBody := src.RegisterInput{
		Email:    "test@test.com",
		Username: "Bob",
		Password: "pass",
	}
	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatalf("Failed to marshal request body: %s", err)
	}

	w := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonBody))
	if err != nil {
		log.Fatalf("impossible to build request: %s", err)
	}
	router.ServeHTTP(w, req)

	assert.Equal(t, 201, w.Code)
	assert.Equal(t, "{\"created\":\"User created successfully\"}", w.Body.String())
}
