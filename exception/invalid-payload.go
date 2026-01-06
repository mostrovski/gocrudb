package exception

import "fmt"

type InvalidPayload struct {
	Reason string
}

func (e InvalidPayload) Error() string {
	return fmt.Sprintf("%s", e.Reason)
}
