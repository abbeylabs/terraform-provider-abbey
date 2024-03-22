package unmarshal

import (
	"reflect"
	"strconv"
)

func ToString(source []byte, target reflect.Value) error {
	target.Elem().SetString(string(source))
	return nil
}

func ToInt(source []byte, target reflect.Value) error {
	intBody, err := strconv.ParseInt(string(source), 10, 64)
	if err != nil {
		return err
	}

	target.Elem().SetInt(intBody)

	return nil
}

func ToFloat(source []byte, target reflect.Value) error {
	floatBody, err := strconv.ParseFloat(string(source), 64)
	if err != nil {
		return err
	}

	target.Elem().SetFloat(floatBody)

	return nil
}

func ToBool(source []byte, target reflect.Value) error {
	boolBody, err := strconv.ParseBool(string(source))
	if err != nil {
		return err
	}

	target.Elem().SetBool(boolBody)

	return nil
}
