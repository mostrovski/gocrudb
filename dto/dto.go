package dto

type RequestDTO interface {
	ToRequestMap() map[string]any
}
