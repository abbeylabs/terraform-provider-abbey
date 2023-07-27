// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package sdk

import (
	"abbey/v2/internal/sdk/pkg/models/shared"
	"abbey/v2/internal/sdk/pkg/utils"
	"fmt"
	"net/http"
	"time"
)

const (
	// ServerProd - prod
	ServerProd string = "prod"
	// ServerDev - dev
	ServerDev string = "dev"
)

// ServerList contains the list of servers available to the SDK
var ServerList = map[string]string{
	ServerProd: "https://api.abbey.io/v1",
	ServerDev:  "http://localhost:8080/v1",
}

// HTTPClient provides an interface for suplying the SDK with a custom HTTP client
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// String provides a helper function to return a pointer to a string
func String(s string) *string { return &s }

// Bool provides a helper function to return a pointer to a bool
func Bool(b bool) *bool { return &b }

// Int provides a helper function to return a pointer to an int
func Int(i int) *int { return &i }

// Int64 provides a helper function to return a pointer to an int64
func Int64(i int64) *int64 { return &i }

// Float32 provides a helper function to return a pointer to a float32
func Float32(f float32) *float32 { return &f }

// Float64 provides a helper function to return a pointer to a float64
func Float64(f float64) *float64 { return &f }

type sdkConfiguration struct {
	DefaultClient     HTTPClient
	SecurityClient    HTTPClient
	Security          *shared.Security
	ServerURL         string
	Server            string
	Language          string
	OpenAPIDocVersion string
	SDKVersion        string
	GenVersion        string
}

func (c *sdkConfiguration) GetServerDetails() (string, map[string]string) {
	if c.ServerURL != "" {
		return c.ServerURL, nil
	}

	if c.Server == "" {
		c.Server = "prod"
	}

	return ServerList[c.Server], nil
}

// SDK - Abbey API: The public Abbey API. Used for integrating with Abbey and building interfaces to extend the Abbey platform.
// See https://docs.abbey.io for more information.
type SDK struct {
	// APIKeys - API Keys are used to authenticate to the Abbey API.
	//
	// https://docs.abbey.io/product/managing-api-keys
	APIKeys *apiKeys
	// ConnectionSpecs - Connection Specs are the templates for creating connections.
	// They are used to validate connection parameters and to provide a UI for creating connections.
	//
	// https://docs.abbey.io
	ConnectionSpecs *connectionSpecs
	// Connections - Connections are authenticated, with scopes if available, and made available to Abbey Grant Kits at runtime.
	//
	// https://docs.abbey.io
	Connections *connections
	// Demo - Abbey Demo
	// https://docs.abbey.io/getting-started/quickstart
	Demo *demo
	// GrantKits - Grant Kits are what you configure in code to control and automatically right-size permissions for resources.
	// A Grant Kit has 3 components:
	//
	// 1. Workflow to configure how someone should get access.
	// 2. Policies to configure if someone should get access.
	// 3. Output to configure how and where Grants should materialize.
	//
	// https://docs.abbey.io/getting-started/concepts#grant-kits
	GrantKits *grantKits
	// Grants - Grants are permissions that reflect the result of an access request going through the process of evaluating
	// policies and approval workflows where all approval conditions are met.
	//
	// Grants may be revoked manually by a user or automatically if a time-based or attribute-based policy is
	// included in the corresponding Grant Kit's policy.
	//
	// https://docs.abbey.io/getting-started/concepts#grants
	Grants *grants
	// Identities - User metadata used for enriching data.
	// Enriched data is used to write richer policies, workflows, and outputs.
	//
	// https://docs.abbey.io
	Identities *identities
	// Requests - Requests are Access Requests that users make to get access to a resource.
	//
	// https://docs.abbey.io/getting-started/concepts#access-requests
	Requests *requests
	// Reviews - Reviews are decisions made by a reviewer on an Access Request.
	//
	// A Reviewer might approve or deny a request.
	//
	// https://docs.abbey.io/product/approving-or-denying-access-requests
	Reviews *reviews

	sdkConfiguration sdkConfiguration
}

type SDKOption func(*SDK)

// WithServerURL allows the overriding of the default server URL
func WithServerURL(serverURL string) SDKOption {
	return func(sdk *SDK) {
		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithTemplatedServerURL allows the overriding of the default server URL with a templated URL populated with the provided parameters
func WithTemplatedServerURL(serverURL string, params map[string]string) SDKOption {
	return func(sdk *SDK) {
		if params != nil {
			serverURL = utils.ReplaceParameters(serverURL, params)
		}

		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithServer allows the overriding of the default server by name
func WithServer(server string) SDKOption {
	return func(sdk *SDK) {
		_, ok := ServerList[server]
		if !ok {
			panic(fmt.Errorf("server %s not found", server))
		}

		sdk.sdkConfiguration.Server = server
	}
}

// WithClient allows the overriding of the default HTTP client used by the SDK
func WithClient(client HTTPClient) SDKOption {
	return func(sdk *SDK) {
		sdk.sdkConfiguration.DefaultClient = client
	}
}

// WithSecurity configures the SDK to use the provided security details
func WithSecurity(security shared.Security) SDKOption {
	return func(sdk *SDK) {
		sdk.sdkConfiguration.Security = &security
	}
}

// New creates a new instance of the SDK with the provided options
func New(opts ...SDKOption) *SDK {
	sdk := &SDK{
		sdkConfiguration: sdkConfiguration{
			Language:          "terraform",
			OpenAPIDocVersion: "v1",
			SDKVersion:        "2.1.3",
			GenVersion:        "2.73.1",
		},
	}
	for _, opt := range opts {
		opt(sdk)
	}

	// Use WithClient to override the default client if you would like to customize the timeout
	if sdk.sdkConfiguration.DefaultClient == nil {
		sdk.sdkConfiguration.DefaultClient = &http.Client{Timeout: 60 * time.Second}
	}
	if sdk.sdkConfiguration.SecurityClient == nil {
		if sdk.sdkConfiguration.Security != nil {
			sdk.sdkConfiguration.SecurityClient = utils.ConfigureSecurityClient(sdk.sdkConfiguration.DefaultClient, sdk.sdkConfiguration.Security)
		} else {
			sdk.sdkConfiguration.SecurityClient = sdk.sdkConfiguration.DefaultClient
		}
	}

	sdk.APIKeys = newAPIKeys(sdk.sdkConfiguration)

	sdk.ConnectionSpecs = newConnectionSpecs(sdk.sdkConfiguration)

	sdk.Connections = newConnections(sdk.sdkConfiguration)

	sdk.Demo = newDemo(sdk.sdkConfiguration)

	sdk.GrantKits = newGrantKits(sdk.sdkConfiguration)

	sdk.Grants = newGrants(sdk.sdkConfiguration)

	sdk.Identities = newIdentities(sdk.sdkConfiguration)

	sdk.Requests = newRequests(sdk.sdkConfiguration)

	sdk.Reviews = newReviews(sdk.sdkConfiguration)

	return sdk
}
