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
	CreatedAt    *time.Time `form:"created_at" time_format:"2006-01-02"`
	CreatedAtMin *time.Time `form:"created_at_min" time_format:"2006-01-02"`
	CreatedAtMax *time.Time `form:"created_at_max" time_format:"2006-01-02"`
	UpdatedAt    *time.Time `form:"updated_at" time_format:"2006-01-02"`
	UpdatedAtMin *time.Time `form:"updated_at_min" time_format:"2006-01-02"`
	UpdatedAtMax *time.Time `form:"updated_at_max" time_format:"2006-01-02"`
	SortBy       *string    `form:"sort_by" binding:"omitnil,item_sort_by"`
	SortOrder    *string    `form:"sort_order" binding:"omitnil,item_sort_order"`
}

func (q QueryItem) ToQuerySorts() []structure.SortBy {
	s := []structure.SortBy{}
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

var ValidItemSortBy validator.Func = func(fl validator.FieldLevel) bool {
	re := regexp.MustCompile(`^(name|stock|price|created_at|updated_at)(,(name|stock|price|created_at|updated_at))*$`)
	return re.MatchString(fl.Field().String())
}

var ValidItemSortOrder validator.Func = func(fl validator.FieldLevel) bool {
	re := regexp.MustCompile(`^(asc|desc)(,(asc|desc))*$`)
	return re.MatchString(fl.Field().String())
}
