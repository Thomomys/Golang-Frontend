package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Define a struct to represent a ToDo item
type ToDoItem struct {
	ID    int
	Title string
	Done  bool
}

func main() {
	r := gin.Default()

	var todoList []ToDoItem

	// Render the index page with the ToDo list
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"TodoList": todoList,
		})
	})

	// Handle form submission to add a new ToDo item
	r.POST("/add", func(c *gin.Context) {
		title := c.PostForm("title")
		newTodo := ToDoItem{
			ID:    len(todoList) + 1,
			Title: title,
			Done:  false,
		}
		todoList = append(todoList, newTodo)
		c.Redirect(http.StatusFound, "/")
	})

	r.LoadHTMLFiles("templates/index.html")

	r.Run(":8080")
}
