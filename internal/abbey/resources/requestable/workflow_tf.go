package requestable

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	. "github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
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
	var w Workflow

	if !value.IsFullyKnown() {
		return WorkflowTf{
			Workflow: w,
			valid:    false,
		}, nil
	}

	var m map[string]tftypes.Value
	if err := value.As(&m); err != nil {
		return nil, err
	}

	if len(m) == 0 {
		return WorkflowTf{
			Workflow: w,
			valid:    false,
		}, nil
	}

	var inner WorkflowEnum

	for key, val := range m {
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

	return WorkflowTf{Workflow: Workflow{value: inner}, valid: true}, nil
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

	return &w, diags
}

type WorkflowTf struct {
	Workflow

	valid bool
}

func (w WorkflowTf) ToObjectValue(ctx context.Context) (object basetypes.ObjectValue, diags Diagnostics) {
	var (
		builtin      BuiltinWorkflow
		builtinValue attr.Value
	)

	if !w.valid {
		return types.ObjectNull(w.AttrTypes(ctx)), nil
	}

	w.value.VisitWorkflow(WorkflowVisitor{
		Builtin: func(workflow BuiltinWorkflow) {
			builtinValue, diags = workflow.ToObjectValue(ctx)
		},
	})
	if diags.HasError() {
		return object, diags
	}

	return types.ObjectValue(
		map[string]attr.Type{
			workflowTypeBuiltinTf: builtin.Type(ctx),
		},
		map[string]attr.Value{
			workflowTypeBuiltinTf: builtinValue,
		},
	)
}

func (w WorkflowTf) AttrTypes(ctx context.Context) map[string]attr.Type {
	var builtin BuiltinWorkflow

	return map[string]attr.Type{
		workflowTypeBuiltinTf: builtin.Type(ctx),
	}
}

func (w WorkflowTf) Type(context.Context) attr.Type {
	return WorkflowType{}
}

func (w WorkflowTf) ToTerraformValue(ctx context.Context) (value tftypes.Value, err error) {
	var (
		builtinValue tftypes.Value
		type_        WorkflowType
	)

	w.value.VisitWorkflow(WorkflowVisitor{
		Builtin: func(workflow BuiltinWorkflow) {
			builtinValue, err = workflow.ToTerraformValue(ctx)
		},
	})
	if err != nil {
		return value, err
	}

	return tftypes.NewValue(
		type_.TerraformType(ctx),
		map[string]tftypes.Value{
			workflowTypeBuiltinTf: builtinValue,
		},
	), nil
}

func (w WorkflowTf) Equal(value attr.Value) bool {
	rhs, err := value.ToTerraformValue(context.Background())
	if err != nil {
		return false
	}

	lhs, err := w.ToTerraformValue(context.Background())
	if err != nil {
		return false
	}

	return lhs.Equal(rhs)
}

func (w WorkflowTf) IsNull() (defined bool) {
	w.value.VisitWorkflow(WorkflowVisitor{
		Builtin: func(workflow BuiltinWorkflow) {
			defined = true
		},
	})

	return !defined
}

func (w WorkflowTf) IsUnknown() (defined bool) {
	w.value.VisitWorkflow(WorkflowVisitor{
		Builtin: func(workflow BuiltinWorkflow) {
			defined = true
		},
	})

	return !defined
}

func (w WorkflowTf) String() string {
	var inner string

	w.value.VisitWorkflow(WorkflowVisitor{
		Builtin: func(workflow BuiltinWorkflow) {
			inner = workflow.String()
		},
	})

	return fmt.Sprintf("Workflow{%s}", inner)
}

func BuiltinWorkflowTfTypesType() tftypes.Type {
	return tftypes.Object{
		AttributeTypes: map[string]tftypes.Type{
			builtinWorkflowTypeAllOfTf: BuiltinWorkflowAllOfTfTypesType(),
			builtinWorkflowTypeOneOfTf: BuiltinWorkflowOneOfTfTypesType(),
		},
		OptionalAttributes: nil,
	}
}

func (w BuiltinWorkflow) ToObjectValue(ctx context.Context) (object basetypes.ObjectValue, diags Diagnostics) {
	var (
		allOf      BuiltinWorkflowAllOf
		oneOf      BuiltinWorkflowOneOf
		allOfValue attr.Value = types.ObjectNull(allOf.AttrTypes(ctx))
		oneOfValue attr.Value = types.ObjectNull(oneOf.AttrTypes(ctx))
	)

	w.value.VisitBuiltinWorkflow(BuiltinWorkflowVisitor{
		AllOf: func(allOf BuiltinWorkflowAllOf) {
			allOfValue, diags = allOf.ToObjectValue(ctx)
		},
		OneOf: func(oneOf BuiltinWorkflowOneOf) {
			oneOfValue, diags = oneOf.ToObjectValue(ctx)
		},
	})
	if diags.HasError() {
		return object, diags
	}

	return types.ObjectValue(
		map[string]attr.Type{
			builtinWorkflowTypeAllOfTf: allOf.Type(ctx),
			builtinWorkflowTypeOneOfTf: oneOf.Type(ctx),
		},
		map[string]attr.Value{
			builtinWorkflowTypeAllOfTf: allOfValue,
			builtinWorkflowTypeOneOfTf: oneOfValue,
		},
	)
}

