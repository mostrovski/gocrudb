package repository

import (
	"context"
	"errors"
	"fmt"
	"gocrudb/exception"
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

	if conditions.Filters != nil {
		query = r.withFilters(query, conditions.Filters)
	}
	if conditions.Sorts != nil {
		query = r.withSorts(query, conditions.Sorts)
	}

	return query.Find(context.Background())
}

func (r SqlRepository[I, R]) Find(id I) (R, error) {

	instance, err := r.manager.Where("id = ?", id).First(context.Background())
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return instance, exception.ResourceNotFound{Id: id}
	}
	return instance, err
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
	if _, err := r.Find(id); err != nil {
		return err
	}
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

func (r SqlRepository[I, R]) withFilters(query gorm.ChainInterface[R], filters []structure.FilterBy) gorm.ChainInterface[R] {
	for _, f := range filters {
		operator := f.Operator
		value := f.Value
		if operator == "like" {
			operator = "ILIKE"
			value = "%" + value.(string) + "%"
		}
		query = query.Where(fmt.Sprintf("%s %s ?", f.Field, operator), value)
	}

	return query
}

func (r SqlRepository[I, R]) withSorts(query gorm.ChainInterface[R], sorts []structure.SortBy) gorm.ChainInterface[R] {
	for _, s := range sorts {
		query = query.Order(fmt.Sprintf("%s %s", s.Field, s.Direction))
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
