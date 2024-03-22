package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-provider-sdk/internal/clients/rest/httptransport"
)

type TerminatingHandler[T any] struct {
	httpClient *http.Client
}

func NewTerminatingHandler[T any]() *TerminatingHandler[T] {
	return &TerminatingHandler[T]{
		httpClient: &http.Client{Timeout: time.Second * 10},
	}
}

func (h *TerminatingHandler[T]) Handle(request httptransport.Request) (*httptransport.Response[T], *httptransport.ErrorResponse[T]) {
	requestClone := request.Clone()
	req, err := requestClone.CreateHttpRequest()
	if err != nil {
		return nil, httptransport.NewErrorResponse[T](err, nil)
	}

	resp, err := h.httpClient.Do(req)
	if err != nil {
		return nil, httptransport.NewErrorResponse[T](err, nil)
	}

	transportResponse, responseErr := httptransport.NewResponse[T](resp)
	if responseErr != nil {
		return nil, httptransport.NewErrorResponse[T](responseErr, transportResponse)
	}

	if transportResponse.StatusCode >= 400 {
		err := fmt.Errorf("HTTP request failed with status code %d", transportResponse.StatusCode)
		return nil, httptransport.NewErrorResponse[T](err, transportResponse)
	}

	return transportResponse, nil
}

func (h *TerminatingHandler[T]) SetNext(handler Handler[T]) {
	fmt.Println("WARNING: SetNext should not be called on the terminating handler.")
}
