package demo

import (
	"context"
	restClient "github.com/go-provider-sdk/internal/clients/rest"
	"github.com/go-provider-sdk/internal/clients/rest/httptransport"
	"github.com/go-provider-sdk/internal/configmanager"
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
func (api *DemoService) GetDemo(ctx context.Context, params GetDemoRequestParams) (*shared.ClientResponse[Demo], *shared.ClientError) {
	config := *api.getConfig()

	client := restClient.NewRestClient[Demo](config)

	request := httptransport.NewRequest(ctx, "GET", "/demo", config)

	request.Options = params

	resp, err := client.Call(request)
	if err != nil {
		return nil, shared.NewClientError[Demo](err)
	}

	return shared.NewClientResponse[Demo](resp), nil
}

// Creates a new Demo access
func (api *DemoService) CreateDemo(ctx context.Context, demoParams DemoParams) (*shared.ClientResponse[Demo], *shared.ClientError) {
	config := *api.getConfig()

	client := restClient.NewRestClient[Demo](config)

	request := httptransport.NewRequest(ctx, "POST", "/demo", config)

	request.Body = demoParams

	resp, err := client.Call(request)
	if err != nil {
		return nil, shared.NewClientError[Demo](err)
	}

	return shared.NewClientResponse[Demo](resp), nil
}

// Deletes the Demo access
func (api *DemoService) DeleteDemo(ctx context.Context, demoParams DemoParams) (*shared.ClientResponse[any], *shared.ClientError) {
	config := *api.getConfig()

	client := restClient.NewRestClient[any](config)

	request := httptransport.NewRequest(ctx, "DELETE", "/demo", config)

	request.Body = demoParams

	resp, err := client.Call(request)
	if err != nil {
		return nil, shared.NewClientError[any](err)
	}

	return shared.NewClientResponse[any](resp), nil
}
