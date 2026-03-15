package main

import (
	"todo/controllers"
	"todo/models"

	"github.com/google/uuid"
)

func main() {
	// Implementation of data structures

	repo := &models.TodoRepository{
		Todos: make(map[uuid.UUID]models.TodoStruct),
	}

	methods := models.SaveMethods(repo)

	// Call functions

	controllers.Menu(methods)

}
