// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"time"
)

// Requestable - Created
type Requestable struct {
	CreatedAt time.Time `json:"created_at"`
	Grant     Grant     `json:"grant"`
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Workflow  *Workflow `json:"workflow,omitempty"`
}