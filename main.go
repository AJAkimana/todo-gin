package main

import (
	"net/http"

	"github.com/AJAkimana/todo-gin/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	routes.RigisterRoutes(router)
	router.Run(":8000")
}
