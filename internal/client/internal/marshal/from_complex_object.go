package marshal

import (
	"encoding/json"
	"errors"
	"reflect"

	"github.com/go-provider-sdk/internal/utils"
)

func FromComplexObject(obj any) ([]byte, error) {
	types := utils.GetReflectType(reflect.TypeOf(obj))
	values := utils.GetReflectValue(reflect.ValueOf(obj))

	for i := 0; i < types.NumField(); i++ {
		if !values.Field(i).IsNil() {
			return json.Marshal(values.Field(i).Interface())
		}
	}

	return nil, errors.New("cannot marshal complex object, no non-nil fields found")
}
