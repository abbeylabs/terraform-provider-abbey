package requestable

import (
	"encoding/json"
	"fmt"
)

const userQueryTypeAuthId = "AuthId"

type (
	UserQuery struct {
		value UserQueryEnum
	}

	UserQueryEnum interface {
		VisitUserQuery(UserQueryVisitor)
	}

	UserQueryVisitor struct {
		AuthId func(AuthId)
	}

	AuthId struct {
		value string
	}
)

var _ UserQueryEnum = (*AuthId)(nil)

func (u UserQuery) VisitUserQuery(visitor UserQueryVisitor) {
	u.value.VisitUserQuery(visitor)
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

	*u = UserQuery{value: value}

	return nil
}

func (a AuthId) Value() string {
	return a.value
}

func (a AuthId) VisitUserQuery(visitor UserQueryVisitor) {
	visitor.AuthId(a)
}

func (a AuthId) MarshalJSON() ([]byte, error) {
	return json.Marshal(a.value)
}

func (a *AuthId) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	*a = AuthId{value: s}

	return nil
}
