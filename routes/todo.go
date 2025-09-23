package routes

import (
	"github.com/AJAkimana/todo-gin/handlers"
	"github.com/gin-gonic/gin"
)

func RigisterRoutes(r *gin.Engine) {
	todo := r.Group("/todos")
	{
		todo.GET("", handlers.GetTodo)
		todo.POST("", handlers.AddTodo)
		todo.DELETE("/:id", handlers.DeleteTodo)
	}
}
