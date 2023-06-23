// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"abbey/internal/sdk/pkg/models/shared"
	"net/http"
)

type GetGrantKitByIDRequest struct {
	// The ID of the grant kit or resource to retrieve.
	GrantKitIDOrName string `pathParam:"style=simple,explode=false,name=grant_kit_id_or_name"`
}

type GetGrantKitByIDResponse struct {
	ContentType string
	// Authentication Failed
	Error *shared.Error
	// Success
	GrantKit    *shared.GrantKit
	StatusCode  int
	RawResponse *http.Response
}
