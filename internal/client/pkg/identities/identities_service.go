package identities

import (
	restClient "github.com/go-provider-sdk/internal/clients/rest"
	"github.com/go-provider-sdk/internal/clients/rest/httptransport"
	"github.com/go-provider-sdk/internal/configmanager"
	"github.com/go-provider-sdk/internal/unmarshal"
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
func (api *IdentitiesService) ListEnrichedIdentities() (*shared.ClientResponse[[]Identity], error) {
	config := *api.getConfig()

	client := restClient.NewRestClient(config)

	request := httptransport.NewRequest("GET", "/identities", config)

	httpResponse, err := client.Call(request)
	if err != nil {
		return nil, err.GetError()
	}

	data, unmarshalError := unmarshal.ToArray[[]Identity](httpResponse)
	if unmarshalError != nil {
		return nil, unmarshalError
	}

	response := shared.ClientResponse[[]Identity]{
		Data: data,
		Metadata: shared.ClientResponseMetadata{
			Headers:    httpResponse.Headers,
			StatusCode: httpResponse.StatusCode,
		},
	}

	return &response, nil
}

// Creates a new identity.
//
// An identity represents a human, service, or workload.
func (api *IdentitiesService) CreateIdentity(identityParams IdentityParams) (*shared.ClientResponse[Identity], error) {
	config := *api.getConfig()

	client := restClient.NewRestClient(config)

	request := httptransport.NewRequest("POST", "/identities", config)

	request.Body = identityParams

	httpResponse, err := client.Call(request)
	if err != nil {
		return nil, err.GetError()
	}

	data, unmarshalError := unmarshal.ToObject[Identity](httpResponse)
	if unmarshalError != nil {
		return nil, unmarshalError
	}

	response := shared.ClientResponse[Identity]{
		Data: *data,
		Metadata: shared.ClientResponseMetadata{
			Headers:    httpResponse.Headers,
			StatusCode: httpResponse.StatusCode,
		},
	}

	return &response, nil
}

// Returns the details of an identity.
func (api *IdentitiesService) GetIdentity(identityId string) (*shared.ClientResponse[Identity], error) {
	config := *api.getConfig()

	client := restClient.NewRestClient(config)

	request := httptransport.NewRequest("GET", "/identities/{identity_id}", config)

	request.SetPathParam("identity_id", identityId)

	httpResponse, err := client.Call(request)
	if err != nil {
		return nil, err.GetError()
	}

	data, unmarshalError := unmarshal.ToObject[Identity](httpResponse)
	if unmarshalError != nil {
		return nil, unmarshalError
	}

	response := shared.ClientResponse[Identity]{
		Data: *data,
		Metadata: shared.ClientResponseMetadata{
			Headers:    httpResponse.Headers,
			StatusCode: httpResponse.StatusCode,
		},
	}

	return &response, nil
}

// Updates an identity.
func (api *IdentitiesService) UpdateIdentity(identityId string, identityParams IdentityParams) (*shared.ClientResponse[Identity], error) {
	config := *api.getConfig()

	client := restClient.NewRestClient(config)

	request := httptransport.NewRequest("PUT", "/identities/{identity_id}", config)

	request.Body = identityParams

	request.SetPathParam("identity_id", identityId)

	httpResponse, err := client.Call(request)
	if err != nil {
		return nil, err.GetError()
	}

	data, unmarshalError := unmarshal.ToObject[Identity](httpResponse)
	if unmarshalError != nil {
		return nil, unmarshalError
	}

	response := shared.ClientResponse[Identity]{
		Data: *data,
		Metadata: shared.ClientResponseMetadata{
			Headers:    httpResponse.Headers,
			StatusCode: httpResponse.StatusCode,
		},
	}

	return &response, nil
}

// Deletes the specified identity.
func (api *IdentitiesService) DeleteIdentity(identityId string) error {
	config := *api.getConfig()

	client := restClient.NewRestClient(config)

	request := httptransport.NewRequest("DELETE", "/identities/{identity_id}", config)

	request.SetPathParam("identity_id", identityId)

	_, err := client.Call(request)
	if err != nil {
		return err.GetError()
	}

	return nil
}
