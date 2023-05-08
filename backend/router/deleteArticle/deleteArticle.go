package deleteArticle

import (
	"github.com/gin-gonic/gin"
	"github.com/ipfs/go-ipfs-api"
)

func Delete(c *gin.Context) {
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

