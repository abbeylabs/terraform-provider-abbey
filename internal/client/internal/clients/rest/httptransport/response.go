package httptransport

type Response struct {
	StatusCode int
	Headers    map[string]string
	Body       []byte
}

func (r *Response) Clone() Response {
	if r == nil {
		return Response{
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

func (r *Response) GetStatusCode() int {
	return r.StatusCode
}

func (r *Response) SetStatusCode(statusCode int) {
	r.StatusCode = statusCode
}

func (r *Response) GetHeader(header string) string {
	return r.Headers[header]
}

func (r *Response) SetHeader(header string, value string) {
	r.Headers[header] = value
}

func (r *Response) GetBody() []byte {
	return r.Body
}

func (r *Response) SetBody(body []byte) {
	r.Body = body
}
