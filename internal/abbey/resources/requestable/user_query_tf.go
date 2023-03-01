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

const userQueryTypeAuthIdTf = "auth_id"

func UserQueryType() attr.Type {
	return types.ObjectType{AttrTypes: map[string]attr.Type{
		userQueryTypeAuthIdTf: AuthIdType(),
	}}
}

func UserQueryTfTypesType() tftypes.Type {
	return tftypes.Object{
		AttributeTypes: map[string]tftypes.Type{
			userQueryTypeAuthIdTf: AuthIdTfTypesType(),
		},
		OptionalAttributes: nil,
	}
}

func (q UserQuery) ToObjectValue(context.Context) (object basetypes.ObjectValue, diags Diagnostics) {
	var authIdValue attr.Value = types.StringNull()

	q.value.VisitUserQuery(UserQueryVisitor{
		AuthId: func(a AuthId) {
			authIdValue = types.StringValue(a.value)
		},
	})

	return types.ObjectValue(
		map[string]attr.Type{
			userQueryTypeAuthIdTf: types.StringType,
		},
		map[string]attr.Value{
			userQueryTypeAuthIdTf: authIdValue,
		},
	)
}

func (q UserQuery) Type(ctx context.Context) attr.Type {
	var authId AuthId

	return types.ObjectType{AttrTypes: map[string]attr.Type{
		userQueryTypeAuthIdTf: authId.Type(ctx),
	}}
}

func (q UserQuery) ToTerraformValue(ctx context.Context) (value tftypes.Value, err error) {
	var authIdValue tftypes.Value

	q.value.VisitUserQuery(UserQueryVisitor{
		AuthId: func(authId AuthId) {
			authIdValue, err = authId.ToTerraformValue(ctx)
		},
	})
	if err != nil {
		return value, err
	}

	return tftypes.NewValue(
		UserQueryTfTypesType(),
		map[string]tftypes.Value{
			userQueryTypeAuthIdTf: authIdValue,
		},
	), nil
}

func (q UserQuery) Equal(value attr.Value) bool {
	rhs, err := value.ToTerraformValue(context.Background())
	if err != nil {
		return false
	}

	lhs, err := q.ToTerraformValue(context.Background())
	if err != nil {
		return false
	}

	return lhs.Equal(rhs)
}

func (q UserQuery) IsNull() (defined bool) {
	return false
}

func (q UserQuery) IsUnknown() (defined bool) {
	return false
}

func (q UserQuery) String() string {
	var inner string

	q.value.VisitUserQuery(UserQueryVisitor{
		func(authId AuthId) {
			inner = authId.String()
		},
	})

	return fmt.Sprintf("UserQuery{%s}", inner)
}

func AuthIdType() attr.Type {
	return types.StringType
}

func AuthIdTfTypesType() tftypes.Type {
	return tftypes.String
}

func (a AuthId) Type(context.Context) attr.Type {
	return types.StringType
}

func (a AuthId) ToTerraformValue(context.Context) (value tftypes.Value, err error) {
	return tftypes.NewValue(tftypes.String, a.value), nil
}

func (a AuthId) Equal(value attr.Value) bool {
	rhs, err := value.ToTerraformValue(context.Background())
	if err != nil {
		return false
	}

	lhs, err := a.ToTerraformValue(context.Background())
	if err != nil {
		return false
	}

	return lhs.Equal(rhs)
}

func (a AuthId) IsNull() (defined bool) {
	return false
}

func (a AuthId) IsUnknown() (defined bool) {
	return false
}

func (a AuthId) String() string {
	return a.value
}
