package handlers

import (
	"net/http"
	"strconv"

	"github.com/AJAkimana/todo-gin/models"
	"github.com/gin-gonic/gin"
)

var todos = []models.Todo{}
var nextID = 1

func GetTodo(c *gin.Context) {
	c.JSON(http.StatusOK, todos)
}

func AddTodo(c *gin.Context) {
	var todo models.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo.ID = nextID
	nextID++
	todos = append(todos, todo)
	c.JSON(http.StatusCreated, todo)
}

func DeleteTodo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	for i, t := range todos {
		if t.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "delete"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
}
