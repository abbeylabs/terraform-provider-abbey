// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"abbey/v2/internal/sdk/pkg/models/shared"
	"net/http"
)

type ListGrantKitVersionsByIDRequest struct {
	// The ID of the grant kit or resource to retrieve.
	GrantKitIDOrName string `pathParam:"style=simple,explode=false,name=grant_kit_id_or_name"`
}

type ListGrantKitVersionsByIDResponse struct {
	ContentType string
	// Authentication Failed
	Error *shared.Error
	// Success
	GrantKitVersions []shared.GrantKitVersion
	StatusCode       int
	RawResponse      *http.Response
}
