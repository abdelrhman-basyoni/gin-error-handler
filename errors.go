package ginHandler

type DataBaseError struct {
	Message string
	code    int
}

func (e *DataBaseError) Error() string {
	return e.Message
}

type ValidationError struct {
	Message string
	code    int
}

func (e *ValidationError) Error() string {
	return e.Message
}
