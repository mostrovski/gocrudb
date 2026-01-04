package dto

type UpdateItem struct {
	Name  *string  `json:"name" binding:"omitnil,min=1"`
	Stock *uint    `json:"stock" binding:"omitnil,gte=0"`
	Price *float64 `json:"price" binding:"omitnil,gte=0"`
}

func (request UpdateItem) ToRequestMap() map[string]any {
	m := map[string]any{}

	if request.Name != nil {
		m["Name"] = *request.Name
	}
	if request.Stock != nil {
		m["Stock"] = *request.Stock
	}
	if request.Price != nil {
		m["Price"] = *request.Price
	}

	return m
}
