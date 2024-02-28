package httptransport

import "fmt"

type ErrorResponse struct {
	Err         error
	IsHttpError bool
	Response
}

func NewErrorResponse(err error, resp *Response) *ErrorResponse {
	if resp == nil {
		return &ErrorResponse{
			Err:         err,
			IsHttpError: false,
		}
	}

	return &ErrorResponse{
		Err:         err,
		IsHttpError: true,
		Response:    *resp,
	}
}

func (r *ErrorResponse) Clone() ErrorResponse {
	if r == nil {
		return ErrorResponse{}
	}

	clone := *r
	clone.Headers = make(map[string]string)
	for header, value := range r.Headers {
		clone.Headers[header] = value
	}
	return clone
}

func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%s", r.Err)
}

func (r *ErrorResponse) GetError() error {
	return r.Err
}
