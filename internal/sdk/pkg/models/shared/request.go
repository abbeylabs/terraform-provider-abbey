// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"time"
)

// Request - Created
type Request struct {
	CreatedAt time.Time     `json:"created_at"`
	ID        string        `json:"id"`
	Reason    *string       `json:"reason,omitempty"`
	Reviews   []Review      `json:"reviews,omitempty"`
	Status    RequestStatus `json:"status"`
}