package error

type HttpError struct {
	Code int
	Err  error
}

func New(code int, err error) HttpError {
	return HttpError{Code: code, Err: err}
}

func (e HttpError) Error() string {
	return e.Err.Error()
}
