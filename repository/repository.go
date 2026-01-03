package repository

import "gocrudb/resource"

type Repository[I resource.IdType, R resource.Resource[I]] interface {
	Get() ([]R, error)
	Find(id I) (R, error)
	Create(instance R) (R, error)
	Update(instance R) (R, error)
	Delete(id I) error
}
