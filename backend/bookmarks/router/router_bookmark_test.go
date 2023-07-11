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

func TestRouterBookmarkCreation(t *testing.T) {

	router := Router(fakeDB())

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/bookmarks", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 201, w.Code)
	assert.Equal(t, "{\"message\":\"Bookmarks added successfully\"}", w.Body.String())
}

func TestRouterBookmarkRecuperation(t *testing.T) {

	router := Router(fakeDB())

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/bookmarks/1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\"message\":\"Bookmark retreived successfully\"}", w.Body.String())
}

func TestRouterBookmarkSuppression(t *testing.T) {

	router := Router(fakeDB())

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/bookmarks/1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\"message\":\"Bookmark deleted successfully\"}", w.Body.String())
}

func TestRouterBookmarkModification(t *testing.T) {

	router := Router(fakeDB())

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/bookmarks/1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\"message\":\"Bookmark updated successfully\"}", w.Body.String())
}
