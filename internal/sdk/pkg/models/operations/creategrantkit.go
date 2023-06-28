// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"abbey/internal/sdk/pkg/models/shared"
	"net/http"
)

type CreateGrantKitResponse struct {
	ContentType string
	// Request Failed
	Error *shared.Error
	// Created
	GrantKit    *shared.GrantKit
	StatusCode  int
	RawResponse *http.Response
}