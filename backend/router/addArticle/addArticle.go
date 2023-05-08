package addArticle

import (
    "github.com/gin-gonic/gin"
    "github.com/ipfs/go-ipfs-api"
)

func Add(c *gin.Context) {
    sh := shell.NewShell("localhost:5001")

    file, header, err := c.Request.FormFile("file")
    if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }
    defer file.Close()

    hash, err := sh.Add(file)
    if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }

    c.JSON(200, gin.H{"filename": header.Filename, "hash": hash})
}

