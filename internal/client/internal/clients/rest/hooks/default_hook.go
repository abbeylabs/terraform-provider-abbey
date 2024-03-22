package hooks

type DefaultHook struct{}

func NewDefaultHook() Hook {
	return &DefaultHook{}
}

func (h *DefaultHook) BeforeRequest(req Request) Request {
	return req
}

func (h *DefaultHook) AfterResponse(req Request, resp Response) Response {
	return resp
}

func (h *DefaultHook) OnError(req Request, resp ErrorResponse) ErrorResponse {
	return resp
}
