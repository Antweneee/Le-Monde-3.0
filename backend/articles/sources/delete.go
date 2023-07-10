package sources

import (
	"github.com/gin-gonic/gin"
	"github.com/ipfs/go-ipfs-api"
)

func DeleteIPFS(c *gin.Context) {
	cid := c.Param("cid")

	sh := shell.NewShell("localhost:5001")

	err := sh.Unpin(cid)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{"message": "cid deleted successfully : " + cid,})
}

// func DeleteDB(c *gin.Context, db *gorm.DB) {
// 	article := new(database.Article);

// 	if err := c.Bind(&article); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	cid := c.Query("cid")
// 	res := db.Where(database.Article{Cid: cid}).Find(&article)
// 	if res.Error != nil {
//         if errors.Is(res.Error, gorm.ErrRecordNotFound) {
//             c.JSON(http.StatusNotFound, gin.H{"error": res.Error})
// 			return
//         }
//         c.JSON(http.StatusInternalServerError, gin.H{"error": res.Error})
// 		return
//     }
// 	if article.Id == 0 {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "article not found"})
// 		return
// 	}

// 	id := article.Id
// 	condition := database.Article{Id: id}

//     result := db.Delete(&condition)
//     if result.Error != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error})
//     } else {
// 		c.JSON(http.StatusOK, gin.H{"delete" : "Article deleted successfully"})
// 	}
// }