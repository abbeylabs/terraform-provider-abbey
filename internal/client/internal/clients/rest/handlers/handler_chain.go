package handlers

import (
	"github.com/go-provider-sdk/internal/clients/rest/httptransport"
)

type Handler[T any] interface {
	Handle(req httptransport.Request) (*httptransport.Response[T], *httptransport.ErrorResponse[T])
	SetNext(handler Handler[T])
}

type HandlerChain[T any] struct {
	head Handler[T]
	tail Handler[T]
}

func BuildHandlerChain[T any]() *HandlerChain[T] {
	return &HandlerChain[T]{}
}

func (chain *HandlerChain[T]) AddHandler(handler Handler[T]) *HandlerChain[T] {
	if chain.head == nil {
		chain.head = handler
		chain.tail = handler
		return chain
	}

	chain.tail.SetNext(handler)
	chain.tail = handler

	return chain
}

func (chain *HandlerChain[T]) CallApi(request httptransport.Request) (*httptransport.Response[T], *httptransport.ErrorResponse[T]) {
	return chain.head.Handle(request)
}
