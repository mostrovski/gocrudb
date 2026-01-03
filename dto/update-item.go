package dto

type UpdateItem struct {
	Name  *string  `json:"name" binding:"omitnil,min=1"`
	Stock *uint    `json:"stock" binding:"omitnil,gte=0"`
	Price *float64 `json:"price" binding:"omitnil,gte=0"`
}

func (d UpdateItem) ToRequestMap() map[string]any {
	m := map[string]any{}

	if d.Name != nil {
		m["Name"] = *d.Name
	}
	if d.Stock != nil {
		m["Stock"] = *d.Stock
	}
	if d.Price != nil {
		m["Price"] = *d.Price
	}

	return m
}
