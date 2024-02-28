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

type RetryHandler struct {
	nextHandler Handler
}

func NewRetryHandler() *RetryHandler {
	return &RetryHandler{
		nextHandler: nil,
	}
}

func (h *RetryHandler) Handle(request httptransport.Request) (*httptransport.Response, *httptransport.ErrorResponse) {
	if h.nextHandler == nil {
		err := errors.New("Handler chain terminated without terminating handler")
		return nil, httptransport.NewErrorResponse(err, nil)
	}

	var err *httptransport.ErrorResponse
	for tryCount := 0; tryCount < maxRetries; tryCount++ {
		var resp *httptransport.Response
		resp, err = h.nextHandler.Handle(request)
		if err != nil {
			return nil, err
		}

		if resp.StatusCode <= 399 {
			return resp, nil
		}

		backoffDuration := time.Duration(tryCount) * retryDelay
		time.Sleep(backoffDuration)
	}
	return nil, httptransport.NewErrorResponse(err, nil)
}

func (h *RetryHandler) SetNext(handler Handler) {
	h.nextHandler = handler
}
