package router

import (
	"bytes"
	"encoding/json"
	"github.com/go-playground/assert/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	src "main/sources"
	"net/http"
	"net/http/httptest"
	"testing"
)

type RegisterInput struct {
	Email    string `json:"email" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func fakeDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("fakedb.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}

	db.AutoMigrate(&src.User{})

	return db
}

func TestRegisterAdmin(t *testing.T) {

	requestBody := RegisterInput{
		Email:    "test@test.com",
		Username: "john",
		Password: "pass",
	}

	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatalf("Failed to marshal request body: %s", err)
	}

	router := Router(fakeDB())

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonBody))
	router.ServeHTTP(w, req)

	assert.Equal(t, 201, w.Code)
	assert.Equal(t, "{\"Created\":\"User created successfully\"}", w.Body.String())
}
