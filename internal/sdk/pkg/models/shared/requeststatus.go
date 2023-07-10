// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type RequestStatus string

const (
	RequestStatusPending  RequestStatus = "Pending"
	RequestStatusDenied   RequestStatus = "Denied"
	RequestStatusApproved RequestStatus = "Approved"
	RequestStatusCanceled RequestStatus = "Canceled"
)

func (e RequestStatus) ToPointer() *RequestStatus {
	return &e
}

func (e *RequestStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Pending":
		fallthrough
	case "Denied":
		fallthrough
	case "Approved":
		fallthrough
	case "Canceled":
		*e = RequestStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RequestStatus: %v", v)
	}
}
