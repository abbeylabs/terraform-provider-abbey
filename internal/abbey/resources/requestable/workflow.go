package requestable

import (
	"encoding/json"
	"fmt"
)

const (
	workflowTypeBuiltin = "Builtin"

	builtinWorkflowTypeAllOf = "AllOf"
	builtinWorkflowTypeOneOf = "OneOf"
)

type (
	Workflow struct {
		value WorkflowEnum
	}

	WorkflowEnum interface {
		VisitWorkflow(WorkflowVisitor)
	}

	WorkflowVisitor struct {
		Builtin func(BuiltinWorkflow)
	}

	BuiltinWorkflow struct {
		value BuiltinWorkflowEnum
	}

	BuiltinWorkflowEnum interface {
		VisitBuiltinWorkflow(BuiltinWorkflowVisitor)
	}

	BuiltinWorkflowVisitor struct {
		AllOf func(BuiltinWorkflowAllOf)
		OneOf func(BuiltinWorkflowOneOf)
	}

	BuiltinWorkflowAllOf struct {
		Reviewers []UserQuery `json:"reviewers" tfsdk:"reviewers"`
	}

	BuiltinWorkflowOneOf struct {
		Reviewers []UserQuery `json:"reviewers" tfsdk:"reviewers"`
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

	w.value.VisitWorkflow(WorkflowVisitor{
		Builtin: func(workflow BuiltinWorkflow) {
			type_ = workflowTypeBuiltin
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
	default:
		return fmt.Errorf("unknown workflow type: %s", e.Type)
	}

	*w = Workflow{value: value}

	return nil
}

func (b *BuiltinWorkflow) VisitWorkflow(visitor WorkflowVisitor) {
	visitor.Builtin(*b)
}

func (w BuiltinWorkflow) MarshalJSON() ([]byte, error) {
	var (
		err   error
		type_ string
		value json.RawMessage
	)

	w.value.VisitBuiltinWorkflow(BuiltinWorkflowVisitor{
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

	*b = BuiltinWorkflow{value: value}

	return nil
}

func (b *BuiltinWorkflowAllOf) VisitBuiltinWorkflow(visitor BuiltinWorkflowVisitor) {
	visitor.AllOf(*b)
}

func (b *BuiltinWorkflowOneOf) VisitBuiltinWorkflow(visitor BuiltinWorkflowVisitor) {
	visitor.OneOf(*b)
}
