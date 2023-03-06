package entity

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/moznion/go-optional"
)

type (
	PolicySet struct {
		GrantIf  Option[[]Policy] `json:"grant_if"`
		RevokeIf Option[[]Policy] `json:"revoke_if"`
	}

	Policy struct {
		Bundle Option[string] `json:"bundle"`
		Query  Option[string] `json:"query"`
	}
)

func PolicyFromObject(ctx context.Context, object types.Object) (invalid Policy, err error) {
	var (
		attrs  = object.Attributes()
		bundle Option[string]
		query  Option[string]
	)

	bundleValue, err := attrs["bundle"].ToTerraformValue(ctx)
	if err != nil {
		return invalid, err
	}

	queryValue, err := attrs["query"].ToTerraformValue(ctx)
	if err != nil {
		return invalid, err
	}

	if !bundleValue.IsNull() {
		var s string

		if err := bundleValue.As(&s); err != nil {
			return invalid, err
		}

		bundle = Some(s)
	}

	if !queryValue.IsNull() {
		var s string

		if err := queryValue.As(&s); err != nil {
			return invalid, err
		}

		query = Some(s)
	}

	return Policy{
		Bundle: bundle,
		Query:  query,
	}, nil
}
