// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"abbey/v2/internal/sdk/pkg/models/shared"
	"net/http"
)

type UpdateIdentityRequest struct {
	IdentityParams shared.IdentityParams `request:"mediaType=application/json"`
	// The ID of the identity to retrieve
	IdentityID string `pathParam:"style=simple,explode=false,name=identity_id"`
}

type UpdateIdentityResponse struct {
	ContentType string
	// Authentication Failed
	Error *shared.Error
	// Success
	Identity    *shared.Identity
	StatusCode  int
	RawResponse *http.Response
}