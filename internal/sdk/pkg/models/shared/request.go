// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"time"
)

type Request struct {
	CreatedAt         time.Time     `json:"created_at"`
	GrantID           string        `json:"grant_id"`
	GrantKitID        string        `json:"grant_kit_id"`
	GrantKitName      *string       `json:"grant_kit_name,omitempty"`
	GrantKitVersionID string        `json:"grant_kit_version_id"`
	ID                string        `json:"id"`
	PullRequest       string        `json:"pull_request"`
	Reason            string        `json:"reason"`
	Reviews           []Review      `json:"reviews,omitempty"`
	Status            RequestStatus `json:"status"`
	UpdatedAt         time.Time     `json:"updated_at"`
	UserID            string        `json:"user_id"`
}
