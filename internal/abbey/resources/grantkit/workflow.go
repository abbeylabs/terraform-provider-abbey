package grantkit

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"abbey.so/terraform-provider-abbey/internal/abbey/resources/requestable"
)

type Workflow struct {
	Steps types.List `tfsdk:"steps"`
}

func (self Workflow) ToObject() (types.Object, diag.Diagnostics) {
	return types.ObjectValue(WorkflowAttrTypes(), map[string]attr.Value{
		"steps": self.Steps,
	})
}

func WorkflowAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"steps": types.ListType{
			ElemType: StepType(),
		},
	}
}

func WorkflowFromRequestableBuiltinWorkflow(builtinWorkflow requestable.BuiltinWorkflow) (*Workflow, diag.Diagnostics) {
	step, diags := StepFromRequestableBuiltinWorkflow(builtinWorkflow)
	if diags.HasError() {
		return nil, diags
	}

	stepObject, diags_ := step.ToObject()
	diags.Append(diags_...)
	if diags.HasError() {
		return nil, diags
	}

	steps, diags := types.ListValue(StepType(), []attr.Value{stepObject})

	return &Workflow{Steps: steps}, nil
}

func WorkflowFromObject(ctx context.Context, object types.Object) (*requestable.Workflow, diag.Diagnostics) {
	var (
		workflow Workflow
		diags    diag.Diagnostics
	)

	diags.Append(object.As(ctx, &workflow, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    false,
		UnhandledUnknownAsEmpty: false,
	})...)
	if diags.HasError() {
		return nil, diags
	}

	if workflow.Steps.IsNull() {
		return nil, nil
	}

	elements := workflow.Steps.Elements()

	if len(elements) < 1 {
		diags.AddError(
			"Invalid Input",
			fmt.Sprintf("Expected at least 1 review step, got %d.", len(elements)),
		)
		return nil, diags
	}

	var steps []Step

	diags.Append(workflow.Steps.ElementsAs(ctx, &steps, false)...)
	if diags.HasError() {
		return nil, diags
	}

	reviewSteps := make([]requestable.ReviewStep, 0, len(steps))

	for _, step := range steps {
		reviewStep, diags_ := step.ToReviewStep(ctx)
		diags.Append(diags_...)
		if diags.HasError() {
			return nil, diags
		}

		reviewSteps = append(reviewSteps, *reviewStep)
	}

	return &requestable.Workflow{Value: requestable.ReviewWorkflow{
		Steps: reviewSteps,
	}}, nil
}
