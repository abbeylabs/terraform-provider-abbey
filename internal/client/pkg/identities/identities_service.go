package identities

import (
	"context"
	restClient "github.com/go-provider-sdk/internal/clients/rest"
	"github.com/go-provider-sdk/internal/clients/rest/httptransport"
	"github.com/go-provider-sdk/internal/configmanager"
	"github.com/go-provider-sdk/pkg/clientconfig"
	"github.com/go-provider-sdk/pkg/shared"
)

type IdentitiesService struct {
	manager *configmanager.ConfigManager
}

func NewIdentitiesService(manager *configmanager.ConfigManager) *IdentitiesService {
	return &IdentitiesService{
		manager: manager,
	}
}

func (api *IdentitiesService) getConfig() *clientconfig.Config {
	return api.manager.GetIdentities()
}

func (api *IdentitiesService) SetBaseUrl(baseUrl string) {
	config := api.getConfig()
	config.SetBaseUrl(baseUrl)
}

func (api *IdentitiesService) SetAccessToken(accessToken string) {
	config := api.getConfig()
	config.SetAccessToken(accessToken)
}

// Returns all Identities with enriched metadata in the org
func (api *IdentitiesService) ListEnrichedIdentities(ctx context.Context) (*shared.ClientResponse[[]Identity], *shared.ClientError) {
	config := *api.getConfig()

	client := restClient.NewRestClient[[]Identity](config)

	request := httptransport.NewRequest(ctx, "GET", "/identities", config)

	resp, err := client.Call(request)
	if err != nil {
		return nil, shared.NewClientError[[]Identity](err)
	}

	return shared.NewClientResponse[[]Identity](resp), nil
}

// Creates a new identity.
//
// An identity represents a human, service, or workload.
func (api *IdentitiesService) CreateIdentity(ctx context.Context, identityParams IdentityParams) (*shared.ClientResponse[Identity], *shared.ClientError) {
	config := *api.getConfig()

	client := restClient.NewRestClient[Identity](config)

	request := httptransport.NewRequest(ctx, "POST", "/identities", config)

	request.Body = identityParams

	resp, err := client.Call(request)
	if err != nil {
		return nil, shared.NewClientError[Identity](err)
	}

	return shared.NewClientResponse[Identity](resp), nil
}

// Returns the details of an identity.
func (api *IdentitiesService) GetIdentity(ctx context.Context, identityId string) (*shared.ClientResponse[Identity], *shared.ClientError) {
	config := *api.getConfig()

	client := restClient.NewRestClient[Identity](config)

	request := httptransport.NewRequest(ctx, "GET", "/identities/{identity_id}", config)

	request.SetPathParam("identity_id", identityId)

	resp, err := client.Call(request)
	if err != nil {
		return nil, shared.NewClientError[Identity](err)
	}

	return shared.NewClientResponse[Identity](resp), nil
}

// Updates an identity.
func (api *IdentitiesService) UpdateIdentity(ctx context.Context, identityId string, identityParams IdentityParams) (*shared.ClientResponse[Identity], *shared.ClientError) {
	config := *api.getConfig()

	client := restClient.NewRestClient[Identity](config)

	request := httptransport.NewRequest(ctx, "PUT", "/identities/{identity_id}", config)

	request.Body = identityParams

	request.SetPathParam("identity_id", identityId)

	resp, err := client.Call(request)
	if err != nil {
		return nil, shared.NewClientError[Identity](err)
	}

	return shared.NewClientResponse[Identity](resp), nil
}

// Deletes the specified identity.
func (api *IdentitiesService) DeleteIdentity(ctx context.Context, identityId string) (*shared.ClientResponse[any], *shared.ClientError) {
	config := *api.getConfig()

	client := restClient.NewRestClient[any](config)

	request := httptransport.NewRequest(ctx, "DELETE", "/identities/{identity_id}", config)

	request.SetPathParam("identity_id", identityId)

	resp, err := client.Call(request)
	if err != nil {
		return nil, shared.NewClientError[any](err)
	}

	return shared.NewClientResponse[any](resp), nil
}
