package validation

import (
	"gocrudb/dto"

	"github.com/go-playground/validator/v10"
)

func TagValidators() map[string]validator.Func {
	return map[string]validator.Func{
		"item_sort_by":    dto.ValidItemSortBy,
		"item_sort_order": dto.ValidItemSortOrder,
	}
}
