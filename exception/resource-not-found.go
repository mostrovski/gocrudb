package exception

import "fmt"

type ResourceNotFound struct {
	Resource string
	Id       any
}

func (e ResourceNotFound) Error() string {
	r := "Resource"
	if e.Resource != "" {
		r = e.Resource
	}

	return fmt.Sprintf("%s not found. ID: %v", r, e.Id)
}
