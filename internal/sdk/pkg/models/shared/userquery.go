// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

type UserQueryType string

const (
	UserQueryTypeAuthID UserQueryType = "AuthId"
)

type UserQuery struct {
	UserQueryAuthIDVariant *UserQueryAuthIDVariant

	Type UserQueryType
}

func CreateUserQueryAuthID(authID UserQueryAuthIDVariant) UserQuery {
	typ := UserQueryTypeAuthID
	typStr := UserQueryAuthIDVariantType(typ)
	authID.Type = typStr

	return UserQuery{
		UserQueryAuthIDVariant: &authID,
		Type:                   typ,
	}
}

func (u *UserQuery) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	type discriminator struct {
		Type string
	}

	dis := new(discriminator)
	if err := json.Unmarshal(data, &dis); err != nil {
		return fmt.Errorf("could not unmarshal discriminator: %w", err)
	}

	switch dis.Type {
	case "AuthId":
		d = json.NewDecoder(bytes.NewReader(data))
		d.DisallowUnknownFields()
		userQueryAuthIDVariant := new(UserQueryAuthIDVariant)
		if err := d.Decode(&userQueryAuthIDVariant); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.UserQueryAuthIDVariant = userQueryAuthIDVariant
		u.Type = UserQueryTypeAuthID
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u UserQuery) MarshalJSON() ([]byte, error) {
	if u.UserQueryAuthIDVariant != nil {
		return json.Marshal(u.UserQueryAuthIDVariant)
	}

	return nil, nil
}
