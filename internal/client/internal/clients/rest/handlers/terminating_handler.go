package handlers

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/go-provider-sdk/internal/clients/rest/httptransport"
)

type TerminatingHandler struct {
	httpClient *http.Client
}

func NewTerminatingHandler() *TerminatingHandler {
	return &TerminatingHandler{
		httpClient: &http.Client{Timeout: time.Second * 10},
	}
}

func (h *TerminatingHandler) Handle(request httptransport.Request) (*httptransport.Response, *httptransport.ErrorResponse) {
	requestClone := request.Clone()
	req, err := requestClone.CreateHttpRequest()
	if err != nil {
		return nil, httptransport.NewErrorResponse(err, nil)
	}

	resp, err := h.httpClient.Do(req)
	if err != nil {
		return nil, httptransport.NewErrorResponse(err, nil)
	}

	return h.handleResponse(resp)
}

func (h *TerminatingHandler) handleResponse(resp *http.Response) (*httptransport.Response, *httptransport.ErrorResponse) {
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, httptransport.NewErrorResponse(err, nil)
	}

	responseHeaders := make(map[string]string)
	for key := range resp.Header {
		responseHeaders[key] = resp.Header.Get(key)
	}

	response := &httptransport.Response{
		StatusCode: resp.StatusCode,
		Headers:    responseHeaders,
		Body:       body,
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 400 {
		err := fmt.Errorf("HTTP request failed with status code %d", resp.StatusCode)
		return nil, httptransport.NewErrorResponse(err, response)
	}
	return response, nil
}

func (h *TerminatingHandler) SetNext(handler Handler) {
	fmt.Println("WARNING: SetNext should not be called on the terminating handler.")
}
