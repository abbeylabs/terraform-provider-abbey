package handlers

import (
	"errors"

	"github.com/go-provider-sdk/internal/clients/rest/hooks"
	"github.com/go-provider-sdk/internal/clients/rest/httptransport"
)

type HookHandler struct {
	nextHandler Handler
	hook        hooks.Hook
}

func NewHookHandler(hook hooks.Hook) *HookHandler {
	return &HookHandler{
		hook:        hook,
		nextHandler: nil,
	}
}

func (h *HookHandler) Handle(request httptransport.Request) (*httptransport.Response, *httptransport.ErrorResponse) {
	clonedReq := request.Clone()
	hookReq := h.hook.BeforeRequest(&clonedReq)

	if h.nextHandler == nil {
		err := errors.New("Handler chain terminated without terminating handler")
		return nil, httptransport.NewErrorResponse(err, nil)
	}

	nextRequest, ok := hookReq.(*httptransport.Request)
	if !ok {
		err := errors.New("hook returned invalid request")
		return nil, httptransport.NewErrorResponse(err, nil)
	}

	response, err := h.nextHandler.Handle(*nextRequest)
	if err != nil && err.IsHttpError {
		clonedError := err.Clone()
		hookError := h.hook.OnError(hookReq, &clonedError)
		nextError, ok := hookError.(*httptransport.ErrorResponse)
		if !ok {
			err := errors.New("hook returned invalid error")
			return nil, httptransport.NewErrorResponse(err, nil)
		}

		return nil, nextError
	} else if err != nil {
		return nil, err
	}

	clonedResp := response.Clone()
	hookResp := h.hook.AfterResponse(hookReq, &clonedResp)
	nextResponse, ok := hookResp.(*httptransport.Response)
	if !ok {
		err := errors.New("hook returned invalid response")
		return nil, httptransport.NewErrorResponse(err, nil)
	}

	return nextResponse, nil
}

func (h *HookHandler) SetNext(handler Handler) {
	h.nextHandler = handler
}
