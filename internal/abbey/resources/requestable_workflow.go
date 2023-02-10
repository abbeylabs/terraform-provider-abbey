package resources

import (
	"encoding/json"
	"fmt"
)

type enum struct {
	Type  string          `json:"type"`
	Value json.RawMessage `json:"value"`
}

type (
	WorkflowType        string
	BuiltinWorkflowType string
)

const (
	WorkflowTypeBuiltin WorkflowType = "Builtin"

	BuiltinWorkflowTypeOneOf BuiltinWorkflowType = "OneOf"
)

func (t WorkflowType) String() string {
	return string(t)
}

func (t BuiltinWorkflowType) String() string {
	return string(t)
}

type WorkflowEnum struct {
	Type  WorkflowType `json:"type"`
	Value Workflow     `json:"value"`
}

var _ Workflow = (*WorkflowEnum)(nil)

func (w *WorkflowEnum) ToWorkflowEnum() WorkflowEnum {
	return *w
}

func (w *WorkflowEnum) VisitWorkflow(visitor WorkflowVisitor) {
	w.Value.VisitWorkflow(visitor)
}

func (w *WorkflowEnum) UnmarshalJSON(b []byte) error {
	var e enum
	if err := json.Unmarshal(b, &e); err != nil {
		return err
	}

	switch e.Type {
	case "":
	case WorkflowTypeBuiltin.String():
		var x BuiltinWorkflowEnum
		if err := json.Unmarshal(e.Value, &x); err != nil {
			return err
		}

		w.Type = WorkflowTypeBuiltin
		w.Value = &x
	default:
		return fmt.Errorf("unknown workflow type: %s", e.Type)
	}

	return nil
}

type Workflow interface {
	ToWorkflowEnum() WorkflowEnum
	VisitWorkflow(WorkflowVisitor)
}

type WorkflowVisitor struct {
	VisitBuiltinWorkflow func(BuiltinWorkflow)
}

type BuiltinWorkflowEnum struct {
	Type  BuiltinWorkflowType `json:"type"`
	Value BuiltinWorkflow     `json:"value"`
}

var (
	_ Workflow        = (*BuiltinWorkflowEnum)(nil)
	_ BuiltinWorkflow = (*BuiltinWorkflowEnum)(nil)
)

func (b *BuiltinWorkflowEnum) ToEnum() WorkflowEnum {
	return WorkflowEnum{
		Type:  WorkflowTypeBuiltin,
		Value: b,
	}
}

func (b *BuiltinWorkflowEnum) ToWorkflowEnum() WorkflowEnum {
	return b.ToEnum()
}

func (b *BuiltinWorkflowEnum) VisitWorkflow(visitor WorkflowVisitor) {
	visitor.VisitBuiltinWorkflow(b.Value)
}

func (b *BuiltinWorkflowEnum) VisitBuiltinWorkflow(visitor BuiltinWorkflowVisitor) {
	b.Value.VisitBuiltinWorkflow(visitor)
}

func (b *BuiltinWorkflowEnum) UnmarshalJSON(bs []byte) error {
	var e enum
	if err := json.Unmarshal(bs, &e); err != nil {
		return err
	}

	switch e.Type {
	case "":
	case BuiltinWorkflowTypeOneOf.String():
		var x BuiltinWorkflowOneOf
		if err := json.Unmarshal(e.Value, &x); err != nil {
			return err
		}

		b.Type = BuiltinWorkflowTypeOneOf
		b.Value = &x
	default:
		return fmt.Errorf("unknown builtin workflow type: %s", e.Type)
	}

	return nil
}

type BuiltinWorkflow interface {
	VisitBuiltinWorkflow(BuiltinWorkflowVisitor)
}

type BuiltinWorkflowVisitor struct {
	VisitOneOf func(BuiltinWorkflowOneOf)
	VisitAllOf func(BuiltinWorkflowAllOf)
}

type BuiltinWorkflowAllOf struct {
	Reviewers []UserQueryEnum `json:"reviewers"`
}

type BuiltinWorkflowOneOf struct {
	Reviewers []UserQueryEnum `json:"reviewers"`
}

var _ BuiltinWorkflow = (*BuiltinWorkflowOneOf)(nil)

func (b *BuiltinWorkflowOneOf) VisitBuiltinWorkflow(visitor BuiltinWorkflowVisitor) {
	visitor.VisitOneOf(*b)
}

func (b *BuiltinWorkflowOneOf) ToEnum() BuiltinWorkflowEnum {
	return BuiltinWorkflowEnum{
		Type:  BuiltinWorkflowTypeOneOf,
		Value: b,
	}
}
