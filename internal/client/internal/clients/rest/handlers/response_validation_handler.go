package handlers

import (
	"errors"

	"github.com/go-provider-sdk/internal/clients/rest/httptransport"
	"github.com/go-provider-sdk/internal/validation"
)

type ResponseValidationHandler[T any] struct {
	nextHandler Handler[T]
}

func NewResponseValidationHandler[T any]() *ResponseValidationHandler[T] {
	return &ResponseValidationHandler[T]{
		nextHandler: nil,
	}
}

func (h *ResponseValidationHandler[T]) Handle(request httptransport.Request) (*httptransport.Response[T], *httptransport.ErrorResponse[T]) {
	if h.nextHandler == nil {
		err := errors.New("Handler chain terminated without terminating handler")
		return nil, httptransport.NewErrorResponse[T](err, nil)
	}

	resp, handlerError := h.nextHandler.Handle(request)
	if handlerError != nil {
		return nil, handlerError
	}

	err := validation.ValidateData(resp.Data)
	if err != nil {
		return nil, httptransport.NewErrorResponse[T](err, nil)
	}

	return resp, nil
}

func (h *ResponseValidationHandler[T]) SetNext(handler Handler[T]) {
	h.nextHandler = handler
}
