package httperr

type errorCoded struct {
	err  error
	code int
}

func (err errorCoded) Error() string {
	return err.err.Error()
}

func (err errorCoded) StatusCode() int { return err.code }

func (err errorCoded) Unwrap() error { return err.err }

func NewError(code int, err error) errorCoded {
	return errorCoded{err, code}
}