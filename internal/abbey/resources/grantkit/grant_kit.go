package grantkit

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/moznion/go-optional"

	"abbey.so/terraform-provider-abbey/internal/abbey/entity"
	"abbey.so/terraform-provider-abbey/internal/abbey/resources/requestable"
)

type Model struct {
	Id          types.String `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Description types.String `tfsdk:"description"`
	Workflow    types.Object `tfsdk:"workflow"`
	Output      types.Object `tfsdk:"output"`
	Policies    types.Object `tfsdk:"policies"`
}

func (self Model) ToRequestableView(ctx context.Context) (*requestable.View, diag.Diagnostics) {
	workflow, diags := WorkflowFromObject(ctx, self.Workflow)
	if diags.HasError() {
		return nil, diags
	}

	grant, diags_ := RequestableGrantFromOutputObject(ctx, self.Output)
	diags.Append(diags_...)
	if diags.HasError() {
		return nil, diags
	}

	var policies Option[entity.PolicySet]

	if !self.Policies.IsNull() {
		policySet, diags_ := PolicySetFromObject(ctx, self.Policies)
		diags.Append(diags_...)
		if diags.HasError() {
			return nil, diags
		}

		policies_, diags_ := policySet.ToView(ctx)
		diags.Append(diags_...)
		if diags.HasError() {
			return nil, diags
		}

		policies = Some(policies_)
	}

	return &requestable.View{
		Id:          self.Id.ValueString(),
		Name:        self.Name.ValueString(),
		Description: self.Description.ValueString(),
		Workflow:    workflow,
		Grant:       grant,
		Policies:    policies,
	}, nil
}

func ModelFromRequestableView(view requestable.View) (*Model, diag.Diagnostics) {
	var (
		diags          diag.Diagnostics
		diags_         diag.Diagnostics
		workflowObject = types.ObjectNull(WorkflowAttrTypes())
		outputObject   = types.ObjectNull(OutputAttrTypes())
		policiesObject = types.ObjectNull(PolicySetAttrTypes())
	)

	if view.Workflow != nil {
		view.Workflow.Value.VisitWorkflow(requestable.WorkflowVisitor{
			GrantKit: func(reviewWorkflow requestable.ReviewWorkflow) {
				workflowObject, diags_ = reviewWorkflow.ToObject()
				diags.Append(diags_...)
				if diags.HasError() {
					return
				}
			},
			Builtin: func(builtinWorkflow requestable.BuiltinWorkflow) {
				workflow, diags_ := WorkflowFromRequestableBuiltinWorkflow(builtinWorkflow)
				diags.Append(diags_...)
				if diags.HasError() {
					return
				}

				workflowObject, diags_ = workflow.ToObject()
				diags.Append(diags_...)
				if diags.HasError() {
					return
				}
			},
		})
	}
	if diags.HasError() {
		return nil, diags
	}

	if view.Grant != nil {
		view.Grant.Value.VisitGrant(requestable.GrantVisitor{
			Generate: func(generate requestable.GenerateGrant) {
				generate.Value.VisitGenerateGrant(requestable.GenerateGrantVisitor{
					Github: func(github requestable.GithubGenerateDestination) {
						output := OutputFromRequestableGithubDestination(github)
						outputObject, diags_ = output.ToObject()
						diags.Append(diags_...)
						if diags.HasError() {
							return
						}
					},
				})
			},
		})
	}
	if diags.HasError() {
		return nil, diags
	}

	view.Policies.IfSome(func(view entity.PolicySet) {
		policySet, diags_ := PolicySetFromView(view)
		diags.Append(diags_...)
		if diags.HasError() {
			return
		}

		policiesObject, diags_ = policySet.ToObject()
		diags.Append(diags_...)
		if diags.HasError() {
			return
		}
	})
	if diags.HasError() {
		return nil, diags
	}

	return &Model{
		Id:          types.StringValue(view.Id),
		Name:        types.StringValue(view.Name),
		Description: types.StringValue(view.Description),
		Workflow:    workflowObject,
		Output:      outputObject,
		Policies:    policiesObject,
	}, nil
}
