package handlers

import (
	"errors"

	"github.com/go-provider-sdk/internal/clients/rest/httptransport"
	"github.com/go-provider-sdk/internal/unmarshal"
)

type UnmarshalHandler[T any] struct {
	nextHandler Handler[T]
}

func NewUnmarshalHandler[T any]() *UnmarshalHandler[T] {
	return &UnmarshalHandler[T]{
		nextHandler: nil,
	}
}

func (h *UnmarshalHandler[T]) Handle(request httptransport.Request) (*httptransport.Response[T], *httptransport.ErrorResponse[T]) {
	if h.nextHandler == nil {
		err := errors.New("Handler chain terminated without terminating handler")
		return nil, httptransport.NewErrorResponse[T](err, nil)
	}

	resp, handlerError := h.nextHandler.Handle(request)
	if handlerError != nil {
		return nil, handlerError
	}

	target := new(T)
	err := unmarshal.Unmarshal(resp.Body, target)
	if err != nil {
		err := errors.New("failed to unmarshal response body into struct")
		return nil, httptransport.NewErrorResponse[T](err, nil)
	}

	resp.Data = *target

	return resp, nil
}

func (h *UnmarshalHandler[T]) SetNext(handler Handler[T]) {
	h.nextHandler = handler
}
