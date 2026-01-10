package exception

import "fmt"

type InvalidRequest struct {
	Reason string
}

func (e InvalidRequest) Error() string {
	return fmt.Sprintf("%s", e.Reason)
}
