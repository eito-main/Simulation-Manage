package server

import (
	"Simulation-Manage/handler"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/users/:id", handler.GetUser)
	r.GET("/hello", handler.HelloHandler)
	
	return r
}

func Run() {
	r := SetupRouter()
	r.Run(":8080")
}