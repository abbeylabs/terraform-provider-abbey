package demo

import (
	restClient "github.com/go-provider-sdk/internal/clients/rest"
	"github.com/go-provider-sdk/internal/clients/rest/httptransport"
	"github.com/go-provider-sdk/internal/configmanager"
	"github.com/go-provider-sdk/internal/unmarshal"
	"github.com/go-provider-sdk/pkg/clientconfig"
	"github.com/go-provider-sdk/pkg/shared"
)

type DemoService struct {
	manager *configmanager.ConfigManager
}

func NewDemoService(manager *configmanager.ConfigManager) *DemoService {
	return &DemoService{
		manager: manager,
	}
}

func (api *DemoService) getConfig() *clientconfig.Config {
	return api.manager.GetDemo()
}

func (api *DemoService) SetBaseUrl(baseUrl string) {
	config := api.getConfig()
	config.SetBaseUrl(baseUrl)
}

func (api *DemoService) SetAccessToken(accessToken string) {
	config := api.getConfig()
	config.SetAccessToken(accessToken)
}

// Get the demo response
func (api *DemoService) GetDemo(params GetDemoRequestParams) (*shared.ClientResponse[Demo], error) {
	config := *api.getConfig()

	client := restClient.NewRestClient(config)

	request := httptransport.NewRequest("GET", "/demo", config)

	request.Options = params

	httpResponse, err := client.Call(request)
	if err != nil {
		return nil, err.GetError()
	}

	data, unmarshalError := unmarshal.ToObject[Demo](httpResponse)
	if unmarshalError != nil {
		return nil, unmarshalError
	}

	response := shared.ClientResponse[Demo]{
		Data: *data,
		Metadata: shared.ClientResponseMetadata{
			Headers:    httpResponse.Headers,
			StatusCode: httpResponse.StatusCode,
		},
	}

	return &response, nil
}

// Creates a new Demo access
func (api *DemoService) CreateDemo(demoParams DemoParams) (*shared.ClientResponse[Demo], error) {
	config := *api.getConfig()

	client := restClient.NewRestClient(config)

	request := httptransport.NewRequest("POST", "/demo", config)

	request.Body = demoParams

	httpResponse, err := client.Call(request)
	if err != nil {
		return nil, err.GetError()
	}

	data, unmarshalError := unmarshal.ToObject[Demo](httpResponse)
	if unmarshalError != nil {
		return nil, unmarshalError
	}

	response := shared.ClientResponse[Demo]{
		Data: *data,
		Metadata: shared.ClientResponseMetadata{
			Headers:    httpResponse.Headers,
			StatusCode: httpResponse.StatusCode,
		},
	}

	return &response, nil
}

// Deletes the Demo access
func (api *DemoService) DeleteDemo(demoParams DemoParams) error {
	config := *api.getConfig()

	client := restClient.NewRestClient(config)

	request := httptransport.NewRequest("DELETE", "/demo", config)

	request.Body = demoParams

	_, err := client.Call(request)
	if err != nil {
		return err.GetError()
	}

	return nil
}
