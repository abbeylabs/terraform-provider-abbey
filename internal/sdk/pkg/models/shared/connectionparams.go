// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

type ConnectionParamsType string

const (
	ConnectionParamsTypeGithub    ConnectionParamsType = "Github"
	ConnectionParamsTypePagerduty ConnectionParamsType = "Pagerduty"
)

type ConnectionParams struct {
	ConnectionParamsGithubVariant *ConnectionParamsGithubVariant
	PagerdutyConnectionValue      *PagerdutyConnectionValue

	Type ConnectionParamsType
}

func CreateConnectionParamsGithub(github ConnectionParamsGithubVariant) ConnectionParams {
	typ := ConnectionParamsTypeGithub
	typStr := ConnectionType(typ)
	github.Type = typStr

	return ConnectionParams{
		ConnectionParamsGithubVariant: &github,
		Type:                          typ,
	}
}

func CreateConnectionParamsPagerduty(pagerduty PagerdutyConnectionValue) ConnectionParams {
	typ := ConnectionParamsTypePagerduty
	typStr := ConnectionType(typ)
	pagerduty.Type = typStr

	return ConnectionParams{
		PagerdutyConnectionValue: &pagerduty,
		Type:                     typ,
	}
}

func (u *ConnectionParams) UnmarshalJSON(data []byte) error {
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
		connectionParamsGithubVariant := new(ConnectionParamsGithubVariant)
		if err := d.Decode(&connectionParamsGithubVariant); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.ConnectionParamsGithubVariant = connectionParamsGithubVariant
		u.Type = ConnectionParamsTypeGithub
		return nil
	case "Pagerduty":
		d = json.NewDecoder(bytes.NewReader(data))
		pagerdutyConnectionValue := new(PagerdutyConnectionValue)
		if err := d.Decode(&pagerdutyConnectionValue); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.PagerdutyConnectionValue = pagerdutyConnectionValue
		u.Type = ConnectionParamsTypePagerduty
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u ConnectionParams) MarshalJSON() ([]byte, error) {
	if u.ConnectionParamsGithubVariant != nil {
		return json.Marshal(u.ConnectionParamsGithubVariant)
	}

	if u.PagerdutyConnectionValue != nil {
		return json.Marshal(u.PagerdutyConnectionValue)
	}

	return nil, nil
}
