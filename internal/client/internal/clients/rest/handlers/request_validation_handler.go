package handlers

import (
	"errors"

	"github.com/go-provider-sdk/internal/clients/rest/httptransport"
	"github.com/go-provider-sdk/internal/validation"
)

type RequestValidationHandler[T any] struct {
	nextHandler Handler[T]
}

func NewRequestValidationHandler[T any]() *RequestValidationHandler[T] {
	return &RequestValidationHandler[T]{
		nextHandler: nil,
	}
}

func (h *RequestValidationHandler[T]) Handle(request httptransport.Request) (*httptransport.Response[T], *httptransport.ErrorResponse[T]) {
	if h.nextHandler == nil {
		err := errors.New("Handler chain terminated without terminating handler")
		return nil, httptransport.NewErrorResponse[T](err, nil)
	}

	err := validation.ValidateData(request.Body)
	if err != nil {
		return nil, httptransport.NewErrorResponse[T](err, nil)
	}

	err = validation.ValidateData(request.Options)
	if err != nil {
		return nil, httptransport.NewErrorResponse[T](err, nil)
	}

	return h.nextHandler.Handle(request)
}

func (h *RequestValidationHandler[T]) SetNext(handler Handler[T]) {
	h.nextHandler = handler
}
