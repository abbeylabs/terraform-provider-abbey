package grantkit

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"abbey.so/terraform-provider-abbey/internal/abbey/resources/requestable"
)

type Model struct {
	Id          types.String `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Description types.String `tfsdk:"description"`
	Workflow    types.Object `tfsdk:"workflow"`
	Output      types.Object `tfsdk:"output"`

	// Policies    types.Object `tfsdk:"policies"`
}

func (m Model) ToRequestableView(ctx context.Context) (*requestable.View, diag.Diagnostics) {
	workflow, diags := WorkflowFromObject(ctx, m.Workflow)
	if diags.HasError() {
		return nil, diags
	}

	grant, diags_ := RequestableGrantFromOutputObject(ctx, m.Output)
	diags.Append(diags_...)
	if diags.HasError() {
		return nil, diags
	}

	return &requestable.View{
		Id:       m.Id.ValueString(),
		Name:     m.Name.ValueString(),
		Workflow: workflow,
		Grant:    grant,
	}, nil
}

func ModelFromRequestableView(view requestable.View) (*Model, diag.Diagnostics) {
	var (
		diags          diag.Diagnostics
		workflowObject = types.ObjectNull(WorkflowAttrTypes())
		outputObject   = types.ObjectNull(OutputAttrTypes())
	)

	if view.Workflow != nil {
		view.Workflow.Value.VisitWorkflow(requestable.WorkflowVisitor{
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
						var diags_ diag.Diagnostics

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

	return &Model{
		Id:          types.StringValue(view.Id),
		Name:        types.StringValue(view.Name),
		Description: types.StringNull(),
		Workflow:    workflowObject,
		Output:      outputObject,
	}, nil
}
