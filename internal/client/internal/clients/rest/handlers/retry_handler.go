package handlers

import (
	"errors"
	"time"

	"github.com/go-provider-sdk/internal/clients/rest/httptransport"
)

const (
	maxRetries = 3
	retryDelay = 150 * time.Millisecond
)

type RetryHandler[T any] struct {
	nextHandler Handler[T]
}

func NewRetryHandler[T any]() *RetryHandler[T] {
	return &RetryHandler[T]{
		nextHandler: nil,
	}
}

func (h *RetryHandler[T]) Handle(request httptransport.Request) (*httptransport.Response[T], *httptransport.ErrorResponse[T]) {
	if h.nextHandler == nil {
		err := errors.New("Handler chain terminated without terminating handler")
		return nil, httptransport.NewErrorResponse[T](err, nil)
	}

	var err *httptransport.ErrorResponse[T]
	for tryCount := 0; tryCount < maxRetries; tryCount++ {
		nextRequest := request.Clone()

		var resp *httptransport.Response[T]
		resp, err = h.nextHandler.Handle(nextRequest)
		if err != nil {
			return nil, err
		}

		if resp.StatusCode < 400 {
			return resp, nil
		}

		backoffDuration := time.Duration(tryCount) * retryDelay
		time.Sleep(backoffDuration)
	}
	return nil, httptransport.NewErrorResponse[T](err, nil)
}

func (h *RetryHandler[T]) SetNext(handler Handler[T]) {
	h.nextHandler = handler
}
