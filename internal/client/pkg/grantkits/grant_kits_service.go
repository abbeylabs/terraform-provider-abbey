package grantkits

import (
	restClient "github.com/go-provider-sdk/internal/clients/rest"
	"github.com/go-provider-sdk/internal/clients/rest/httptransport"
	"github.com/go-provider-sdk/internal/configmanager"
	"github.com/go-provider-sdk/internal/unmarshal"
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
func (api *GrantKitsService) ListGrantKits() (*shared.ClientResponse[[]GrantKit], error) {
	config := *api.getConfig()

	client := restClient.NewRestClient(config)

	request := httptransport.NewRequest("GET", "/grant-kits", config)

	httpResponse, err := client.Call(request)
	if err != nil {
		return nil, err.GetError()
	}

	data, unmarshalError := unmarshal.ToArray[[]GrantKit](httpResponse)
	if unmarshalError != nil {
		return nil, unmarshalError
	}

	response := shared.ClientResponse[[]GrantKit]{
		Data: data,
		Metadata: shared.ClientResponseMetadata{
			Headers:    httpResponse.Headers,
			StatusCode: httpResponse.StatusCode,
		},
	}

	return &response, nil
}

// Creates a new Grant Kit
func (api *GrantKitsService) CreateGrantKit(grantKitCreateParams GrantKitCreateParams) (*shared.ClientResponse[GrantKit], error) {
	config := *api.getConfig()

	client := restClient.NewRestClient(config)

	request := httptransport.NewRequest("POST", "/grant-kits", config)

	request.Body = grantKitCreateParams

	httpResponse, err := client.Call(request)
	if err != nil {
		return nil, err.GetError()
	}

	data, unmarshalError := unmarshal.ToObject[GrantKit](httpResponse)
	if unmarshalError != nil {
		return nil, unmarshalError
	}

	response := shared.ClientResponse[GrantKit]{
		Data: *data,
		Metadata: shared.ClientResponseMetadata{
			Headers:    httpResponse.Headers,
			StatusCode: httpResponse.StatusCode,
		},
	}

	return &response, nil
}

// Returns the details of a Grant Kit.
func (api *GrantKitsService) GetGrantKitById(grantKitIdOrName string) (*shared.ClientResponse[GrantKit], error) {
	config := *api.getConfig()

	client := restClient.NewRestClient(config)

	request := httptransport.NewRequest("GET", "/grant-kits/{grant_kit_id_or_name}", config)

	request.SetPathParam("grant_kit_id_or_name", grantKitIdOrName)

	httpResponse, err := client.Call(request)
	if err != nil {
		return nil, err.GetError()
	}

	data, unmarshalError := unmarshal.ToObject[GrantKit](httpResponse)
	if unmarshalError != nil {
		return nil, unmarshalError
	}

	response := shared.ClientResponse[GrantKit]{
		Data: *data,
		Metadata: shared.ClientResponseMetadata{
			Headers:    httpResponse.Headers,
			StatusCode: httpResponse.StatusCode,
		},
	}

	return &response, nil
}

// Updates the specified grant kit.
func (api *GrantKitsService) UpdateGrantKit(grantKitIdOrName string, grantKitUpdateParams GrantKitUpdateParams) (*shared.ClientResponse[GrantKit], error) {
	config := *api.getConfig()

	client := restClient.NewRestClient(config)

	request := httptransport.NewRequest("PUT", "/grant-kits/{grant_kit_id_or_name}", config)

	request.Body = grantKitUpdateParams

	request.SetPathParam("grant_kit_id_or_name", grantKitIdOrName)

	httpResponse, err := client.Call(request)
	if err != nil {
		return nil, err.GetError()
	}

	data, unmarshalError := unmarshal.ToObject[GrantKit](httpResponse)
	if unmarshalError != nil {
		return nil, unmarshalError
	}

	response := shared.ClientResponse[GrantKit]{
		Data: *data,
		Metadata: shared.ClientResponseMetadata{
			Headers:    httpResponse.Headers,
			StatusCode: httpResponse.StatusCode,
		},
	}

	return &response, nil
}

// Deletes the specified grant kit.
func (api *GrantKitsService) DeleteGrantKit(grantKitIdOrName string) (*shared.ClientResponse[GrantKit], error) {
	config := *api.getConfig()

	client := restClient.NewRestClient(config)

	request := httptransport.NewRequest("DELETE", "/grant-kits/{grant_kit_id_or_name}", config)

	request.SetPathParam("grant_kit_id_or_name", grantKitIdOrName)

	httpResponse, err := client.Call(request)
	if err != nil {
		return nil, err.GetError()
	}

	data, unmarshalError := unmarshal.ToObject[GrantKit](httpResponse)
	if unmarshalError != nil {
		return nil, unmarshalError
	}

	response := shared.ClientResponse[GrantKit]{
		Data: *data,
		Metadata: shared.ClientResponseMetadata{
			Headers:    httpResponse.Headers,
			StatusCode: httpResponse.StatusCode,
		},
	}

	return &response, nil
}
