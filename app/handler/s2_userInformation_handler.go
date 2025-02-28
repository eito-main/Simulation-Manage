package handler

import (
	"Simulation-Manage/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GetUser は指定されたIDのユーザ情報を返すハンドラです
// http://localhost:8080/users/1
func GetUser(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "無効なID"})
		return
	}

	user, err := model.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}
