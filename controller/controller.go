package controller

import (
	"gocrudb/exception"
	"net/http"
)

type Response struct {
	StatusCode int
	Body       map[string]any
}

func ErrorResponse(e error) Response {
	var statusCode int
	var message string

	switch e.(type) {
	case exception.ResourceNotFound:
		statusCode = http.StatusNotFound
	case exception.InvalidRequest:
		statusCode = http.StatusBadRequest
	case exception.InvalidPayload:
		statusCode = http.StatusUnprocessableEntity
	default:
		statusCode = http.StatusInternalServerError
		message = exception.InternalServerError{}.Error()
	}

	if message == "" {
		message = e.Error()
	}

	return Response{StatusCode: statusCode, Body: map[string]any{"error": message}}
}
