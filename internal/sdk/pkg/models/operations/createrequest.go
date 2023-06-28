// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"abbey/internal/sdk/pkg/models/shared"
	"net/http"
)

type CreateRequestResponse struct {
	ContentType string
	// Request Failed
	Error *shared.Error
	// Created
	Request     *shared.Request
	StatusCode  int
	RawResponse *http.Response
}