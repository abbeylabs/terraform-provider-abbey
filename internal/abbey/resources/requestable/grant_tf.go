package requestable

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	. "github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"

	"abbey.so/terraform-provider-abbey/internal/abbey/value"
)

type GrantTf struct {
	Grant

	state value.State
}

func NewGrant(g Grant) GrantTf {
	return GrantTf{
		Grant: g,
		state: value.NewValidState(),
	}
}

func NewNullGrant() GrantTf {
	var g Grant

	return GrantTf{
		Grant: g,
		state: value.NewNullState(),
	}
}

func NewUnknownGrant() GrantTf {
	var g Grant

	return GrantTf{
		Grant: g,
		state: value.NewUnknownState(),
	}
}

func (g GrantTf) ToObjectValue(ctx context.Context) (object basetypes.ObjectValue, diags Diagnostics) {
	g.state.Visit(value.StateVisitor{
		Null: func() {
			object = types.ObjectNull(g.AttrTypes(ctx))
		},
		Unknown: func() {
			object = types.ObjectUnknown(g.AttrTypes(ctx))
		},
		Valid: func() {
			var (
				generate      GenerateGrant
				generateValue attr.Value = types.ObjectNull(generate.AttrTypes(ctx))
				diags_        Diagnostics
			)

			g.Value.VisitGrant(GrantVisitor{
				Generate: func(grant GenerateGrant) {
					generateValue, diags = grant.ToObjectValue(ctx)
				},
			})
			if diags.HasError() {
				return
			}

			object, diags_ = types.ObjectValue(
				map[string]attr.Type{
					grantTypeGenerateTf: generate.Type(ctx),
				},
				map[string]attr.Value{
					grantTypeGenerateTf: generateValue,
				},
			)
			diags.Append(diags_...)
		},
	})

	return object, diags
}

func (g GrantTf) AttrTypes(ctx context.Context) map[string]attr.Type {
	var generate GenerateGrant

	return map[string]attr.Type{
		grantTypeGenerateTf: generate.Type(ctx),
	}
}

func GrantTfTfTypesType() tftypes.Type {
	return tftypes.Object{
		AttributeTypes: map[string]tftypes.Type{
			grantTypeGenerateTf: GenerateGrantTfTypesType(),
		},
		OptionalAttributes: nil,
	}
}

func (g GrantTf) Type(context.Context) attr.Type {
	return GrantType{}
}

func (g GrantTf) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	var (
		val tftypes.Value
		err error
	)

	g.state.Visit(value.StateVisitor{
		Null: func() {
			val = tftypes.NewValue(GrantTfTfTypesType(), nil)
		},
		Unknown: func() {
			val = tftypes.NewValue(GrantTfTfTypesType(), tftypes.UnknownValue)
		},
		Valid: func() {
			var generateValue tftypes.Value

			g.Value.VisitGrant(GrantVisitor{
				Generate: func(grant GenerateGrant) {
					generateValue, err = grant.ToTerraformValue(ctx)
				},
			})
			if err != nil {
				return
			}

			val = tftypes.NewValue(
				GrantTfTfTypesType(),
				map[string]tftypes.Value{
					grantTypeGenerateTf: generateValue,
				},
			)
		},
	})
	if err != nil {
		return val, err
	}

	return val, nil
}

func (g GrantTf) Equal(value attr.Value) bool {
	rhs, err := value.ToTerraformValue(context.Background())
	if err != nil {
		return false
	}

	lhs, err := g.ToTerraformValue(context.Background())
	if err != nil {
		return false
	}

	return lhs.Equal(rhs)
}

func (g GrantTf) IsNull() (isNull bool) {
	g.state.Visit(value.StateVisitor{
		Null: func() {
			isNull = true
		},
		Unknown: func() {},
		Valid:   func() {},
	})

	return isNull
}

func (g GrantTf) IsUnknown() (unknown bool) {
	g.state.Visit(value.StateVisitor{
		Null: func() {},
		Unknown: func() {
			unknown = true
		},
		Valid: func() {},
	})

	return unknown
}

func (g GrantTf) String() string {
	var inner string

	g.Value.VisitGrant(GrantVisitor{
		Generate: func(grant GenerateGrant) {
			inner = grant.String()
		},
	})

	return fmt.Sprintf("Grant{%s}", inner)
}

func GenerateGrantTfTypesType() tftypes.Type {
	return tftypes.Object{
		AttributeTypes: map[string]tftypes.Type{
			generateGrantTypeGithubTf: GithubGenerateDestinationTfTypesType(),
		},
		OptionalAttributes: nil,
	}
}

func (g GenerateGrant) AttrTypes(ctx context.Context) map[string]attr.Type {
	var github GithubGenerateDestination

	return map[string]attr.Type{
		generateGrantTypeGithubTf: github.Type(ctx),
	}
}

