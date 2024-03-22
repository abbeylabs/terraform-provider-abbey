package grantkits

import (
	"context"
	restClient "github.com/go-provider-sdk/internal/clients/rest"
	"github.com/go-provider-sdk/internal/clients/rest/httptransport"
	"github.com/go-provider-sdk/internal/configmanager"
	"github.com/go-provider-sdk/pkg/clientconfig"
	"github.com/go-provider-sdk/pkg/shared"
)

type GrantKitsService struct {
	manager *configmanager.ConfigManager
}

func NewGrantKitsService(manager *configmanager.ConfigManager) *GrantKitsService {
	return &GrantKitsService{
		manager: manager,
	}
}

func (api *GrantKitsService) getConfig() *clientconfig.Config {
	return api.manager.GetGrantKits()
}

func (api *GrantKitsService) SetBaseUrl(baseUrl string) {
	config := api.getConfig()
	config.SetBaseUrl(baseUrl)
}

func (api *GrantKitsService) SetAccessToken(accessToken string) {
	config := api.getConfig()
	config.SetAccessToken(accessToken)
}

// Returns a list of the latest versions of each grant kit in the organization.
//
// Grant Kits are sorted by creation date, descending.
func (api *GrantKitsService) ListGrantKits(ctx context.Context) (*shared.ClientResponse[[]GrantKit], *shared.ClientError) {
	config := *api.getConfig()

	client := restClient.NewRestClient[[]GrantKit](config)

	request := httptransport.NewRequest(ctx, "GET", "/grant-kits", config)

	resp, err := client.Call(request)
	if err != nil {
		return nil, shared.NewClientError[[]GrantKit](err)
	}

	return shared.NewClientResponse[[]GrantKit](resp), nil
}

// Creates a new Grant Kit
func (api *GrantKitsService) CreateGrantKit(ctx context.Context, grantKitCreateParams GrantKitCreateParams) (*shared.ClientResponse[GrantKit], *shared.ClientError) {
	config := *api.getConfig()

	client := restClient.NewRestClient[GrantKit](config)

	request := httptransport.NewRequest(ctx, "POST", "/grant-kits", config)

	request.Body = grantKitCreateParams

	resp, err := client.Call(request)
	if err != nil {
		return nil, shared.NewClientError[GrantKit](err)
	}

	return shared.NewClientResponse[GrantKit](resp), nil
}

// Returns the details of a Grant Kit.
func (api *GrantKitsService) GetGrantKitById(ctx context.Context, grantKitIdOrName string) (*shared.ClientResponse[GrantKit], *shared.ClientError) {
	config := *api.getConfig()

	client := restClient.NewRestClient[GrantKit](config)

	request := httptransport.NewRequest(ctx, "GET", "/grant-kits/{grant_kit_id_or_name}", config)

	request.SetPathParam("grant_kit_id_or_name", grantKitIdOrName)

	resp, err := client.Call(request)
	if err != nil {
		return nil, shared.NewClientError[GrantKit](err)
	}

	return shared.NewClientResponse[GrantKit](resp), nil
}

// Updates the specified grant kit.
func (api *GrantKitsService) UpdateGrantKit(ctx context.Context, grantKitIdOrName string, grantKitUpdateParams GrantKitUpdateParams) (*shared.ClientResponse[GrantKit], *shared.ClientError) {
	config := *api.getConfig()

	client := restClient.NewRestClient[GrantKit](config)

	request := httptransport.NewRequest(ctx, "PUT", "/grant-kits/{grant_kit_id_or_name}", config)

	request.Body = grantKitUpdateParams

	request.SetPathParam("grant_kit_id_or_name", grantKitIdOrName)

	resp, err := client.Call(request)
	if err != nil {
		return nil, shared.NewClientError[GrantKit](err)
	}

	return shared.NewClientResponse[GrantKit](resp), nil
}

// Deletes the specified grant kit.
func (api *GrantKitsService) DeleteGrantKit(ctx context.Context, grantKitIdOrName string) (*shared.ClientResponse[GrantKit], *shared.ClientError) {
	config := *api.getConfig()

	client := restClient.NewRestClient[GrantKit](config)

	request := httptransport.NewRequest(ctx, "DELETE", "/grant-kits/{grant_kit_id_or_name}", config)

	request.SetPathParam("grant_kit_id_or_name", grantKitIdOrName)

	resp, err := client.Call(request)
	if err != nil {
		return nil, shared.NewClientError[GrantKit](err)
	}

	return shared.NewClientResponse[GrantKit](resp), nil
}
