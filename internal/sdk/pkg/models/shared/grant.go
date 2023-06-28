// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

type GrantType string

const (
	GrantTypeGenerate GrantType = "Generate"
)

type Grant struct {
	GenerateVariant *GenerateVariant

	Type GrantType
}

func CreateGrantGenerate(generate GenerateVariant) Grant {
	typ := GrantTypeGenerate
	typStr := GenerateVariantType(typ)
	generate.Type = typStr

	return Grant{
		GenerateVariant: &generate,
		Type:            typ,
	}
}

func (u *Grant) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	type discriminator struct {
		Type string
	}

	dis := new(discriminator)
	if err := json.Unmarshal(data, &dis); err != nil {
		return fmt.Errorf("could not unmarshal discriminator: %w", err)
	}

	switch dis.Type {
	case "Generate":
		d = json.NewDecoder(bytes.NewReader(data))
		d.DisallowUnknownFields()
		generateVariant := new(GenerateVariant)
		if err := d.Decode(&generateVariant); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.GenerateVariant = generateVariant
		u.Type = GrantTypeGenerate
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u Grant) MarshalJSON() ([]byte, error) {
	if u.GenerateVariant != nil {
		return json.Marshal(u.GenerateVariant)
	}

	return nil, nil
}