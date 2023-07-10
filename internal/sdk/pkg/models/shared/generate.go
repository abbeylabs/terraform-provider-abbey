// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

type GenerateType string

const (
	GenerateTypeGithub GenerateType = "Github"
)

type Generate struct {
	GenerateGithubVariant *GenerateGithubVariant

	Type GenerateType
}

func CreateGenerateGithub(github GenerateGithubVariant) Generate {
	typ := GenerateTypeGithub
	typStr := GenerateGithubVariantType(typ)
	github.Type = typStr

	return Generate{
		GenerateGithubVariant: &github,
		Type:                  typ,
	}
}

func (u *Generate) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	type discriminator struct {
		Type string
	}

	dis := new(discriminator)
	if err := json.Unmarshal(data, &dis); err != nil {
		return fmt.Errorf("could not unmarshal discriminator: %w", err)
	}

	switch dis.Type {
	case "Github":
		d = json.NewDecoder(bytes.NewReader(data))
		generateGithubVariant := new(GenerateGithubVariant)
		if err := d.Decode(&generateGithubVariant); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.GenerateGithubVariant = generateGithubVariant
		u.Type = GenerateTypeGithub
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u Generate) MarshalJSON() ([]byte, error) {
	if u.GenerateGithubVariant != nil {
		return json.Marshal(u.GenerateGithubVariant)
	}

	return nil, nil
}
