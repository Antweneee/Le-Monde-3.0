package router

import (
	"bytes"
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/lib/pq"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	src "main/sources"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

func fakeDB() *gorm.DB {
	if _, err := os.Stat("fakeBookmarks.db"); err == nil {
		if err := os.Remove("fakeBookmarks.db"); err != nil {
			log.Fatal("failed to remove existing database:", err)
		}
	}

	db, err := gorm.Open(sqlite.Open("fakeBookmarks.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}

	db.AutoMigrate(&src.Bookmark{})

	return db
}

var db *gorm.DB
var token string
var fakeBookmark src.Bookmark
var router *gin.Engine

func generateFakeToken() (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = 1
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("secret_key")))
}

func setUp() {
	db = fakeDB()
	token, _ = generateFakeToken()
	fakeBookmark = src.Bookmark{
		Id:       1,
		Title:    "TestTitle",
		UserId:   1,
		Articles: pq.Int32Array{},
	}
	router = Router(db)
}

func TestAddBookmark(t *testing.T) {

	setUp()

	articleJSON, _ := json.Marshal(fakeBookmark)
	w := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/bookmarks", bytes.NewReader(articleJSON))
	if err != nil {
		log.Fatalf("impossible to build request: %s", err)
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	assert.Equal(t, 201, w.Code)

	var responseBookmark src.Bookmark
	err = json.Unmarshal([]byte(w.Body.String()), &responseBookmark)
	if err != nil {
		log.Fatalf("error unmarshaling response: %s", err)
	}
	assert.Equal(t, "TestTitle", responseBookmark.Title)
}

func TestAddArticleInBookmark(t *testing.T) {

	setUp()

	result := db.Create(&fakeBookmark)
	if result.Error != nil {
		panic(result.Error)
	}
	articleJSON, _ := json.Marshal(fakeBookmark)
	w := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/bookmarks/1/articles/1", bytes.NewReader(articleJSON))
	if err != nil {
		log.Fatalf("impossible to build request: %s", err)
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	var responseBookmark src.Bookmark
	err = json.Unmarshal([]byte(w.Body.String()), &responseBookmark)
	if err != nil {
		log.Fatalf("error unmarshaling response: %s", err)
	}
	assert.Equal(t, pq.Int32Array{1}, responseBookmark.Articles)
}

func TestGetAllBookmarks(t *testing.T) {
	setUp()

	result := db.Create(&fakeBookmark)
	if result.Error != nil {
		panic(result.Error)
	}
	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/bookmarks", nil)
	if err != nil {
		log.Fatalf("impossible to build request: %s", err)
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	var responseBookmarks []src.Bookmark
	err = json.Unmarshal([]byte(w.Body.String()), &responseBookmarks)
	if err != nil {
		log.Fatalf("error unmarshaling response: %s", err)
	}
	assert.Equal(t, []src.Bookmark{fakeBookmark}, responseBookmarks)
}

func TestGetBookmark(t *testing.T) {
	setUp()

	result := db.Create(&fakeBookmark)
	if result.Error != nil {
		panic(result.Error)
	}
	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/bookmarks/1", nil)
	if err != nil {
		log.Fatalf("impossible to build request: %s", err)
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	var responseBookmark src.Bookmark
	err = json.Unmarshal([]byte(w.Body.String()), &responseBookmark)
	if err != nil {
		log.Fatalf("error unmarshaling response: %s", err)
	}
	assert.Equal(t, "TestTitle", responseBookmark.Title)
}

func TestGetAllArticlesBookmark(t *testing.T) {

}

func TestEditBookmark(t *testing.T) {
	setUp()

	result := db.Create(&fakeBookmark)
	if result.Error != nil {
		panic(result.Error)
	}
	editedBookmark := src.EditedBookmark{Title: "EditedTitle"}
	articleJSON, _ := json.Marshal(editedBookmark)
	w := httptest.NewRecorder()
	req, err := http.NewRequest("PUT", "/bookmarks/1", bytes.NewReader(articleJSON))
	if err != nil {
		log.Fatalf("impossible to build request: %s", err)
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	var responseBookmark src.Bookmark
	err = json.Unmarshal([]byte(w.Body.String()), &responseBookmark)
	if err != nil {
		log.Fatalf("error unmarshaling response: %s", err)
	}
	assert.Equal(t, "EditedTitle", responseBookmark.Title)
}

func TestDeleteAllBookmarks(t *testing.T) {
	setUp()

	result := db.Create(&fakeBookmark)
	if result.Error != nil {
		panic(result.Error)
	}

	w := httptest.NewRecorder()
	req, err := http.NewRequest("DELETE", "/bookmarks", nil)
	if err != nil {
		log.Fatalf("impossible to build request: %s", err)
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	var body map[string]string
	err = json.Unmarshal([]byte(w.Body.String()), &body)
	if err != nil {
		log.Fatalf("error unmarshaling response: %s", err)
	}
	assert.Equal(t, "all bookmarks have been successfully deleted", body["delete"])
}

func TestDeleteBookmark(t *testing.T) {
	setUp()

	result := db.Create(&fakeBookmark)
	if result.Error != nil {
		panic(result.Error)
	}

	w := httptest.NewRecorder()
	req, err := http.NewRequest("DELETE", "/bookmarks/1", nil)
	if err != nil {
		log.Fatalf("impossible to build request: %s", err)
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	var body map[string]string
	err = json.Unmarshal([]byte(w.Body.String()), &body)
	if err != nil {
		log.Fatalf("error unmarshaling response: %s", err)
	}
	assert.Equal(t, "bookmark has been deleted successfully", body["delete"])
}

func TestDeleteAllArticlesBookmark(t *testing.T) {
	setUp()

	fakeBookmark.Articles = pq.Int32Array{1, 2, 3}

	result := db.Create(&fakeBookmark)
	if result.Error != nil {
		panic(result.Error)
	}

	w := httptest.NewRecorder()
	req, err := http.NewRequest("DELETE", "/bookmarks/1/articles", nil)
	if err != nil {
		log.Fatalf("impossible to build request: %s", err)
	}
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	var responseBookmark src.Bookmark
	err = json.Unmarshal([]byte(w.Body.String()), &responseBookmark)
	if err != nil {
		log.Fatalf("error unmarshaling response: %s", err)
	}
	assert.Equal(t, pq.Int32Array{}, responseBookmark.Articles)
}

func TestDeleteArticleBookmark(t *testing.T) {
	setUp()

	fakeBookmark.Articles = pq.Int32Array{1, 2, 3}

	result := db.Create(&fakeBookmark)
	if result.Error != nil {
		panic(result.Error)
	}

	w := httptest.NewRecorder()
	req, err := http.NewRequest("DELETE", "/bookmarks/1/articles/1", nil)
	if err != nil {
		log.Fatalf("impossible to build request: %s", err)
	}
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	var responseBookmark src.Bookmark
	err = json.Unmarshal([]byte(w.Body.String()), &responseBookmark)
	if err != nil {
		log.Fatalf("error unmarshaling response: %s", err)
	}
	assert.Equal(t, pq.Int32Array{2, 3}, responseBookmark.Articles)
}
