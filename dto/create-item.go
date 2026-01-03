package dto

type CreateItem struct {
	Name  *string  `json:"name" binding:"required,min=1"`
	Stock *uint    `json:"stock" binding:"required,gte=0"`
	Price *float64 `json:"price" binding:"required,gte=0"`
}

func (d CreateItem) ToRequestMap() map[string]any {
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