func (g GenerateGrant) Type(ctx context.Context) attr.Type {
	return types.ObjectType{AttrTypes: g.AttrTypes(ctx)}
}

func (g GenerateGrant) ToObjectValue(ctx context.Context) (object basetypes.ObjectValue, diags Diagnostics) {
	var (
		github      GithubGenerateDestination
		githubValue attr.Value = types.ObjectNull(github.AttrTypes())
	)

	g.Value.VisitGenerateGrant(GenerateGrantVisitor{
		Github: func(dest GithubGenerateDestination) {
			githubValue, diags = dest.ToObjectValue(ctx)
		},
	})
	if diags.HasError() {
		return object, diags
	}

	return types.ObjectValue(
		map[string]attr.Type{
			generateGrantTypeGithubTf: github.Type(ctx),
		},
		map[string]attr.Value{
			generateGrantTypeGithubTf: githubValue,
		},
	)
}

func (g GenerateGrant) ToTerraformValue(ctx context.Context) (value tftypes.Value, err error) {
	var githubValue tftypes.Value

	g.Value.VisitGenerateGrant(GenerateGrantVisitor{
		Github: func(github GithubGenerateDestination) {
			githubValue, err = github.ToTerraformValue(ctx)
		},
	})
	if err != nil {
		return value, err
	}

	return tftypes.NewValue(
		GenerateGrantTfTypesType(),
		map[string]tftypes.Value{
			generateGrantTypeGithubTf: githubValue,
		},
	), nil
}

func (g GenerateGrant) Equal(value attr.Value) bool {
	rhs, err := value.ToTerraformValue(context.Background())
	if err != nil {
		return false
	}

	lhs, err := g.ToTerraformValue(context.Background())
	if err != nil {
		return false
	}

	return lhs.Equal(rhs)
}

func (g GenerateGrant) IsNull() bool {
	defined := false

	g.Value.VisitGenerateGrant(GenerateGrantVisitor{
		Github: func(GithubGenerateDestination) {
			defined = true
		},
	})

	return !defined
}

func (g GenerateGrant) IsUnknown() bool {
	defined := false

	g.Value.VisitGenerateGrant(GenerateGrantVisitor{
		Github: func(GithubGenerateDestination) {
			defined = true
		},
	})

	return !defined
}

func (g GenerateGrant) String() string {
	var inner string

	g.Value.VisitGenerateGrant(GenerateGrantVisitor{
		Github: func(github GithubGenerateDestination) {
			inner = github.String()
		},
	})

	return fmt.Sprintf("GenerateGrant{%s}", inner)
}

func GithubGenerateDestinationTfTypesType() tftypes.Type {
	return tftypes.Object{
		AttributeTypes: map[string]tftypes.Type{
			"repo":   tftypes.String,
			"path":   tftypes.String,
			"append": tftypes.String,
		},
		OptionalAttributes: nil,
	}
}

func (g GithubGenerateDestination) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"repo":   types.StringType,
		"path":   types.StringType,
		"append": types.StringType,
	}
}

func (g GithubGenerateDestination) Type(context.Context) attr.Type {
	return types.ObjectType{AttrTypes: g.AttrTypes()}
}

func (g GithubGenerateDestination) ToObjectValue(context.Context) (object basetypes.ObjectValue, diags Diagnostics) {
	return types.ObjectValue(
		g.AttrTypes(),
		map[string]attr.Value{
			"repo":   basetypes.NewStringValue(g.Repo),
			"path":   basetypes.NewStringValue(g.Path),
			"append": basetypes.NewStringValue(g.Append),
		},
	)
}

func (g GithubGenerateDestination) ToTerraformValue(context.Context) (value tftypes.Value, err error) {
	return tftypes.NewValue(
		tftypes.Object{
			AttributeTypes: map[string]tftypes.Type{
				"repo":   tftypes.String,
				"path":   tftypes.String,
				"append": tftypes.String,
			},
			OptionalAttributes: nil,
		},
		map[string]tftypes.Value{
			"repo":   tftypes.NewValue(tftypes.String, g.Repo),
			"path":   tftypes.NewValue(tftypes.String, g.Path),
			"append": tftypes.NewValue(tftypes.String, g.Append),
		},
	), nil
}

func (g GithubGenerateDestination) Equal(value attr.Value) bool {
	rhs, err := value.ToTerraformValue(context.Background())
	if err != nil {
		return false
	}

	lhs, err := g.ToTerraformValue(context.Background())
	if err != nil {
		return false
	}

	return lhs.Equal(rhs)
}

func (g GithubGenerateDestination) IsNull() bool {
	return false
}

func (g GithubGenerateDestination) IsUnknown() bool {
	return false
}

func (g GithubGenerateDestination) String() string {
	return fmt.Sprintf("GenerateGrant{Repo: %q, Path: %q, Append: %q}", g.Repo, g.Path, g.Append)
}
