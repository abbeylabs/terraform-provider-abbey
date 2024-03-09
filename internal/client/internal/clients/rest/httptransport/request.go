package httptransport

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"strings"

	"github.com/go-provider-sdk/internal/utils"
	"github.com/go-provider-sdk/pkg/clientconfig"
)

type paramMap struct {
	Key   string
	Value string
}

type Request struct {
	Context     context.Context
	Method      string
	Path        string
	Headers     map[string]string
	QueryParams map[string]string
	PathParams  map[string]string
	Options     any
	Body        any
	Config      clientconfig.Config
}

func NewRequest(ctx context.Context, method string, path string, config clientconfig.Config) Request {
	return Request{
		Context:     ctx,
		Method:      method,
		Path:        path,
		Headers:     make(map[string]string),
		QueryParams: make(map[string]string),
		PathParams:  make(map[string]string),
		Config:      config,
	}
}

func (r *Request) Clone() Request {
	if r == nil {
		return Request{
			Headers:     make(map[string]string),
			QueryParams: make(map[string]string),
			PathParams:  make(map[string]string),
		}
	}

	clone := *r
	clone.PathParams = utils.CloneMap(r.PathParams)
	clone.Headers = utils.CloneMap(r.Headers)
	clone.QueryParams = utils.CloneMap(r.QueryParams)

	return clone
}

func (r *Request) GetMethod() string {
	return r.Method
}

func (r *Request) SetMethod(method string) {
	r.Method = method
}

func (r *Request) GetBaseUrl() string {
	return *r.Config.BaseUrl
}

func (r *Request) SetBaseUrl(baseUrl string) {
	r.Config.SetBaseUrl(baseUrl)
}

func (r *Request) GetPath() string {
	return r.Path
}

func (r *Request) SetPath(path string) {
	r.Path = path
}

func (r *Request) GetHeader(header string) string {
	return r.Headers[header]
}

func (r *Request) SetHeader(header string, value string) {
	r.Headers[header] = value
}

func (r *Request) GetPathParam(param string) string {
	return r.PathParams[param]
}

func (r *Request) SetPathParam(param string, value any) {
	r.PathParams[param] = fmt.Sprintf("%v", value)
}

func (r *Request) GetQueryParam(header string) string {
	return r.QueryParams[header]
}

func (r *Request) SetQueryParam(header string, value string) {
	r.QueryParams[header] = value
}

func (r *Request) GetOptions() any {
	return r.Options
}

func (r *Request) SetOptions(options any) {
	r.Options = options
}

func (r *Request) GetBody() any {
	return r.Body
}

func (r *Request) SetBody(body any) {
	r.Body = body
}

func (r *Request) GetContext() context.Context {
	return r.Context
}

func (r *Request) SetContext(ctx context.Context) {
	r.Context = ctx
}

func (r *Request) CreateHttpRequest() (*http.Request, error) {
	requestUrl := r.getRequestUrl()

	requestBody, err := r.bodyToBytesReader()
	if err != nil {
		return nil, err
	}

	var httpRequest *http.Request
	if requestBody == nil {
		httpRequest, err = http.NewRequestWithContext(r.Context, r.Method, requestUrl, nil)
	} else {
		httpRequest, err = http.NewRequestWithContext(r.Context, r.Method, requestUrl, requestBody)
	}

	httpRequest.Header = r.getRequestHeaders()

	return httpRequest, err
}

func (r *Request) getRequestUrl() string {
	requestPath := r.Path
	for paramName, paramValue := range r.PathParams {
		placeholder := "{" + paramName + "}"
		requestPath = strings.ReplaceAll(requestPath, placeholder, url.PathEscape(paramValue))
	}

	requestOptions := ""
	params := r.getRequestQueryParams()
	if len(params) > 0 {
		requestOptions = fmt.Sprintf("?%s", params.Encode())
	}

	return *r.Config.BaseUrl + requestPath + requestOptions
}

func (r *Request) bodyToBytesReader() (*bytes.Reader, error) {
	if r.Body == nil {
		return nil, nil
	}

	marshalledBody, err := json.Marshal(r.Body)
	if err != nil {
		return nil, err
	}
	reqBody := bytes.NewReader(marshalledBody)

	return reqBody, nil
}

func (r *Request) getRequestQueryParams() url.Values {
	params := url.Values{}
	for key, value := range r.QueryParams {
		params.Add(key, value)
	}

	for _, p := range tagsToMap("queryParam", r.Options) {
		params.Add(p.Key, p.Value)
	}

	return params
}

func (r *Request) getRequestHeaders() http.Header {
	headers := http.Header{}
	for key, value := range r.Headers {
		headers.Add(key, value)
	}

	for _, p := range tagsToMap("headerParam", r.Options) {
		headers.Add(p.Key, p.Value)
	}

	return headers
}

func tagsToMap(tag string, obj any) []paramMap {
	tagMap := make([]paramMap, 0)

	if obj == nil {
		return tagMap
	}

	values := utils.GetReflectValue(reflect.ValueOf(obj))
	for i := 0; i < values.NumField(); i++ {
		key, found := values.Type().Field(i).Tag.Lookup(tag)
		if !found || values.Field(i).Type().Kind() == reflect.Pointer && values.Field(i).IsNil() {
			continue
		}

		field := utils.GetReflectValue(values.Field(i))

		fieldKind := utils.GetReflectKind(field.Type())
		if fieldKind == reflect.Array || fieldKind == reflect.Slice {
			for j := 0; j < field.Len(); j++ {
				p := paramMap{Key: key, Value: fmt.Sprint(field.Index(j))}
				tagMap = append(tagMap, p)
			}
		} else {
			p := paramMap{Key: key, Value: fmt.Sprint(field)}
			tagMap = append(tagMap, p)
		}
	}

	return tagMap
}
