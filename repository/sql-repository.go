package repository

import (
	"context"
	"gocrudb/resource"

	"gorm.io/gorm"
)

type SqlRepository[I resource.IdType, R resource.Resource[I]] struct {
	manager     gorm.Interface[R]
	initialized bool
}

func (r SqlRepository[I, R]) Init(db *gorm.DB) SqlRepository[I, R] {
	if r.initialized {
		return r
	}

	r.manager = gorm.G[R](db)
	r.initialized = true
	return r
}

func (r SqlRepository[I, R]) Get() ([]R, error) {
	return r.manager.Find(context.Background())
}

func (r SqlRepository[I, R]) Find(id I) (R, error) {
	return r.manager.Where("id = ?", id).First(context.Background())
}

func (r SqlRepository[I, R]) Create(instance R) (R, error) {
	err := r.manager.Omit(instance.GetProtectedFields()...).Create(context.Background(), &instance)
	return instance, err
}

func (r SqlRepository[I, R]) Update(instance R) (R, error) {
	query := r.manager.Where("id = ?", instance.GetId())

	requestMap := instance.GetRequestMap()
	if len(requestMap) > 0 {
		for field := range requestMap {
			query = query.Select(field) // https://github.com/go-gorm/gorm/issues/7498
		}
	}

	_, err := query.Omit(instance.GetProtectedFields()...).Updates(context.Background(), instance) // https://github.com/go-gorm/gorm/issues/7658

	if err != nil {
		return instance, err
	}

	return r.Find(instance.GetId())
}

func (r SqlRepository[I, R]) Delete(id I) error {
	_, err := r.manager.Where("id = ?", id).Delete(context.Background())
	return err
}
