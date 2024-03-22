package handlers

import (
	"errors"

	"github.com/go-provider-sdk/internal/clients/rest/hooks"
	"github.com/go-provider-sdk/internal/clients/rest/httptransport"
)

type HookHandler[T any] struct {
	nextHandler Handler[T]
	hook        hooks.Hook
}

func NewHookHandler[T any](hook hooks.Hook) *HookHandler[T] {
	return &HookHandler[T]{
		hook:        hook,
		nextHandler: nil,
	}
}

func (h *HookHandler[T]) Handle(request httptransport.Request) (*httptransport.Response[T], *httptransport.ErrorResponse[T]) {
	if h.nextHandler == nil {
		err := errors.New("Handler chain terminated without terminating handler")
		return nil, httptransport.NewErrorResponse[T](err, nil)
	}

	clonedReq := request.Clone()
	hookReq := h.hook.BeforeRequest(&clonedReq)

	nextRequest, ok := hookReq.(*httptransport.Request)
	if !ok {
		err := errors.New("hook returned invalid request")
		return nil, httptransport.NewErrorResponse[T](err, nil)
	}

	response, err := h.nextHandler.Handle(*nextRequest)
	if err != nil && err.IsHttpError {
		clonedError := err.Clone()
		hookError := h.hook.OnError(hookReq, &clonedError)
		nextError, ok := hookError.(*httptransport.ErrorResponse[T])
		if !ok {
			err := errors.New("hook returned invalid error")
			return nil, httptransport.NewErrorResponse[T](err, nil)
		}

		return nil, nextError
	} else if err != nil {
		return nil, err
	}

	clonedResp := response.Clone()
	hookResp := h.hook.AfterResponse(hookReq, &clonedResp)
	nextResponse, ok := hookResp.(*httptransport.Response[T])
	if !ok {
		err := errors.New("hook returned invalid response")
		return nil, httptransport.NewErrorResponse[T](err, nil)
	}

	return nextResponse, nil
}

func (h *HookHandler[T]) SetNext(handler Handler[T]) {
	h.nextHandler = handler
}
