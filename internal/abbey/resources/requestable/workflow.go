package requestable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"abbey.so/terraform-provider-abbey/internal/abbey/entity"
)

const (
	workflowTypeBuiltin = "Builtin"

	builtinWorkflowTypeAllOf = "AllOf"
	builtinWorkflowTypeOneOf = "OneOf"
)

type (
	Workflow struct {
		Value WorkflowEnum
	}

	WorkflowEnum interface {
		VisitWorkflow(WorkflowVisitor)
	}

	WorkflowVisitor struct {
		Builtin  func(BuiltinWorkflow)
		GrantKit func(ReviewWorkflow)
	}

	BuiltinWorkflow struct {
		Value BuiltinWorkflowEnum
	}

	BuiltinWorkflowEnum interface {
		VisitBuiltinWorkflow(BuiltinWorkflowVisitor)
	}

	BuiltinWorkflowVisitor struct {
		AllOf func(BuiltinWorkflowAllOf)
		OneOf func(BuiltinWorkflowOneOf)
	}

	BuiltinWorkflowAllOf struct {
		Reviewers []UserQuery `json:"reviewers"`
	}

	BuiltinWorkflowOneOf struct {
		Reviewers []UserQuery `json:"reviewers"`
	}
)

var (
	_ WorkflowEnum        = (*BuiltinWorkflow)(nil)
	_ BuiltinWorkflowEnum = (*BuiltinWorkflowAllOf)(nil)
	_ BuiltinWorkflowEnum = (*BuiltinWorkflowOneOf)(nil)
)

func (w Workflow) MarshalJSON() ([]byte, error) {
	var (
		err   error
		type_ string
		value json.RawMessage
	)

	w.Value.VisitWorkflow(WorkflowVisitor{
		Builtin: func(workflow BuiltinWorkflow) {
			type_ = workflowTypeBuiltin
			value, err = json.Marshal(workflow)
		},
		GrantKit: func(workflow ReviewWorkflow) {
			type_ = "GrantKit"
			value, err = json.Marshal(workflow)
		},
	})
	if err != nil {
		return nil, err
	}

	return json.Marshal(enum{
		Type:  type_,
		Value: value,
	})
}

func (w *Workflow) UnmarshalJSON(b []byte) error {
	var e enum
	if err := json.Unmarshal(b, &e); err != nil {
		return err
	}

	var value WorkflowEnum

	switch e.Type {
	case workflowTypeBuiltin:
		var x BuiltinWorkflow
		if err := json.Unmarshal(e.Value, &x); err != nil {
			return err
		}

		value = &x
	case "GrantKit":
		var x ReviewWorkflow
		if err := json.Unmarshal(e.Value, &x); err != nil {
			return err
		}

		value = x
	default:
		return fmt.Errorf("unknown workflow type: %s", e.Type)
	}

	*w = Workflow{Value: value}

	return nil
}

func (b BuiltinWorkflow) VisitWorkflow(visitor WorkflowVisitor) {
	visitor.Builtin(b)
}

func (w BuiltinWorkflow) MarshalJSON() ([]byte, error) {
	var (
		err   error
		type_ string
		value json.RawMessage
	)

	w.Value.VisitBuiltinWorkflow(BuiltinWorkflowVisitor{
		AllOf: func(allOf BuiltinWorkflowAllOf) {
			type_ = builtinWorkflowTypeAllOf
			value, err = json.Marshal(allOf)
		},
		OneOf: func(oneOf BuiltinWorkflowOneOf) {
			type_ = builtinWorkflowTypeOneOf
			value, err = json.Marshal(oneOf)
		},
	})
	if err != nil {
		return nil, err
	}

	return json.Marshal(enum{
		Type:  type_,
		Value: value,
	})
}

func (b *BuiltinWorkflow) UnmarshalJSON(bs []byte) error {
	var e enum
	if err := json.Unmarshal(bs, &e); err != nil {
		return err
	}

	var value BuiltinWorkflowEnum

	switch e.Type {
	case builtinWorkflowTypeAllOf:
		var x BuiltinWorkflowAllOf
		if err := json.Unmarshal(e.Value, &x); err != nil {
			return err
		}

		value = &x
	case builtinWorkflowTypeOneOf:
		var x BuiltinWorkflowOneOf
		if err := json.Unmarshal(e.Value, &x); err != nil {
			return err
		}

		value = &x
	default:
		return fmt.Errorf("unknown builtin workflow type: %s", e.Type)
	}

	*b = BuiltinWorkflow{Value: value}

	return nil
}

func (b BuiltinWorkflowAllOf) VisitBuiltinWorkflow(visitor BuiltinWorkflowVisitor) {
	visitor.AllOf(b)
}

func (b BuiltinWorkflowOneOf) VisitBuiltinWorkflow(visitor BuiltinWorkflowVisitor) {
	visitor.OneOf(b)
}

type (
	ReviewWorkflow struct {
		Steps []ReviewStep `json:"steps"`
	}

	ReviewStep struct {
		Reviewers ReviewerQuantifierEnvelope `json:"reviewers"`
		SkipIf    []entity.Policy            `json:"skip_if"`
	}

	ReviewerQuantifierEnvelope struct {
		Value ReviewerQuantifier
	}

	ReviewerQuantifier interface {
		VisitReviewerQuantifier(ReviewerQuantifierVisitor)
	}

	ReviewerQuantifierVisitor struct {
		AllOf func([]string)
		OneOf func([]string)
	}

	ReviewerQuantifierAllOf []string
	ReviewerQuantifierOneOf []string
)

func ReviewWorkflowAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"steps": types.ListType{
			types.ObjectType{
				ReviewStepAttrTypes(),
			},
		},
	}
}

func (self ReviewerQuantifierEnvelope) ToObject() (obj types.Object, diags diag.Diagnostics) {
	self.Value.VisitReviewerQuantifier(ReviewerQuantifierVisitor{
		AllOf: func(ss []string) {
			var values []attr.Value
			for _, s := range ss {
				values = append(values, types.StringValue(s))
			}

			obj, diags = types.ObjectValue(
				map[string]attr.Type{
					"all_of": types.ListType{
						ElemType: types.StringType,
					},
					"one_of": types.ListType{
						ElemType: types.StringType,
					},
				},
				map[string]attr.Value{
					"all_of": types.ListValueMust(types.StringType, values),
					"one_of": types.ListNull(types.StringType),
				},
			)
		},
		OneOf: func(ss []string) {
			var values []attr.Value
			for _, s := range ss {
				values = append(values, types.StringValue(s))
			}

			obj, diags = types.ObjectValue(
				map[string]attr.Type{
					"all_of": types.ListType{
						ElemType: types.StringType,
					},
					"one_of": types.ListType{
						ElemType: types.StringType,
					},
				},
				map[string]attr.Value{
					"one_of": types.ListValueMust(types.StringType, values),
					"all_of": types.ListNull(types.StringType),
				},
			)
		},
	})
	return obj, diags
}

func (self ReviewWorkflow) ToObject() (invalid types.Object, diags diag.Diagnostics) {
	var steps []attr.Value
	for _, step := range self.Steps {
		obj, diags_ := step.ToObject()
		diags.Append(diags_...)
		if diags.HasError() {
			return invalid, diags
		}

		steps = append(steps, obj)
	}

	return types.ObjectValue(
		ReviewWorkflowAttrTypes(),
		map[string]attr.Value{
			"steps": types.ListValueMust(types.ObjectType{ReviewStepAttrTypes()}, steps),
		},
	)
}

func ReviewStepAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"reviewers": types.ObjectType{map[string]attr.Type{
			"all_of": types.ListType{
				ElemType: types.StringType,
			},
			"one_of": types.ListType{
				ElemType: types.StringType,
			},
		}},
		"skip_if": types.ListType{
			types.ObjectType{
				map[string]attr.Type{
					"bundle": types.StringType,
					"query":  types.StringType,
				},
			},
		},
	}
}

func (self ReviewStep) ToObject() (invalid types.Object, diags diag.Diagnostics) {
	var (
		skipIfs     []attr.Value
		skipIfValue attr.Value
	)

	reviewers, diags_ := self.Reviewers.ToObject()
	diags.Append(diags_...)
	if diags.HasError() {
		return invalid, diags
	}

	policyType := types.ObjectType{
		map[string]attr.Type{
			"bundle": types.StringType,
			"query":  types.StringType,
		},
	}

	for _, skipIf := range self.SkipIf {
		obj, diags_ := types.ObjectValue(
			map[string]attr.Type{
				"bundle": types.StringType,
				"query":  types.StringType,
			},
			map[string]attr.Value{
				"bundle": types.StringValue(skipIf.Bundle.Unwrap()),
				"query":  types.StringValue(skipIf.Query.Unwrap()),
			},
		)
		diags.Append(diags_...)
		if diags.HasError() {
			return invalid, diags
		}

		skipIfs = append(skipIfs, obj)
	}

	if skipIfs == nil {
		skipIfValue = types.ListNull(policyType)
	} else {
		skipIfValue = types.ListValueMust(policyType, skipIfs)
	}

	return types.ObjectValue(
		ReviewStepAttrTypes(),
		map[string]attr.Value{
			"reviewers": reviewers,
			"skip_if":   skipIfValue,
		},
	)
}

func (self ReviewWorkflow) VisitWorkflow(visitor WorkflowVisitor) {
	visitor.GrantKit(self)
}

func (self ReviewerQuantifierAllOf) VisitReviewerQuantifier(visitor ReviewerQuantifierVisitor) {
	visitor.AllOf(self)
}

func (self ReviewerQuantifierOneOf) VisitReviewerQuantifier(visitor ReviewerQuantifierVisitor) {
	visitor.OneOf(self)
}

func (self ReviewerQuantifierEnvelope) MarshalJSON() ([]byte, error) {
	var (
		tag string
	)

	self.Value.VisitReviewerQuantifier(ReviewerQuantifierVisitor{
		AllOf: func([]string) { tag = "AllOf" },
		OneOf: func([]string) { tag = "OneOf" },
	})

	content, err := json.Marshal(self.Value)
	if err != nil {
		return nil, err
	}

	return json.Marshal(enum{
		Type:  tag,
		Value: content,
	})
}

func (self *ReviewerQuantifierEnvelope) UnmarshalJSON(data []byte) error {
	var e enum
	if err := json.Unmarshal(data, &e); err != nil {
		return err
	}

	var value ReviewerQuantifier

	switch e.Type {
	case "AllOf":
		var x ReviewerQuantifierAllOf
		if err := json.Unmarshal(e.Value, &x); err != nil {
			return err
		}

		value = x
	case "OneOf":
		var x ReviewerQuantifierOneOf
		if err := json.Unmarshal(e.Value, &x); err != nil {
			return err
		}

		value = x
	default:
		return fmt.Errorf("unknown review workflow type: %s", e.Type)
	}

	*self = ReviewerQuantifierEnvelope{Value: value}

	return nil
}
