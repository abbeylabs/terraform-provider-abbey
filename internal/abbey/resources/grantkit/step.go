package grantkit

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"abbey.so/terraform-provider-abbey/internal/abbey/resources/requestable"
)

type Step struct {
	Reviewers types.Object `tfsdk:"reviewers"`
	SkipIf    types.List   `tfsdk:"skip_if"`
}

func StepAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"reviewers": ReviewerSpecType(),
		"skip_if":   types.ListType{ElemType: PolicyType()},
	}
}

func StepType() attr.Type {
	return types.ObjectType{AttrTypes: StepAttrTypes()}
}

func StepFromRequestableBuiltinWorkflow(
	builtinWorkflow requestable.BuiltinWorkflow,
) (step *Step, diags diag.Diagnostics) {
	var reviewers ReviewerSpec

	builtinWorkflow.Value.VisitBuiltinWorkflow(requestable.BuiltinWorkflowVisitor{
		AllOf: func(allOf requestable.BuiltinWorkflowAllOf) {
			steps := make([]attr.Value, 0, len(allOf.Reviewers))

			for _, reviewer := range allOf.Reviewers {
				reviewer.Value.VisitUserQuery(requestable.UserQueryVisitor{
					AuthId: func(authId requestable.AuthId) {
						steps = append(steps, types.StringValue(authId.Value))
					},
				})
			}

			allOfValue, diags_ := types.ListValue(types.StringType, steps)
			diags.Append(diags_...)
			if diags.HasError() {
				return
			}

			reviewers = ReviewerSpec{
				AllOf: allOfValue,
				OneOf: types.ListNull(types.StringType),
			}
		},
		OneOf: func(oneOf requestable.BuiltinWorkflowOneOf) {
			steps := make([]attr.Value, 0, len(oneOf.Reviewers))

			for _, reviewer := range oneOf.Reviewers {
				reviewer.Value.VisitUserQuery(requestable.UserQueryVisitor{
					AuthId: func(authId requestable.AuthId) {
						steps = append(steps, types.StringValue(authId.Value))
					},
				})
			}

			oneOfValue, diags_ := types.ListValue(types.StringType, steps)
			diags.Append(diags_...)
			if diags.HasError() {
				return
			}

			reviewers = ReviewerSpec{
				AllOf: types.ListNull(types.StringType),
				OneOf: oneOfValue,
			}
		},
	})
	if diags.HasError() {
		return nil, diags
	}

	reviewersValue, diags_ := reviewers.ToObject()
	diags.Append(diags_...)
	if diags.HasError() {
		return nil, diags
	}

	return &Step{
		Reviewers: reviewersValue,
		SkipIf:    types.ListNull(PolicyType()),
	}, nil
}

func (self Step) ToBuiltinWorkflowView(ctx context.Context) (*requestable.BuiltinWorkflow, diag.Diagnostics) {
	var (
		diags     diag.Diagnostics
		reviewers ReviewerSpec
	)

	if !self.SkipIf.IsNull() {
		diags.AddWarning("Future Behavior May Change", "`skip_if` is unimplemented in review workflow. Coming soon!")
	}

	diags.Append(self.Reviewers.As(ctx, &reviewers, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    false,
		UnhandledUnknownAsEmpty: false,
	})...)
	if diags.HasError() {
		return nil, diags
	}

	builtinWorkflow, diags_ := reviewers.ToBuiltinWorkflowView()
	diags.Append(diags_...)

	return builtinWorkflow, diags
}

func (self Step) ToObject() (types.Object, diag.Diagnostics) {
	return types.ObjectValue(StepAttrTypes(), map[string]attr.Value{
		"reviewers": self.Reviewers,
		"skip_if":   self.SkipIf,
	})
}

type ReviewerSpec struct {
	AllOf types.List `tfsdk:"all_of"`
	OneOf types.List `tfsdk:"one_of"`
}

func ReviewerSpecAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"all_of": types.ListType{ElemType: types.StringType},
		"one_of": types.ListType{ElemType: types.StringType},
	}
}

func ReviewerSpecType() attr.Type {
	return types.ObjectType{AttrTypes: ReviewerSpecAttrTypes()}
}

func (self ReviewerSpec) ToObject() (types.Object, diag.Diagnostics) {
	return types.ObjectValue(ReviewerSpecAttrTypes(), map[string]attr.Value{
		"all_of": self.AllOf,
		"one_of": self.OneOf,
	})
}

func (self ReviewerSpec) ToBuiltinWorkflowView() (*requestable.BuiltinWorkflow, diag.Diagnostics) {
	var (
		diags diag.Diagnostics
		enum  requestable.BuiltinWorkflowEnum
	)

	if !self.AllOf.IsNull() && !self.OneOf.IsNull() {
		diags.AddError("Invalid Input", "Only one of `all_of` and `one_of` expected.")
		return nil, diags
	}

	if !self.AllOf.IsNull() {
		elements := self.AllOf.Elements()
		userQueries := make([]requestable.UserQuery, 0, len(elements))

		for _, element := range elements {
			userQueries = append(userQueries, requestable.UserQuery{
				Value: requestable.AuthId{
					Value: element.(types.String).ValueString(),
				},
			})
		}

		enum = requestable.BuiltinWorkflowAllOf{Reviewers: userQueries}
	}

	if !self.OneOf.IsNull() {
		elements := self.OneOf.Elements()
		userQueries := make([]requestable.UserQuery, 0, len(elements))

		for _, element := range elements {
			userQueries = append(userQueries, requestable.UserQuery{
				Value: requestable.AuthId{
					Value: element.(types.String).ValueString(),
				},
			})
		}

		enum = requestable.BuiltinWorkflowOneOf{Reviewers: userQueries}
	}

	return &requestable.BuiltinWorkflow{Value: enum}, nil
}
