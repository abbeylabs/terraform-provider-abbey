package handlers

import (
	"errors"
	"fmt"

	"github.com/go-provider-sdk/internal/clients/rest/httptransport"
)

type AccessTokenHandler[T any] struct {
	nextHandler Handler[T]
}

func NewAccessTokenHandler[T any]() *AccessTokenHandler[T] {
	return &AccessTokenHandler[T]{
		nextHandler: nil,
	}
}

func (h *AccessTokenHandler[T]) Handle(request httptransport.Request) (*httptransport.Response[T], *httptransport.ErrorResponse[T]) {
	if h.nextHandler == nil {
		err := errors.New("Handler chain terminated without terminating handler")
		return nil, httptransport.NewErrorResponse[T](err, nil)
	}

	nextRequest := request.Clone()
	if request.Config.AccessToken == nil {
		return h.nextHandler.Handle(nextRequest)
	}

	nextRequest.SetHeader("Authorization", fmt.Sprintf("Bearer %s", *request.Config.AccessToken))

	return h.nextHandler.Handle(nextRequest)
}

func (h *AccessTokenHandler[T]) SetNext(handler Handler[T]) {
	h.nextHandler = handler
}
