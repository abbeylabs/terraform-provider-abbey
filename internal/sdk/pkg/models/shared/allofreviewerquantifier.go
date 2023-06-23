// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type AllOfReviewerQuantifierType string

const (
	AllOfReviewerQuantifierTypeAllOf AllOfReviewerQuantifierType = "AllOf"
)

func (e AllOfReviewerQuantifierType) ToPointer() *AllOfReviewerQuantifierType {
	return &e
}

func (e *AllOfReviewerQuantifierType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "AllOf":
		*e = AllOfReviewerQuantifierType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AllOfReviewerQuantifierType: %v", v)
	}
}

type AllOfReviewerQuantifier struct {
	Reviewers []UserQuery                 `json:"reviewers,omitempty"`
	Type      AllOfReviewerQuantifierType `json:"type"`
}
