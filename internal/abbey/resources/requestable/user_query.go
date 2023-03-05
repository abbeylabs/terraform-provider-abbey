package requestable

import (
	"encoding/json"
	"fmt"
)

const userQueryTypeAuthId = "AuthId"

type (
	UserQuery struct {
		Value UserQueryEnum
	}

	UserQueryEnum interface {
		VisitUserQuery(UserQueryVisitor)
	}

	UserQueryVisitor struct {
		AuthId func(AuthId)
	}

	AuthId struct {
		Value string
	}
)

var _ UserQueryEnum = (*AuthId)(nil)

func (u UserQuery) VisitUserQuery(visitor UserQueryVisitor) {
	u.Value.VisitUserQuery(visitor)
}

func (u UserQuery) MarshalJSON() ([]byte, error) {
	var (
		err   error
		type_ string
		value json.RawMessage
	)

	u.Value.VisitUserQuery(UserQueryVisitor{
		AuthId: func(authId AuthId) {
			type_ = userQueryTypeAuthId
			value, err = json.Marshal(authId)
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

func (u *UserQuery) UnmarshalJSON(b []byte) error {
	var (
		e     enum
		value UserQueryEnum
	)

	if err := json.Unmarshal(b, &e); err != nil {
		return err
	}

	switch e.Type {
	case userQueryTypeAuthId:
		var x AuthId
		if err := json.Unmarshal(e.Value, &x); err != nil {
			return err
		}

		value = &x
	default:
		return fmt.Errorf("unknown user query type: %s", e.Type)
	}

	*u = UserQuery{Value: value}

	return nil
}

func (a AuthId) VisitUserQuery(visitor UserQueryVisitor) {
	visitor.AuthId(a)
}

func (a AuthId) MarshalJSON() ([]byte, error) {
	return json.Marshal(a.Value)
}

func (a *AuthId) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	*a = AuthId{Value: s}

	return nil
}
