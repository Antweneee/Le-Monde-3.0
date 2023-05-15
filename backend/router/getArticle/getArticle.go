package getArticle

import (
    "github.com/gin-gonic/gin"
    "github.com/ipfs/go-ipfs-api"
    "log"
    "os"
)

func Get(c *gin.Context) {
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
