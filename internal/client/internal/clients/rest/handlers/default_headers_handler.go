package handlers

import (
	"errors"

	"github.com/go-provider-sdk/internal/clients/rest/httptransport"
)

type DefaultHeadersHandler[T any] struct {
	defaultHeaders map[string]string
	nextHandler    Handler[T]
}

func NewDefaultHeadersHandler[T any]() *DefaultHeadersHandler[T] {
	defaultHeaders := map[string]string{
		"User-Agent":   "liblab/2.0.20 go/1.18",
		"Content-type": "application/json",
	}

	return &DefaultHeadersHandler[T]{
		defaultHeaders: defaultHeaders,
		nextHandler:    nil,
	}
}

func (h *DefaultHeadersHandler[T]) Handle(request httptransport.Request) (*httptransport.Response[T], *httptransport.ErrorResponse[T]) {
	if h.nextHandler == nil {
		err := errors.New("Handler chain terminated without terminating handler")
		return nil, httptransport.NewErrorResponse[T](err, nil)
	}

	nextRequest := request.Clone()

	for key, value := range h.defaultHeaders {
		nextRequest.SetHeader(key, value)
	}

	return h.nextHandler.Handle(nextRequest)
}

func (h *DefaultHeadersHandler[T]) SetNext(handler Handler[T]) {
	h.nextHandler = handler
}
