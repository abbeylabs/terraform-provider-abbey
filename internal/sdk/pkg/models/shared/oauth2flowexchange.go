// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// Oauth2FlowExchange - Describes how the client should conduct the authorization code exchange.
type Oauth2FlowExchange struct {
	Enabled bool        `json:"enabled"`
	Request RequestSpec `json:"request"`
}