// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"abbey/internal/sdk/pkg/models/shared"
	"net/http"
)

type ListConnectionSpecsResponse struct {
	// Success
	ConnectionSpecListing *shared.ConnectionSpecListing
	ContentType           string
	// Authentication Failed
	Error       *shared.Error
	StatusCode  int
	RawResponse *http.Response
}
