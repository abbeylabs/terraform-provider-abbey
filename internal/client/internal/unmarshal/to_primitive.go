package unmarshal

import (
	"strconv"

	"github.com/go-provider-sdk/internal/clients/rest/httptransport"
)

func ToString(r *httptransport.Response) (*string, error) {
	result := string(r.Body)
	return &result, nil
}

func ToInt(r *httptransport.Response) (*int64, error) {
	result, err := strconv.ParseInt(string(r.Body), 10, 64)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func ToFloat(r *httptransport.Response) (*float64, error) {
	result, err := strconv.ParseFloat(string(r.Body), 64)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func ToBool(r *httptransport.Response) (*bool, error) {
	result, err := strconv.ParseBool(string(r.Body))
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func ToEnum[T ~string](r *httptransport.Response) (*T, error) {
	result := T(r.Body)
	return &result, nil
}
