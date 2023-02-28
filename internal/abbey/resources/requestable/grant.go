package requestable

import (
	"encoding/json"
	"fmt"
)

const (
	grantTypeGenerate = "Generate"

	generateGrantTypeGithub = "Github"
)

type (
	Grant struct {
		value GrantEnum
	}

	GrantEnum interface {
		VisitGrant(GrantVisitor)
	}

	GrantVisitor struct {
		Generate func(GenerateGrant)
	}
)

var _ GrantEnum = (*GenerateGrant)(nil)

func (g GenerateGrant) VisitGrant(visitor GrantVisitor) {
	visitor.Generate(g)
}

func (g Grant) MarshalJSON() ([]byte, error) {
	var (
		err   error
		type_ string
		value json.RawMessage
	)

	g.value.VisitGrant(GrantVisitor{
		Generate: func(grant GenerateGrant) {
			type_ = grantTypeGenerate
			value, err = json.Marshal(grant)
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

func (g *Grant) UnmarshalJSON(b []byte) error {
	var (
		e     enum
		value GrantEnum
	)

	if err := json.Unmarshal(b, &e); err != nil {
		return err
	}

	switch e.Type {
	case grantTypeGenerate:
		var x GenerateGrant
		if err := json.Unmarshal(e.Value, &x); err != nil {
			return err
		}

		value = &x
	default:
		return fmt.Errorf("unknown grant type: %s", e.Type)
	}

	*g = Grant{value: value}

	return nil
}

type (
	GenerateGrant struct {
		value GenerateGrantEnum
	}

	GenerateGrantEnum interface {
		VisitGenerateGrant(GenerateGrantVisitor)
	}

	GenerateGrantVisitor struct {
		Github func(GithubGenerateDestination)
	}

	GithubGenerateDestination struct {
		Repo   string `json:"repo" tfsdk:"repo"`
		Path   string `json:"path" tfsdk:"path"`
		Append string `json:"append" tfsdk:"append"`
	}
)

var _ GenerateGrantEnum = (*GithubGenerateDestination)(nil)

func (g GithubGenerateDestination) VisitGenerateGrant(visitor GenerateGrantVisitor) {
	visitor.Github(g)
}

func (g GenerateGrant) MarshalJSON() ([]byte, error) {
	var (
		err   error
		type_ string
		value json.RawMessage
	)

	g.value.VisitGenerateGrant(GenerateGrantVisitor{
		Github: func(destination GithubGenerateDestination) {
			type_ = generateGrantTypeGithub
			value, err = json.Marshal(destination)
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

func (g *GenerateGrant) UnmarshalJSON(b []byte) error {
	var (
		e     enum
		value GenerateGrantEnum
	)

	if err := json.Unmarshal(b, &e); err != nil {
		return err
	}

	switch e.Type {
	case generateGrantTypeGithub:
		var x GithubGenerateDestination
		if err := json.Unmarshal(e.Value, &x); err != nil {
			return err
		}

		value = &x
	default:
		return fmt.Errorf("unknown generate type: %s", e.Type)
	}

	g.value = value

	return nil
}
