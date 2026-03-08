package dto

import (
	"gocrudb/structure"
	"regexp"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
)

type QueryItem struct {
	Name         *string    `form:"name"`
	Stock        *uint      `form:"stock"`
	StockMin     *uint      `form:"stock_min"`
	StockMax     *uint      `form:"stock_max"`
	Price        *float64   `form:"price"`
	PriceMin     *float64   `form:"price_min"`
	PriceMax     *float64   `form:"price_max"`
	CreatedAtMin *time.Time `form:"created_at_min" time_format:"2006-01-02"`
	CreatedAtMax *time.Time `form:"created_at_max" time_format:"2006-01-02"`
	UpdatedAtMin *time.Time `form:"updated_at_min" time_format:"2006-01-02"`
	UpdatedAtMax *time.Time `form:"updated_at_max" time_format:"2006-01-02"`
	SortBy       *string    `form:"sort_by" binding:"omitnil,item_sort_by"`
	SortOrder    *string    `form:"sort_order" binding:"omitnil,item_sort_order"`
	Page         *uint      `form:"page"`
	PerPage      *uint      `form:"per_page"`
}

func (q QueryItem) ToQueryConditions() structure.Conditions {
	return structure.Conditions{
		Filters:    q.ToQueryFilters(),
		Sorts:      q.ToQuerySorts(),
		Pagination: q.ToQueryPagination(),
	}
}

func (q QueryItem) ToQueryFilters() []structure.FilterBy {
	f := make([]structure.FilterBy, 0, 15)

	if q.Name != nil {
		f = append(f, structure.FilterBy{Field: "name", Operator: "like", Value: *q.Name})
	}
	if q.Stock != nil {
		f = append(f, structure.FilterBy{Field: "stock", Operator: "=", Value: *q.Stock})
	}
	if q.Price != nil {
		f = append(f, structure.FilterBy{Field: "price", Operator: "=", Value: *q.Price})
	}

	if q.StockMin != nil && q.Stock == nil {
		f = append(f, structure.FilterBy{Field: "stock", Operator: ">=", Value: *q.StockMin})
	}
	if q.PriceMin != nil && q.Price == nil {
		f = append(f, structure.FilterBy{Field: "price", Operator: ">=", Value: *q.PriceMin})
	}
	if q.CreatedAtMin != nil {
		f = append(f, structure.FilterBy{Field: "created_at", Operator: ">=", Value: *q.CreatedAtMin})
	}
	if q.UpdatedAtMin != nil {
		f = append(f, structure.FilterBy{Field: "updated_at", Operator: ">=", Value: *q.UpdatedAtMin})
	}

	if q.StockMax != nil && q.Stock == nil {
		f = append(f, structure.FilterBy{Field: "stock", Operator: "<=", Value: *q.StockMax})
	}
	if q.PriceMax != nil && q.Price == nil {
		f = append(f, structure.FilterBy{Field: "price", Operator: "<=", Value: *q.PriceMax})
	}
	if q.CreatedAtMax != nil {
		f = append(f, structure.FilterBy{Field: "created_at", Operator: "<=", Value: *q.CreatedAtMax})
	}
	if q.UpdatedAtMax != nil {
		f = append(f, structure.FilterBy{Field: "updated_at", Operator: "<=", Value: *q.UpdatedAtMax})
	}

	return f
}

func (q QueryItem) ToQuerySorts() []structure.SortBy {
	s := make([]structure.SortBy, 0, 10)
	defaultSortField := "created_at"
	defaultDirection := "asc"

	sortBy := defaultSortField
	if q.SortBy != nil {
		sortBy = *q.SortBy
	}
	sortOrder := defaultDirection
	if q.SortOrder != nil {
		sortOrder = *q.SortOrder
	}

	uniqueSortFields := map[string]bool{}
	sortFields := strings.Split(sortBy, ",")
	directions := strings.Split(sortOrder, ",")
	directionsCount := len(directions)
	for i, field := range sortFields {
		if _, seen := uniqueSortFields[field]; seen {
			continue
		}
		uniqueSortFields[field] = true

		direction := defaultDirection
		if i < directionsCount {
			direction = directions[i]
		}

		s = append(s, structure.SortBy{Field: field, Direction: direction})
	}

	return s
}

func (q QueryItem) ToQueryPagination() structure.Paginate {
	p := structure.Paginate{}

	if q.Page != nil {
		p.Page = *q.Page
	}
	if q.PerPage != nil {
		p.PerPage = *q.PerPage
	}

	return p
}

var ValidItemSortBy validator.Func = func(fl validator.FieldLevel) bool {
	re := regexp.MustCompile(`^(name|stock|price|created_at|updated_at)(,(name|stock|price|created_at|updated_at))*$`)
	return re.MatchString(fl.Field().String())
}

var ValidItemSortOrder validator.Func = func(fl validator.FieldLevel) bool {
	re := regexp.MustCompile(`^(asc|desc)(,(asc|desc))*$`)
	return re.MatchString(fl.Field().String())
}
