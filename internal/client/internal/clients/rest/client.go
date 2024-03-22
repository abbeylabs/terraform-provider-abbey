package rest

import (
	"github.com/go-provider-sdk/internal/clients/rest/handlers"
	"github.com/go-provider-sdk/internal/clients/rest/hooks"
	"github.com/go-provider-sdk/internal/clients/rest/httptransport"
	"github.com/go-provider-sdk/pkg/clientconfig"
)

type RestClient[T any] struct {
	handlers *handlers.HandlerChain[T]
}

func NewRestClient[T any](config clientconfig.Config) *RestClient[T] {
	defaultHeadersHandler := handlers.NewDefaultHeadersHandler[T]()
	retryHandler := handlers.NewRetryHandler[T]()
	bearerTokenHandler := handlers.NewAccessTokenHandler[T]()
	unmarshalHandler := handlers.NewUnmarshalHandler[T]()
	hookHandler := handlers.NewHookHandler[T](hooks.NewDefaultHook())
	terminatingHandler := handlers.NewTerminatingHandler[T]()

	handlers := handlers.BuildHandlerChain[T]().
		AddHandler(defaultHeadersHandler).
		AddHandler(retryHandler).
		AddHandler(bearerTokenHandler).
		AddHandler(unmarshalHandler).
		AddHandler(hookHandler).
		AddHandler(terminatingHandler)

	return &RestClient[T]{
		handlers: handlers,
	}
}

func (client *RestClient[T]) Call(request httptransport.Request) (*httptransport.Response[T], *httptransport.ErrorResponse[T]) {
	return client.handlers.CallApi(request)
}
