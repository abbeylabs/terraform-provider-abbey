package requestable

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	. "github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

type GrantType struct{}

func (t GrantType) TerraformType(context.Context) tftypes.Type {
	return tftypes.Object{
		AttributeTypes: map[string]tftypes.Type{
			grantTypeGenerateTf: GenerateGrantTfTypesType(),
		},
		OptionalAttributes: nil,
	}
}

func (t GrantType) ValueFromTerraform(ctx context.Context, value tftypes.Value) (value_ attr.Value, err error) {
	var g Grant

	if !value.IsFullyKnown() {
		return types.ObjectNull(g.AttrTypes(ctx)), nil
	}

	var m map[string]tftypes.Value
	if err := value.As(&m); err != nil {
		return nil, err
	}

	if len(m) == 0 {
		return types.ObjectNull(g.AttrTypes(ctx)), nil
	}

	var inner GrantEnum

	for key, val := range m {
		switch key {
		case grantTypeGenerateTf:
			inner_, err := GenerateGrantFromTfTypesValue(val)
			if err != nil {
				return value_, err
			}
			if inner_ == nil {
				continue
			}

			inner = inner_
		default:
			return value_, fmt.Errorf("unknown key: %s", key)
		}
	}

	return &Grant{value: inner}, nil
}

func (t GrantType) ValueType(context.Context) attr.Value {
	var w Grant
	return &w
}

func (t GrantType) Equal(ty attr.Type) bool {
	_, ok := ty.(GrantType)
	return ok
}

func (GrantType) String() string {
	return "Grant"
}

func (w GrantType) ApplyTerraform5AttributePathStep(step tftypes.AttributePathStep) (any, error) {
	attrName, ok := step.(tftypes.AttributeName)
	if !ok {
		return nil, fmt.Errorf("cannot apply step %T to GrantType", step)
	}

	switch string(attrName) {
	case grantTypeGenerateTf:
		return BuiltinWorkflowTfTypesType(), nil
	default:
		return nil, fmt.Errorf("undefined attribute name %s in GrantType", attrName)
	}
}

func (t GrantType) ValueFromObject(
	ctx context.Context,
	value basetypes.ObjectValue,
) (basetypes.ObjectValuable, Diagnostics) {
	var g Grant

	diags := value.As(ctx, &g, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    false,
		UnhandledUnknownAsEmpty: false,
	})
	if diags.HasError() {
		return nil, diags
	}

	return &g, diags
}
