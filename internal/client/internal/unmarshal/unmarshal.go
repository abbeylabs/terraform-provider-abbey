package unmarshal

import (
	"fmt"
	"reflect"

	"github.com/go-provider-sdk/internal/utils"
)

func Unmarshal(source []byte, target any) error {

	targetValue := reflect.ValueOf(target)
	if targetValue.Kind() != reflect.Ptr || targetValue.IsNil() {
		return fmt.Errorf("target must be a non-nil pointer")
	}

	if isComplexObject(target) || isObject(target) || isArray(target) {
		return ToObject(source, target)
	} else if isString(targetValue.Elem().Kind()) {
		return ToString(source, targetValue)
	} else if isInteger(targetValue.Elem().Kind()) {
		return ToInt(source, targetValue)
	} else if isFloat(targetValue.Elem().Kind()) {
		return ToFloat(source, targetValue)
	} else if isBool(targetValue.Elem().Kind()) {
		return ToBool(source, targetValue)
	}

	return nil
}

func isArray(target any) bool {
	targetType := reflect.TypeOf(target)
	kind := utils.GetReflectKind(targetType)
	return kind == reflect.Array || kind == reflect.Slice
}

func isObject(target any) bool {
	targetType := reflect.TypeOf(target)
	return utils.GetReflectKind(targetType) == reflect.Struct
}

func isComplexObject(target any) bool {
	targetType := reflect.TypeOf(target)
	if utils.GetReflectKind(targetType) != reflect.Struct {
		return false
	}

	allFieldsAreOneOf := true

	structValue := utils.GetReflectValue(reflect.ValueOf(target))
	for i := 0; i < structValue.NumField(); i++ {
		field := structValue.Type().Field(i)
		allFieldsAreOneOf = isOneOfField(field) && allFieldsAreOneOf
	}

	return allFieldsAreOneOf
}

func isOneOfField(field reflect.StructField) bool {
	_, found := field.Tag.Lookup("oneof")
	return found
}
