// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"abbey/v2/internal/sdk/pkg/models/shared"
	"net/http"
)

type GetConnectionRequest struct {
	// The ID of the connection to retrieve
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
}

type GetConnectionResponse struct {
	// Success
	Connection *shared.Connection
	// HTTP response content type for this operation
	ContentType string
	// Authentication Failed
	Error *shared.Error
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}
