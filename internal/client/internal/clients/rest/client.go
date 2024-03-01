package rest

import (
	"github.com/go-provider-sdk/internal/clients/rest/handlers"
	"github.com/go-provider-sdk/internal/clients/rest/hooks"
	"github.com/go-provider-sdk/internal/clients/rest/httptransport"
	"github.com/go-provider-sdk/pkg/clientconfig"
)

type RestClient struct {
	handlers *handlers.HandlerChain
}

func NewRestClient(config clientconfig.Config) *RestClient {
	defaultHeadersHandler := handlers.NewDefaultHeadersHandler()
	retryHandler := handlers.NewRetryHandler()
	bearerTokenHandler := handlers.NewAccessTokenHandler()
	hookHandler := handlers.NewHookHandler(hooks.NewDefaultHook())
	requestValidationHandler := handlers.NewRequestValidationHandler()
	terminatingHandler := handlers.NewTerminatingHandler()

	handlers := handlers.BuildHandlerChain().
		AddHandler(defaultHeadersHandler).
		AddHandler(retryHandler).
		AddHandler(bearerTokenHandler).
		AddHandler(hookHandler).
		AddHandler(requestValidationHandler).
		AddHandler(terminatingHandler)

	return &RestClient{
		handlers: handlers,
	}
}

func (client *RestClient) Call(request httptransport.Request) (*httptransport.Response, *httptransport.ErrorResponse) {
	return client.handlers.CallApi(request)
}
