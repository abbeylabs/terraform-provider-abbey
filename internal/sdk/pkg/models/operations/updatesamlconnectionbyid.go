// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"abbey/v2/internal/sdk/pkg/models/shared"
	"net/http"
)

type UpdateSamlConnectionByIDRequest struct {
	SamlConnectionUpdateParams shared.SamlConnectionUpdateParams `request:"mediaType=application/json"`
	// The ID of the SAML connection to update
	SamlConnectionID string `pathParam:"style=simple,explode=false,name=saml_connection_id"`
}

type UpdateSamlConnectionByIDResponse struct {
	ContentType string
	// Authentication Failed
	Error *shared.Error
	// Success
	SamlConnection *shared.SamlConnection
	StatusCode     int
	RawResponse    *http.Response
}
