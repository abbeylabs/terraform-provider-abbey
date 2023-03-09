package grantkit

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"abbey.so/terraform-provider-abbey/internal/abbey/resources/requestable"
)

type Output struct {
	Location  types.String `tfsdk:"location"`
	Append    types.String `tfsdk:"append"`
	Overwrite types.String `tfsdk:"overwrite"`
}

func (self Output) ToObject() (types.Object, diag.Diagnostics) {
	return types.ObjectValue(
		OutputAttrTypes(),
		map[string]attr.Value{
			"location":  self.Location,
			"append":    self.Append,
			"overwrite": self.Overwrite,
		},
	)
}

func OutputFromRequestableGithubDestination(github requestable.GithubGenerateDestination) Output {
	return Output{
		Location:  types.StringValue(fmt.Sprintf("github://%s%s", github.Repo, github.Path)),
		Append:    types.StringValue(github.Append),
		Overwrite: types.StringNull(),
	}
}

func OutputAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"location":  types.StringType,
		"append":    types.StringType,
		"overwrite": types.StringType,
	}
}

func RequestableGrantFromOutputObject(
	ctx context.Context,
	object types.Object,
) (*requestable.Grant, diag.Diagnostics) {
	var (
		diags  diag.Diagnostics
		output Output
	)

	diags.Append(object.As(ctx, &output, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    false,
		UnhandledUnknownAsEmpty: false,
	})...)
	if diags.HasError() {
		return nil, diags
	}

	s, ok := strings.CutPrefix(output.Location.ValueString(), "github://")
	if !ok {
		diags.AddError("Invalid Input", "Only `github://` scheme is supported.")
		return nil, diags
	}

	segments := strings.SplitN(s, "/", 3)
	if len(segments) != 3 {
		diags.AddError("Invalid Input", "Failed to parse GitHub location.")
		return nil, diags
	}

	owner := segments[0]
	repo := segments[1]
	path := segments[2]

	return &requestable.Grant{
		Value: requestable.GenerateGrant{
			Value: requestable.GithubGenerateDestination{
				Repo:   strings.Join([]string{owner, repo}, "/"),
				Path:   path,
				Append: output.Append.ValueString(),
			},
		},
	}, diags
}
