package main

import (
	"net/http"
	"github.com/gin-gonic/gin"


)

type todo struct{
	ID 				string  `json:"id"`
	item			string	`json:"title"`
	Completed		bool	`json:"completed"`
}

var todos = []todo{
	{ID: "1", Item: "Clean Room", Completed: false},
	{ID: "1", Item: "Read Book", Completed: false},
	{ID: "1", Item: "Record Video", Completed: false},
	//Need to convert it to JSON... in getTodos function example
}

//* means type 
func getTodos(context *gin.Context){
	// Converting to JSON todos variable with list of objects
	context.IndentedJSON(http.StatusOK, todos)
}

func addTodos(context *gin.Context){
	//Create a new todo
	var newTodo todo

	if err := context.BindJSON(&newTodo); err != nil{
		return 
	}

	todos = append(todos, newTodo)
	
	context.IndentedJSON(http.StatusCreated, newTodo)

}

//Create Server
func main(){
	//It's going to create a server
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.POST("/todos", addTodos)

	//To run the server
	router.Run("localhost:9090") //specify the path

}

//Create first end point