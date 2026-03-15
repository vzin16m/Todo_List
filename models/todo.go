package models

import (
	"github.com/google/uuid"
)

type TodoStruct struct {
	ID          uuid.UUID
	Title       string
	Description string
	Status      bool
}

// Slice

type TodoRepository struct {
	Todos map[uuid.UUID]TodoStruct
}
