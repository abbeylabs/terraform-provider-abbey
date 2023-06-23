// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"abbey/internal/sdk/pkg/models/shared"
	"net/http"
)

type CreateAPIKeyResponse struct {
	// Created
	APIKey      *shared.APIKey
	ContentType string
	// Request Failed
	Error       *shared.Error
	StatusCode  int
	RawResponse *http.Response
}
