package dto

type CreateItem struct {
	Name  *string  `json:"name" binding:"required,min=1"`
	Stock *uint    `json:"stock" binding:"required,gte=0"`
	Price *float64 `json:"price" binding:"required,gte=0"`
}

func (request CreateItem) ToRequestMap() map[string]any {
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
