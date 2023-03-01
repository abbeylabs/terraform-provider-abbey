package requestable

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	. "github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

type WorkflowType struct{}

func (w WorkflowType) TerraformType(context.Context) tftypes.Type {
	return tftypes.Object{
		AttributeTypes: map[string]tftypes.Type{
			workflowTypeBuiltinTf: BuiltinWorkflowTfTypesType(),
		},
		OptionalAttributes: nil,
	}
}

func (t WorkflowType) ValueFromTerraform(_ context.Context, value tftypes.Value) (value_ attr.Value, err error) {
	if !value.IsKnown() {
		return NewUnknownWorkflow(), nil
	}

	var m *map[string]tftypes.Value
	if err := value.As(&m); err != nil {
		return nil, err
	}

	if m == nil {
		return NewNullWorkflow(), nil
	}

	var inner WorkflowEnum

	for key, val := range *m {
		switch key {
		case workflowTypeBuiltinTf:
			inner_, err := BuiltinWorkflowFromTfTypesValue(val)
			if err != nil {
				return nil, err
			}
			if inner_ == nil {
				continue
			}

			inner = inner_
		default:
			return value_, fmt.Errorf("unknown key: %s", key)
		}
	}
	if err != nil {
		return value_, err
	}

	return NewWorkflow(Workflow{value: inner}), nil
}

func (t WorkflowType) ValueType(context.Context) attr.Value {
	var wtf WorkflowTf
	return wtf
}

func (t WorkflowType) Equal(ty attr.Type) bool {
	_, ok := ty.(WorkflowType)
	return ok
}

func (WorkflowType) String() string {
	return "Workflow"
}

func (w WorkflowType) ApplyTerraform5AttributePathStep(step tftypes.AttributePathStep) (interface{}, error) {
	attrName, ok := step.(tftypes.AttributeName)
	if !ok {
		return nil, fmt.Errorf("cannot apply step %T to WorkflowType", step)
	}

	switch string(attrName) {
	case workflowTypeBuiltinTf:
		return BuiltinWorkflowTfTypesType(), nil
	default:
		return nil, fmt.Errorf("undefined attribute name %s in WorkflowType", attrName)
	}
}

func (t WorkflowType) ValueFromObject(
	ctx context.Context,
	value basetypes.ObjectValue,
) (basetypes.ObjectValuable, Diagnostics) {
	var w WorkflowTf

	diags := value.As(ctx, &w, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    false,
		UnhandledUnknownAsEmpty: false,
	})
	if diags.HasError() {
		return nil, diags
	}

	return w, diags
}
