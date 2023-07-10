// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type GenerateVariantType string

const (
	GenerateVariantTypeGenerate GenerateVariantType = "Generate"
)

func (e GenerateVariantType) ToPointer() *GenerateVariantType {
	return &e
}

func (e *GenerateVariantType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Generate":
		*e = GenerateVariantType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GenerateVariantType: %v", v)
	}
}

type GenerateVariant struct {
	Type  GenerateVariantType `json:"type"`
	Value Generate            `json:"value"`
}
