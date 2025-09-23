package main

import (
	"github.com/AJAkimana/todo-gin/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	routes.RigisterRoutes(r)
	r.Run(":8000")
}
