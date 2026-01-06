package exception

type InternalServerError struct {
	Reason string
}

func (e InternalServerError) Error() string {
	return "Internal Server Error"
}
