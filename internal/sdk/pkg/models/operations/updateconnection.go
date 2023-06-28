// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"abbey/internal/sdk/pkg/models/shared"
	"net/http"
)

type UpdateConnectionRequest struct {
	ConnectionUpdateParams shared.ConnectionUpdateParams `request:"mediaType=application/json"`
	// The ID of the connection to update
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
}

type UpdateConnectionResponse struct {
	// Success
	Connection  *shared.Connection
	ContentType string
	// Request Failed
	Error       *shared.Error
	StatusCode  int
	RawResponse *http.Response
}