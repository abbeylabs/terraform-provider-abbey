package shared

import (
	"github.com/go-provider-sdk/internal/clients/rest/httptransport"
)

type ClientError struct {
	Err      error
	Body     []byte
	Metadata ClientErrorMetadata
}

type ClientErrorMetadata struct {
	Headers    map[string]string
	StatusCode int
}

func NewClientError[T any](transportError *httptransport.ErrorResponse[T]) *ClientError {
	return &ClientError{
		Err:  transportError.GetError(),
		Body: transportError.GetBody(),
		Metadata: ClientErrorMetadata{
			StatusCode: transportError.GetStatusCode(),
			Headers:    transportError.GetHeaders(),
		},
	}
}

func (e *ClientError) Error() string {
	return e.Err.Error()
}