func (w BuiltinWorkflow) Type(ctx context.Context) attr.Type {
	var (
		allOf BuiltinWorkflowAllOf
		oneOf BuiltinWorkflowOneOf
	)

	return types.ObjectType{AttrTypes: map[string]attr.Type{
		builtinWorkflowTypeAllOfTf: allOf.Type(ctx),
		builtinWorkflowTypeOneOfTf: oneOf.Type(ctx),
	}}
}

func (w BuiltinWorkflow) ToTerraformValue(ctx context.Context) (value tftypes.Value, err error) {
	var (
		allOfValue = tftypes.NewValue(BuiltinWorkflowAllOfTfTypesType(), nil)
		oneOfValue = tftypes.NewValue(BuiltinWorkflowOneOfTfTypesType(), nil)
	)

	w.value.VisitBuiltinWorkflow(BuiltinWorkflowVisitor{
		AllOf: func(allOf BuiltinWorkflowAllOf) {
			allOfValue, err = allOf.ToTerraformValue(ctx)
		},
		OneOf: func(oneOf BuiltinWorkflowOneOf) {
			oneOfValue, err = oneOf.ToTerraformValue(ctx)
		},
	})
	if err != nil {
		return value, err
	}

	return tftypes.NewValue(
		BuiltinWorkflowTfTypesType(),
		map[string]tftypes.Value{
			builtinWorkflowTypeAllOfTf: allOfValue,
			builtinWorkflowTypeOneOfTf: oneOfValue,
		},
	), nil
}

func (w BuiltinWorkflow) Equal(value attr.Value) bool {
	rhs, err := value.ToTerraformValue(context.Background())
	if err != nil {
		return false
	}

	lhs, err := w.ToTerraformValue(context.Background())
	if err != nil {
		return false
	}

	return lhs.Equal(rhs)
}

func (w BuiltinWorkflow) IsNull() (defined bool) {
	w.value.VisitBuiltinWorkflow(BuiltinWorkflowVisitor{
		AllOf: func(BuiltinWorkflowAllOf) {
			defined = true
		},
		OneOf: func(BuiltinWorkflowOneOf) {
			defined = true
		},
	})

	return !defined
}

func (w BuiltinWorkflow) IsUnknown() (defined bool) {
	w.value.VisitBuiltinWorkflow(BuiltinWorkflowVisitor{
		AllOf: func(BuiltinWorkflowAllOf) {
			defined = true
		},
		OneOf: func(BuiltinWorkflowOneOf) {
			defined = true
		},
	})

	return !defined
}

func (w BuiltinWorkflow) String() string {
	var inner string

	w.value.VisitBuiltinWorkflow(BuiltinWorkflowVisitor{
		AllOf: func(allOf BuiltinWorkflowAllOf) {
			inner = allOf.String()
		},
		OneOf: func(oneOf BuiltinWorkflowOneOf) {
			inner = oneOf.String()
		},
	})

	return fmt.Sprintf("BuiltinWorkflow{%s}", inner)
}

func BuiltinWorkflowAllOfTfTypesType() tftypes.Type {
	return tftypes.Object{
		AttributeTypes: map[string]tftypes.Type{
			"reviewers": tftypes.List{ElementType: UserQueryTfTypesType()},
		},
		OptionalAttributes: nil,
	}
}

func (w BuiltinWorkflowAllOf) AttrTypes(ctx context.Context) map[string]attr.Type {
	var userQuery UserQuery

	return map[string]attr.Type{
		"reviewers": types.ListType{
			ElemType: userQuery.Type(ctx),
		},
	}
}

func (w BuiltinWorkflowAllOf) Type(ctx context.Context) attr.Type {
	return types.ObjectType{AttrTypes: w.AttrTypes(ctx)}
}

func (w BuiltinWorkflowAllOf) ToTerraformValue(ctx context.Context) (value tftypes.Value, err error) {
	reviewers := make([]tftypes.Value, 0, len(w.Reviewers))
	for _, reviewer := range w.Reviewers {
		v, err := reviewer.ToTerraformValue(ctx)
		if err != nil {
			return value, err
		}

		reviewers = append(reviewers, v)
	}

	return tftypes.NewValue(
		BuiltinWorkflowAllOfTfTypesType(),
		map[string]tftypes.Value{
			"reviewers": tftypes.NewValue(
				tftypes.List{ElementType: UserQueryTfTypesType()},
				reviewers,
			),
		},
	), nil
}

