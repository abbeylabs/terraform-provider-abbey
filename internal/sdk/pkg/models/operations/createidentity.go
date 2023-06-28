// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"abbey/internal/sdk/pkg/models/shared"
	"net/http"
)

type CreateIdentityResponse struct {
	ContentType string
	// Request Failed
	Error *shared.Error
	// Created
	Identity    *shared.Identity
	StatusCode  int
	RawResponse *http.Response
}