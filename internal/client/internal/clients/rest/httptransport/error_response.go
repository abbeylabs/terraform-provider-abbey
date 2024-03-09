package httptransport

import "fmt"

type ErrorResponse[T any] struct {
	Err         error
	IsHttpError bool
	Response[T]
}

func NewErrorResponse[T any](err error, resp *Response[T]) *ErrorResponse[T] {
	if resp == nil {
		return &ErrorResponse[T]{
			Err:         err,
			IsHttpError: false,
		}
	}

	return &ErrorResponse[T]{
		Err:         err,
		IsHttpError: true,
		Response:    *resp,
	}
}

func (r *ErrorResponse[T]) Clone() ErrorResponse[T] {
	if r == nil {
		return ErrorResponse[T]{}
	}

	clone := *r
	clone.Headers = make(map[string]string)
	for header, value := range r.Headers {
		clone.Headers[header] = value
	}
	return clone
}

func (r *ErrorResponse[T]) Error() string {
	return fmt.Sprintf("%s", r.Err)
}

func (r *ErrorResponse[T]) GetError() error {
	return r.Err
}
