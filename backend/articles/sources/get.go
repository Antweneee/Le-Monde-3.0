package sources

import (
	"github.com/gin-gonic/gin"
	"os"
	"github.com/ipfs/go-ipfs-api"
	"log"
)

func GetIPFS(c *gin.Context) {
    sh := shell.NewShell("localhost:5001")

    cid := c.Param("cid")
	cwd, _ := os.Getwd()
    c.File(cwd + "/" + cid)

    file, err := os.Create(cwd + "/" + cid)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    err = sh.Get(cid, cwd + "/" + cid)
    if err != nil {
        log.Fatal(err)
    }

    c.File(cwd + "/" + cid)

    c.JSON(200, gin.H{"message": "content successfully retrieved",})
}

// func GetDB(c *gin.Context, db *gorm.DB) {
// 	article := new(database.Article);

// 	if err := c.Bind(&article); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	cid := c.Query("cid")
// 	result := db.Where(database.Article{Cid: cid}).Find(&article)
// 	if result.Error != nil {
//         if errors.Is(result.Error, gorm.ErrRecordNotFound) {
//             c.JSON(http.StatusNotFound, gin.H{"error": result.Error})
// 			return
//         }
//         c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error})
// 		return
//     } else if article.Id == 0 {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "article not found"})
// 		return
// 	} else {
// 		c.JSON(http.StatusOK, article)
// 		return
// 	}
// }