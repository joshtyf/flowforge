package models

type HttpError struct {
	Error string
}

func NewHttpError(err error) *HttpError {
	return &HttpError{Error: err.Error()}
}
