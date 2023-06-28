// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type Oauth2Flow struct {
	CallbackQueryParams []string `json:"callback_query_params,omitempty"`
	// Describes how the client should conduct the authorization code exchange.
	//
	Exchange    *Oauth2FlowExchange `json:"exchange,omitempty"`
	Pkce        *Oauth2FlowPkce     `json:"pkce,omitempty"`
	QueryParams []KeyValuePair      `json:"query_params,omitempty"`
	Type        ConnectionAuthType  `json:"type"`
	URL         string              `json:"url"`
}