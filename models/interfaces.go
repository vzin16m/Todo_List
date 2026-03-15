package models

import (
	"github.com/google/uuid"
)

type ICreateMethods interface {
	Create(title string, description string) (*TodoStruct, error)
	Edit(id uuid.UUID, title string, description string) error
	Update(id uuid.UUID, status bool) error
	Delete(id uuid.UUID) error
}

type IReadMethods interface {
	Get(id uuid.UUID) (*TodoStruct, error)
	List() ([]TodoStruct, error)
}

type SaveMethods interface {
	ICreateMethods
	IReadMethods
}
