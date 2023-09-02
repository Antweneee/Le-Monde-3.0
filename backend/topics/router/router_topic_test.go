package router

import (
	"github.com/go-playground/assert/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func fakeDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("fake.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}

	// db.AutoMigrate(&src.User{})

	return db
}

func TestRouterTopicCreation(t *testing.T) {

	router := Router(fakeDB())

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/topics", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 201, w.Code)
	assert.Equal(t, "{\"message\":\"Topic added successfully\"}", w.Body.String())
}

func TestRouterTopicRecuperation(t *testing.T) {

	router := Router(fakeDB())

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/topics/1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\"message\":\"Topic retreived successfully\"}", w.Body.String())
}

func TestRouterTopicSuppression(t *testing.T) {

	router := Router(fakeDB())

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/topics/1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\"message\":\"Topic deleted successfully\"}", w.Body.String())
}

func TestRouterTopicModification(t *testing.T) {

	router := Router(fakeDB())

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/topics/1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\"message\":\"Topic updated successfully\"}", w.Body.String())
}
