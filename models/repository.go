package models

import (
	"fmt"

	"github.com/google/uuid"
)

func (t *TodoRepository) Create(title string, description string) (*TodoStruct, error) {
	todo := TodoStruct{
		ID:          uuid.New(),
		Title:       title,
		Description: description,
		Status:      false,
	}
	t.Todos[todo.ID] = todo
	return &todo, nil
}

func (t *TodoRepository) Edit(id uuid.UUID, title string, description string) error {
	//if _, ok := t.todos[id]; !ok {
	//	return fmt.Errorf("todo not found")
	//}
	//t.todos[id] = TodoStruct{
	//	ID:          id,
	//	Title:       title,
	//	Description: description,
	//	Status:      false,
	//}
	//return nil

	todo, ok := t.Todos[id]
	if !ok {
		return fmt.Errorf("todo not found")
	}

	todo.Title = title
	todo.Description = description
	t.Todos[id] = todo
	return nil
}

func (t *TodoRepository) Update(id uuid.UUID, status bool) error {
	//if _, ok := t.todos[id]; !ok {
	//	return fmt.Errorf("todo not found")
	//}
	//t.todos[id] = TodoStruct{
	//	ID:     id,
	//	Status: status,
	//}
	//return nil

	todo, ok := t.Todos[id]
	if !ok {
		return fmt.Errorf("todo not found")
	}

	todo.Status = status
	t.Todos[id] = todo
	return nil
}

func (t *TodoRepository) Delete(id uuid.UUID) error {
	if _, ok := t.Todos[id]; !ok {
		return fmt.Errorf("todo not found")
	}
	delete(t.Todos, id)
	return nil
}

func (t *TodoRepository) Get(id uuid.UUID) (*TodoStruct, error) {
	todo, ok := t.Todos[id]
	if !ok {
		return nil, fmt.Errorf("todo not found")
	}

	return &todo, nil
}

func (t *TodoRepository) List() ([]TodoStruct, error) {
	var todos []TodoStruct
	for _, todo := range t.Todos {
		todos = append(todos, todo)
	}
	return todos, nil
}
