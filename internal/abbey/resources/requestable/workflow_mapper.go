package requestable

import (
	"errors"
	"fmt"

	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

const (
	workflowTypeBuiltinTf = "builtin"

	builtinWorkflowTypeAllOfTf = "all_of"
	builtinWorkflowTypeOneOfTf = "one_of"
)

func BuiltinWorkflowFromTfTypesValue(value tftypes.Value) (ret *BuiltinWorkflow, err error) {
	var m map[string]tftypes.Value
	if err := value.As(&m); err != nil {
		return nil, err
	}

	var inner BuiltinWorkflowEnum

	for key, val := range m {
		switch key {
		case builtinWorkflowTypeAllOfTf:
			var inner_ *BuiltinWorkflowAllOf
			inner_, err = AllOfFromGoValue(val)
			if err != nil {
				return nil, err
			}
			if inner_ == nil {
				continue
			}

			inner = inner_
		case builtinWorkflowTypeOneOfTf:
			var inner_ *BuiltinWorkflowOneOf
			inner_, err = OneOfFromGoValue(val)
			if inner_ == nil {
				continue
			}
			if inner_ == nil {
				continue
			}

			inner = inner_
		default:
			return nil, fmt.Errorf("unknown key: %s", key)
		}
		if err != nil {
			return nil, err
		}
	}

	return &BuiltinWorkflow{value: inner}, nil
}

func AllOfFromGoValue(value tftypes.Value) (*BuiltinWorkflowAllOf, error) {
	var m map[string]tftypes.Value
	if err := value.As(&m); err != nil {
		return nil, err
	}

	if len(m) == 0 {
		return nil, nil
	}

	reviewersValue, ok := m["reviewers"]
	if !ok {
		return nil, errors.New("missing reviewers field")
	}

	reviewers, err := UserQueriesFromGoValue(reviewersValue)
	if err != nil {
		return nil, err
	}

	return &BuiltinWorkflowAllOf{Reviewers: reviewers}, nil
}

func OneOfFromGoValue(value tftypes.Value) (*BuiltinWorkflowOneOf, error) {
	var m map[string]tftypes.Value
	if err := value.As(&m); err != nil {
		return nil, err
	}

	if len(m) == 0 {
		return nil, nil
	}

	reviewersValue, ok := m["reviewers"]
	if !ok {
		return nil, errors.New("missing reviewers field")
	}

	reviewers, err := UserQueriesFromGoValue(reviewersValue)
	if err != nil {
		return nil, err
	}

	return &BuiltinWorkflowOneOf{Reviewers: reviewers}, nil
}

func UserQueriesFromGoValue(value tftypes.Value) ([]UserQuery, error) {
	var list []tftypes.Value
	if err := value.As(&list); err != nil {
		return nil, err
	}

	userQueries := make([]UserQuery, 0, len(list))

	for _, v := range list {
		var m map[string]tftypes.Value
		if err := v.As(&m); err != nil {
			return nil, err
		}

		if m == nil {
			return nil, errors.New("got nil user query")
		}

		var inner UserQueryEnum

		for key, val := range m {
			switch key {
			case userQueryTypeAuthIdTf:
				var s string
				if err := val.As(&s); err != nil {
					return nil, err
				}

				inner = &AuthId{value: s}
			default:
				return nil, fmt.Errorf("unknowon key: %s", key)
			}
		}

		userQueries = append(userQueries, UserQuery{value: inner})
	}

	return userQueries, nil
}
