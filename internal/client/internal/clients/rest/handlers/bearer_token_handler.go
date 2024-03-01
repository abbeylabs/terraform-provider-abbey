package handlers

import (
	"errors"
	"fmt"

	"github.com/go-provider-sdk/internal/clients/rest/httptransport"
)

type AccessTokenHandler struct {
	nextHandler Handler
}

func NewAccessTokenHandler() *AccessTokenHandler {
	return &AccessTokenHandler{
		nextHandler: nil,
	}
}

func (h *AccessTokenHandler) Handle(request httptransport.Request) (*httptransport.Response, *httptransport.ErrorResponse) {
	if h.nextHandler == nil {
		err := errors.New("Handler chain terminated without terminating handler")
		return nil, httptransport.NewErrorResponse(err, nil)
	}

	nextRequest := request.Clone()

	if request.Config.AccessToken == nil {
		return h.nextHandler.Handle(nextRequest)
	}

	nextRequest.SetHeader("Authorization", fmt.Sprintf("Bearer %s", *request.Config.AccessToken))

	if h.nextHandler == nil {
		err := errors.New("Handler chain terminated without terminating handler")
		return nil, httptransport.NewErrorResponse(err, nil)
	}

	return h.nextHandler.Handle(nextRequest)
}

func (h *AccessTokenHandler) SetNext(handler Handler) {
	h.nextHandler = handler
}
