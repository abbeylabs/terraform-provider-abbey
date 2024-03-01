package handlers

import (
	"errors"

	"github.com/go-provider-sdk/internal/clients/rest/httptransport"
	"github.com/go-provider-sdk/internal/validation"
)

type RequestValidationHandler struct {
	nextHandler Handler
}

func NewRequestValidationHandler() *RequestValidationHandler {
	return &RequestValidationHandler{
		nextHandler: nil,
	}
}

func (h *RequestValidationHandler) Handle(request httptransport.Request) (*httptransport.Response, *httptransport.ErrorResponse) {
	if h.nextHandler == nil {
		err := errors.New("Handler chain terminated without terminating handler")
		return nil, httptransport.NewErrorResponse(err, nil)
	}

	err := validation.ValidateData(request.Body)
	if err != nil {
		return nil, httptransport.NewErrorResponse(err, nil)
	}

	err = validation.ValidateData(request.Options)
	if err != nil {
		return nil, httptransport.NewErrorResponse(err, nil)
	}

	return h.nextHandler.Handle(request)
}

func (h *RequestValidationHandler) SetNext(handler Handler) {
	h.nextHandler = handler
}
