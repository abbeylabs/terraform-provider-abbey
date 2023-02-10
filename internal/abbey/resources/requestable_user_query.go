package resources

import (
	"encoding/json"
	"fmt"
)

type UserQueryType string

func (u UserQueryType) String() string {
	return string(u)
}

const UserQueryTypeAuthId UserQueryType = "AuthId"

type UserQueryEnum struct {
	Type  UserQueryType `json:"type"`
	Value UserQuery     `json:"value"`
}

var _ UserQuery = (*UserQueryEnum)(nil)

func (u *UserQueryEnum) VisitUserQuery(visitor UserQueryVisitor) {
	u.Value.VisitUserQuery(visitor)
}

func (u *UserQueryEnum) UnmarshalJSON(b []byte) error {
	var e enum
	if err := json.Unmarshal(b, &e); err != nil {
		return err
	}

	switch e.Type {
	case "":
	case UserQueryTypeAuthId.String():
		var x UserQueryAuthId
		if err := json.Unmarshal(e.Value, &x); err != nil {
			return err
		}

		u.Type = UserQueryTypeAuthId
		u.Value = &x
	default:
		return fmt.Errorf("unknown user query type: %s", e.Type)
	}

	return nil
}

type UserQuery interface {
	VisitUserQuery(UserQueryVisitor)
}

type UserQueryVisitor struct {
	VisitAuthId func(*UserQueryAuthId)
}

type UserQueryAuthId string

func (u *UserQueryAuthId) String() string {
	return string(*u)
}

var _ UserQuery = (*UserQueryAuthId)(nil)

func (u *UserQueryAuthId) VisitUserQuery(visitor UserQueryVisitor) {
	visitor.VisitAuthId(u)
}

func (u *UserQueryAuthId) ToEnum() UserQueryEnum {
	return UserQueryEnum{
		Type:  UserQueryTypeAuthId,
		Value: u,
	}
}
