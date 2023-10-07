// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"abbey/v2/internal/sdk/pkg/models/shared"
	"net/http"
)

type CreateGHEAppResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Request Failed
	Error *shared.Error
	// Created
	GithubEnterpriseApp *shared.GithubEnterpriseApp
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}
