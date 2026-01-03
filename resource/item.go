package resource

import (
	"gocrudb/dto"
	"time"

	"github.com/google/uuid"
)

type Item struct {
	ID         uuid.UUID `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Name       string    `json:"name"`
	Stock      uint      `json:"stock"`
	Price      float64   `json:"price" gorm:"type:numeric(8,2)"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	requestMap map[string]any
}

func (item Item) GetId() uuid.UUID {
	return item.ID
}

func (item Item) GetCreatedAt() time.Time {
	return item.CreatedAt
}

func (item Item) GetUpdatedAt() time.Time {
	return item.UpdatedAt
}

func (item Item) GetProtectedFields() []string {
	return []string{"ID"}
}

func (item Item) GetModifiableFields() []string {
	return []string{"Name", "Stock", "Price"}
}

func (item Item) GetRequestMap() map[string]any {
	return item.requestMap
}

func (item Item) FromReuestDto(d dto.RequestDTO) Resource[uuid.UUID] {
	m := d.ToRequestMap()

	if name, exists := m["Name"]; exists {
		value, ok := name.(string)
		if ok {
			item.Name = value
		}
	}
	if stock, exists := m["Stock"]; exists {
		value, ok := stock.(uint)
		if ok {
			item.Stock = value
		}
	}
	if price, exists := m["Price"]; exists {
		value, ok := price.(float64)
		if ok {
			item.Price = value
		}
	}

	item.requestMap = m

	return item
}
