package handler

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// HelloHandler は"Hello"を返すハンドラです
// http://localhost:8080/hello
func HelloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello"})
}