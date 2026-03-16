package exception

import "fmt"

type InvalidRequest struct {
	Reason string
}

func (e InvalidRequest) Error() string {
	return fmt.Sprintf("Bad request: %s", e.Reason)
}
