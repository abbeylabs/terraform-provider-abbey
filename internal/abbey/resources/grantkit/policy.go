package grantkit

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	. "github.com/moznion/go-optional"

	"abbey.so/terraform-provider-abbey/internal/abbey/entity"
)

type (
	Policy struct {
		Bundle types.String `tfsdk:"bundle"`
		Query  types.String `tfsdk:"query"`
	}

	PolicySet struct {
		GrantIf  types.List `tfsdk:"grant_if"`
		RevokeIf types.List `tfsdk:"revoke_if"`
	}
)

func PolicyAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"bundle": types.StringType,
		"query":  types.StringType,
	}
}

func PolicyType() attr.Type {
	return types.ObjectType{AttrTypes: PolicyAttrTypes()}
}

func (self Policy) ToObject() (types.Object, diag.Diagnostics) {
	return types.ObjectValue(PolicyAttrTypes(), map[string]attr.Value{
		"bundle": self.Bundle,
		"query":  self.Query,
	})
}

func PolicyFromView(policy entity.Policy) Policy {
	bundle := types.StringNull()
	query := types.StringNull()

	policy.Bundle.IfSome(func(v string) {
		bundle = types.StringValue(v)
	})
	policy.Query.IfSome(func(v string) {
		query = types.StringValue(v)
	})

	return Policy{
		Bundle: bundle,
		Query:  query,
	}
}

func PolicySetAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"grant_if":  types.ListType{ElemType: PolicyType()},
		"revoke_if": types.ListType{ElemType: PolicyType()},
	}
}

func PolicySetFromView(policySet entity.PolicySet) (invalid PolicySet, diags diag.Diagnostics) {
	var (
		diags_   diag.Diagnostics
		grantIf  = types.ListNull(PolicyType())
		revokeIf = types.ListNull(PolicyType())
	)

	policySet.GrantIf.IfSome(func(policies []entity.Policy) {
		list := make([]attr.Value, 0, len(policies))
		for _, policy := range policies {
			p, diags_ := PolicyFromView(policy).ToObject()
			diags.Append(diags_...)
			if diags.HasError() {
				return
			}

			list = append(list, p)
		}

		grantIf, diags = types.ListValue(PolicyType(), list)
		diags.Append(diags_...)
		if diags.HasError() {
			return
		}
	})
	if diags.HasError() {
		return invalid, diags
	}

	policySet.RevokeIf.IfSome(func(policies []entity.Policy) {
		list := make([]attr.Value, 0, len(policies))
		for _, policy := range policies {
			p, diags_ := PolicyFromView(policy).ToObject()
			diags.Append(diags_...)
			if diags.HasError() {
				return
			}

			list = append(list, p)
		}

		revokeIf, diags = types.ListValue(PolicyType(), list)
		diags.Append(diags_...)
		if diags.HasError() {
			return
		}
	})
	if diags.HasError() {
		return invalid, diags
	}

	return PolicySet{
		GrantIf:  grantIf,
		RevokeIf: revokeIf,
	}, diags
}

func PolicySetFromObject(ctx context.Context, object types.Object) (invalid PolicySet, diags diag.Diagnostics) {
	var policySet PolicySet

	diags.Append(object.As(ctx, &policySet, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    false,
		UnhandledUnknownAsEmpty: false,
	})...)
	if diags.HasError() {
		return invalid, diags
	}

	return policySet, diags
}

func (self PolicySet) ToObject() (types.Object, diag.Diagnostics) {
	return types.ObjectValue(PolicySetAttrTypes(), map[string]attr.Value{
		"grant_if":  self.GrantIf,
		"revoke_if": self.RevokeIf,
	})
}

func (self PolicySet) ToView(ctx context.Context) (invalid entity.PolicySet, diags diag.Diagnostics) {
	var (
		grantIf  Option[[]entity.Policy]
		revokeIf Option[[]entity.Policy]
	)

	if !self.GrantIf.IsNull() {
		policies := make([]entity.Policy, 0, len(self.GrantIf.Elements()))

		for _, element := range self.GrantIf.Elements() {
			policy, err := entity.PolicyFromObject(ctx, element.(types.Object))
			if err != nil {
				diags.AddError("Unexpected", err.Error())
				return invalid, diags
			}

			policies = append(policies, policy)
		}

		grantIf = Some(policies)
	}

	if !self.RevokeIf.IsNull() {
		policies := make([]entity.Policy, 0, len(self.RevokeIf.Elements()))

		for _, element := range self.RevokeIf.Elements() {
			policy, err := entity.PolicyFromObject(ctx, element.(types.Object))
			if err != nil {
				diags.AddError("Unexpected", err.Error())
				return invalid, diags
			}

			policies = append(policies, policy)
		}

		revokeIf = Some(policies)
	}

	return entity.PolicySet{
		GrantIf:  grantIf,
		RevokeIf: revokeIf,
	}, diags
}