func (w BuiltinWorkflowAllOf) Equal(value attr.Value) bool {
	rhs, err := value.ToTerraformValue(context.Background())
	if err != nil {
		return false
	}

	lhs, err := w.ToTerraformValue(context.Background())
	if err != nil {
		return false
	}

	return lhs.Equal(rhs)
}

func (w BuiltinWorkflowAllOf) IsNull() (defined bool) {
	return false
}

func (w BuiltinWorkflowAllOf) IsUnknown() (defined bool) {
	return false
}

func (w BuiltinWorkflowAllOf) String() string {
	elemStrs := make([]string, 0, len(w.Reviewers))
	for _, reviewer := range w.Reviewers {
		elemStrs = append(elemStrs, reviewer.String())
	}

	return fmt.Sprintf("BuiltinWorkflowAllOf{Reviewers: [%s]}", strings.Join(elemStrs, ", "))
}

func (w BuiltinWorkflowAllOf) ToObjectValue(context.Context) (object basetypes.ObjectValue, diags Diagnostics) {
	reviewerValues := make([]attr.Value, 0, len(w.Reviewers))
	for _, reviewer := range w.Reviewers {
		reviewerValues = append(reviewerValues, reviewer)
	}

	reviewersValue, diags := basetypes.NewListValue(UserQueryType(), reviewerValues)
	if diags.HasError() {
		return object, diags
	}

	return types.ObjectValue(
		map[string]attr.Type{
			"reviewers": types.ListType{
				ElemType: UserQueryType(),
			},
		},
		map[string]attr.Value{
			"reviewers": reviewersValue,
		},
	)
}

func BuiltinWorkflowOneOfTfTypesType() tftypes.Type {
	return tftypes.Object{
		AttributeTypes: map[string]tftypes.Type{
			"reviewers": tftypes.List{ElementType: UserQueryTfTypesType()},
		},
		OptionalAttributes: nil,
	}
}

func (w BuiltinWorkflowOneOf) AttrTypes(ctx context.Context) map[string]attr.Type {
	var userQuery UserQuery

	return map[string]attr.Type{
		"reviewers": types.ListType{
			ElemType: userQuery.Type(ctx),
		},
	}
}

func (w BuiltinWorkflowOneOf) Type(ctx context.Context) attr.Type {
	var userQuery UserQuery

	return types.ObjectType{AttrTypes: map[string]attr.Type{
		"reviewers": types.ListType{
			ElemType: userQuery.Type(ctx),
		},
	}}
}

func (w BuiltinWorkflowOneOf) ToTerraformValue(ctx context.Context) (value tftypes.Value, err error) {
	reviewers := make([]tftypes.Value, 0, len(w.Reviewers))
	for _, reviewer := range w.Reviewers {
		v, err := reviewer.ToTerraformValue(ctx)
		if err != nil {
			return value, err
		}

		reviewers = append(reviewers, v)
	}

	return tftypes.NewValue(
		BuiltinWorkflowAllOfTfTypesType(),
		map[string]tftypes.Value{
			"reviewers": tftypes.NewValue(
				tftypes.List{ElementType: UserQueryTfTypesType()},
				reviewers,
			),
		},
	), nil
}

func (w BuiltinWorkflowOneOf) Equal(value attr.Value) bool {
	rhs, err := value.ToTerraformValue(context.Background())
	if err != nil {
		return false
	}

	lhs, err := w.ToTerraformValue(context.Background())
	if err != nil {
		return false
	}

	return lhs.Equal(rhs)
}

func (w BuiltinWorkflowOneOf) IsNull() (defined bool) {
	return false
}

func (w BuiltinWorkflowOneOf) IsUnknown() (defined bool) {
	return false
}

func (w BuiltinWorkflowOneOf) String() string {
	elemStrs := make([]string, 0, len(w.Reviewers))
	for _, reviewer := range w.Reviewers {
		elemStrs = append(elemStrs, reviewer.String())
	}

	return fmt.Sprintf("BuiltinWorkflowOneOf{Reviewers: [%s]}", strings.Join(elemStrs, ", "))
}

func (w BuiltinWorkflowOneOf) ToObjectValue(context.Context) (object basetypes.ObjectValue, diags Diagnostics) {
	reviewerValues := make([]attr.Value, 0, len(w.Reviewers))
	for _, reviewer := range w.Reviewers {
		reviewerValues = append(reviewerValues, reviewer)
	}

	reviewersValue, diags := basetypes.NewListValue(UserQueryType(), reviewerValues)
	if diags.HasError() {
		return object, diags
	}

	return types.ObjectValue(
		map[string]attr.Type{
			"reviewers": types.ListType{
				ElemType: UserQueryType(),
			},
		},
		map[string]attr.Value{
			"reviewers": reviewersValue,
		},
	)
}
