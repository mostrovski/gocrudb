package exception

import "fmt"

type ResourceNotFound struct {
	Id any
}

func (e ResourceNotFound) Error() string {
	return fmt.Sprintf("Resource not found. ID: %v", e.Id)
}
