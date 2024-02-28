package validation

import (
	"fmt"
	"reflect"
)

func validateRequired(fieldValue reflect.Value, fieldType reflect.StructField) error {
	if IsRequiredField(fieldType) && fieldValue.IsNil() {
		return fmt.Errorf("field %s is required", fieldType.Name)
	}

	return nil
}

func IsRequiredField(fieldType reflect.StructField) bool {
	required, found := fieldType.Tag.Lookup("required")
	return found && required == "true"
}

func IsOptionalField(fieldType reflect.StructField) bool {
	required, found := fieldType.Tag.Lookup("required")
	return !found || required == "" || required == "false"
}
