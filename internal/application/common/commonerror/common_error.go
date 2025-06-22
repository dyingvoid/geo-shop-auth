package commonerror

type ValidationError struct {
	Msg string
}

func (e *ValidationError) Error() string {
	return e.Msg
}

type DuplicateError struct {
	Msg string
}

func (e *DuplicateError) Error() string {
	return e.Msg
}
