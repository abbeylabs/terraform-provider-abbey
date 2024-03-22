package httptransport

import (
	"io"
	"net/http"
)

type Response[T any] struct {
	StatusCode int
	Headers    map[string]string
	Body       []byte
	Data       T
}

func (r *Response[T]) Clone() Response[T] {
	if r == nil {
		return Response[T]{
			Headers: make(map[string]string),
		}
	}

	clone := *r
	clone.Headers = make(map[string]string)
	for header, value := range r.Headers {
		clone.Headers[header] = value
	}
	return clone
}

func NewResponse[T any](resp *http.Response) (*Response[T], error) {
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, NewErrorResponse[T](err, nil)
	}

	responseHeaders := make(map[string]string)
	for key := range resp.Header {
		responseHeaders[key] = resp.Header.Get(key)
	}

	placeholderData := new(T)
	return &Response[T]{
		StatusCode: resp.StatusCode,
		Headers:    responseHeaders,
		Body:       body,
		Data:       *placeholderData,
	}, nil
}

func (r *Response[T]) GetStatusCode() int {
	return r.StatusCode
}

func (r *Response[T]) SetStatusCode(statusCode int) {
	r.StatusCode = statusCode
}

func (r *Response[T]) GetHeaders() map[string]string {
	return r.Headers
}

func (r *Response[T]) GetHeader(header string) string {
	return r.Headers[header]
}

func (r *Response[T]) SetHeader(header string, value string) {
	r.Headers[header] = value
}

func (r *Response[T]) GetBody() []byte {
	return r.Body
}

func (r *Response[T]) SetBody(body []byte) {
	r.Body = body
}
