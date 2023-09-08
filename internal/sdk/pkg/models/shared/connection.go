// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"time"
)

type Connection struct {
	CreatedAt time.Time      `json:"created_at"`
	ID        string         `json:"id"`
	Name      *string        `json:"name,omitempty"`
	Type      ConnectionType `json:"type"`
}
