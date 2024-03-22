package shared

import (
	"github.com/go-provider-sdk/internal/clients/rest/httptransport"
)

type ClientResponse[T any] struct {
	Data     T
	Metadata ClientResponseMetadata
}

type ClientResponseMetadata struct {
	Headers    map[string]string
	StatusCode int
}

func NewClientResponse[T any](resp *httptransport.Response[T]) *ClientResponse[T] {
	return &ClientResponse[T]{
		Data: resp.Data,
		Metadata: ClientResponseMetadata{
			StatusCode: resp.StatusCode,
			Headers:    resp.Headers,
		},
	}
}
