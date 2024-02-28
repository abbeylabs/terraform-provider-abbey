package validation

import (
	"reflect"

	"github.com/go-provider-sdk/internal/utils"
)

type validatorFunc = func(fieldValue reflect.Value, fieldType reflect.StructField) error

func ValidateData(data any) error {
	if data == nil {
		return nil
	}

	dataType := reflect.TypeOf(data)
	dataValue := reflect.ValueOf(data)
	if utils.GetReflectKind(dataType) == reflect.Struct {
		return validateStruct(data)
	} else if dataType.Kind() == reflect.Array || dataType.Kind() == reflect.Slice {
		return validateArray(dataValue)
	}

	return nil
}

func validateStruct(data any) error {
	structValue := utils.GetReflectValue(reflect.ValueOf(data))
	for i := 0; i < structValue.NumField(); i++ {
		fieldValue := structValue.Field(i)
		fieldType := structValue.Type().Field(i)

		err := validateField(fieldValue, fieldType)
		if err != nil {
			return err
		}

		if fieldValue.IsNil() {
			continue
		}

		kind := utils.GetReflectKind(fieldType.Type)
		if kind == reflect.Struct || kind == reflect.Array || kind == reflect.Slice {
			err := ValidateData(fieldValue.Interface())
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func validateArray(value reflect.Value) error {
	arrayValue := utils.GetReflectValue(value)
	for j := 0; j < arrayValue.Len(); j++ {
		err := ValidateData(arrayValue.Index(j).Interface())
		if err != nil {
			return err
		}
	}

	return nil
}

func validateField(fieldValue reflect.Value, fieldType reflect.StructField) error {
	validators := getValidators(fieldType)
	for _, validator := range validators {
		err := validator(fieldValue, fieldType)
		if err != nil {
			return err
		}
	}

	return nil
}

func getValidators(fieldType reflect.StructField) []validatorFunc {
	return []validatorFunc{validateRequired}
}
