// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"abbey/v2/internal/sdk/pkg/models/shared"
	"net/http"
)

type DeleteAPIKeyRequest struct {
	// The API Key to delete.
	APIKey string `pathParam:"style=simple,explode=false,name=api_key"`
}

type DeleteAPIKeyResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Authentication Failed
	Error *shared.Error
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}
