// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"time"
)

// Identity - Created
type Identity struct {
	CreatedAt time.Time                `json:"created_at"`
	ID        string                   `json:"id"`
	Linked    map[string][]interface{} `json:"linked"`
	Name      string                   `json:"name"`
}
