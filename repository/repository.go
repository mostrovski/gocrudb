package repository

import (
	"gocrudb/resource"
	"gocrudb/structure"
)

type Repository[I resource.IdType, R resource.Resource[I]] interface {
	Get(conditions structure.Conditions) ([]R, error)
	Find(id I) (R, error)
	Create(instance R) (R, error)
	Update(instance R) (R, error)
	Delete(id I) error
}
