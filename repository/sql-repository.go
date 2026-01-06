package repository

import (
	"context"
	"fmt"
	"gocrudb/resource"
	"gocrudb/structure"
	"slices"

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

func (r SqlRepository[I, R]) Get(conditions structure.Conditions) ([]R, error) {
	query := r.withPagination(conditions.Pagination)

	if conditions.Sorts != nil {
		for _, sort := range conditions.Sorts {
			query = query.Order(fmt.Sprintf("%s %s", sort.Field, sort.Direction))
		}
	}

	return query.Find(context.Background())
}

func (r SqlRepository[I, R]) Find(id I) (R, error) {
	return r.manager.Where("id = ?", id).First(context.Background())
}

func (r SqlRepository[I, R]) Create(instance R) (R, error) {
	err := r.manager.Omit(instance.GetProtectedFields()...).Create(context.Background(), &instance)
	return instance, err
}

func (r SqlRepository[I, R]) Update(instance R) (R, error) {
	query := r.withRequestMap(r.manager.Where("id = ?", instance.GetId()), instance)

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

func (r SqlRepository[I, R]) withPagination(p structure.Paginate) gorm.ChainInterface[R] {
	page := int(p.Page)
	perPage := int(p.PerPage)

	skipPages := 0
	if page > 0 {
		skipPages = page - 1
	}
	query := r.manager.Offset(skipPages * perPage)

	if perPage > 0 {
		query = query.Limit(perPage)
	}

	return query
}

func (r SqlRepository[I, R]) withRequestMap(query gorm.ChainInterface[R], instance R) gorm.ChainInterface[R] {
	requestMap := instance.GetRequestMap()
	if len(requestMap) > 0 {
		for field := range requestMap {
			if r.isModifiableField(field, instance) {
				query = query.Select(field) // https://github.com/go-gorm/gorm/issues/7498
			}
		}
	}

	return query
}

func (r SqlRepository[I, R]) isModifiableField(field string, instance R) bool {
	return slices.Contains(instance.GetModifiableFields(), field)
}
