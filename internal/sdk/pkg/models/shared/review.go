// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"time"
)

type Review struct {
	CreatedAt         time.Time    `json:"created_at"`
	Grant             *Grant       `json:"grant,omitempty"`
	GrantID           string       `json:"grant_id"`
	GrantKitName      string       `json:"grant_kit_name"`
	GrantKitVersionID string       `json:"grant_kit_version_id"`
	ID                string       `json:"id"`
	PullRequest       string       `json:"pull_request"`
	Reason            string       `json:"reason"`
	RequestID         string       `json:"request_id"`
	RequestReason     string       `json:"request_reason"`
	Status            ReviewStatus `json:"status"`
	UpdatedAt         time.Time    `json:"updated_at"`
	UserEmail         *string      `json:"user_email,omitempty"`
	UserID            string       `json:"user_id"`
}
