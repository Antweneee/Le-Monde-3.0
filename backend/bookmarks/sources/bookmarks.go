package sources

import (
	"bytes"
	"encoding/json"
	"errors"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"gorm.io/gorm"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func getUserId(c *gin.Context) (int32, error) {

	bearerToken := c.Request.Header.Get("Authorization")
	tokenString := ""

	if len(strings.Split(bearerToken, " ")) == 2 {
		tokenString = strings.Split(bearerToken, " ")[1]
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid token format"})
	}

	tokenPure, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("secret_key")), nil
	})

	if err != nil {
		return 0, err
	}

	if claims, ok := tokenPure.Claims.(jwt.MapClaims); ok && tokenPure.Valid {
		userID, _ := claims["user_id"].(float64)
		return int32(userID), nil
	} else {
		return 0, errors.New("invalid token")
	}
}

func ExtractToken(c *gin.Context) string {
	token := c.Query("token")
	if token != "" {
		return token
	}
	bearerToken := c.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

func MakeHTTPRequest(c *gin.Context, method string, url string, requestBody interface{}) ([]byte, int, error) {
	jsonParams, err := json.Marshal(requestBody)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	request, err := http.NewRequest(method, url, bytes.NewBuffer(jsonParams))
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+ExtractToken(c))

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	defer response.Body.Close()

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return responseBody, response.StatusCode, nil
}

func getArticleId(c *gin.Context, articleName string) (int32, error) {
	var titleArticle TitleArticle
	responseBody, _, err := MakeHTTPRequest(c, http.MethodGet, "http://articles-lemonde3-0:8082/articles/"+articleName, nil)
	if err != nil {
		return -1, err
	}

	err = json.Unmarshal(responseBody, &titleArticle)
	if err != nil {
		return 0, err
	}
	return titleArticle.Id, nil
}

func getArticlesFromBookmark(c *gin.Context, articlesId []int32) ([]Article, error) {
	var articlesBookmark []Article

	for i := 0; i < len(articlesId); i++ {
		article, err := getArticleById(c, articlesId[i])
		if err != nil {
			return nil, err
		}
		articlesBookmark = append(articlesBookmark, article)
	}
	return articlesBookmark, nil
}

func isArticleInBookmark(articlesBookmark []int32, articleId int32) bool {
	for i := 0; i < len(articlesBookmark); i++ {
		if articleId == articlesBookmark[i] {
			return true
		}
	}
	return false
}

func getArticleByTitle(c *gin.Context, titleArticle string) (Article, error) {
	var article Article
	responseBody, _, err := MakeHTTPRequest(c, http.MethodGet, "http://articles-lemonde3-0:8082/articles/"+titleArticle, nil)
	if err != nil {
		return Article{}, err
	}
	err = json.Unmarshal(responseBody, &article)
	if err != nil {
		return Article{}, err
	}
	return article, nil
}

func getArticleFromBookmark(c *gin.Context, db *gorm.DB, titleBookmark string, titleArticle string) (Article, error) {

	bookmark := new(Bookmark)
	article := Article{}

	userId, err := getUserId(c)
	if err != nil {
		return Article{}, err
	}

	result := db.Where(Bookmark{UserId: userId, Title: titleBookmark}).Find(&bookmark)
	if result.Error != nil {
		return Article{}, result.Error
	} else if bookmark.Title == "" {
		return Article{}, errors.New("bookmark not found")
	}

	articleId, err := getArticleId(c, titleArticle)
	if err != nil {
		return Article{}, err
	}

	if !(isArticleInBookmark(bookmark.Articles, articleId)) {
		return Article{}, errors.New("no article with the provided name was found in the bookmark")
	}

	article, err = getArticleByTitle(c, titleArticle)
	if err != nil {
		return Article{}, err
	}

	return article, nil
}

func getArticleById(c *gin.Context, articleId int32) (Article, error) {
	var article Article
	strArticleId := strconv.Itoa(int(articleId))
	responseBody, _, err := MakeHTTPRequest(c, http.MethodGet, "http://articles-lemonde3-0:8082/articles/"+strArticleId, nil)
	if err != nil {
		return Article{}, err
	}
	err = json.Unmarshal(responseBody, &article)
	if err != nil {
		return Article{}, err
	}
	return article, nil
}

type Bookmark struct {
	gorm.Model
	Id       int32
	UserId   int32
	Title    string
	Articles pq.Int32Array `gorm:"type:integer[]"`
}
