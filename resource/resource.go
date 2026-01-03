package resource

import (
	"gocrudb/dto"
	"time"

	"github.com/google/uuid"
)

type IdType interface {
	uint | uuid.UUID
}

type Resource[I IdType] interface {
	GetId() I
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
	GetProtectedFields() []string
	GetModifiableFields() []string
	GetRequestMap() map[string]any
	FromReuestDto(d dto.RequestDTO) Resource[I]
}
