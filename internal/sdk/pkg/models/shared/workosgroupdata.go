// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"time"
)

type WorkOSGroupDataPreviousAttributes struct {
}

type WorkOSGroupDataRawAttributes struct {
}

type WorkOSGroupData struct {
	CreatedAt          *time.Time                         `json:"created_at,omitempty"`
	DirectoryID        *string                            `json:"directory_id,omitempty"`
	Group              *WorkOSGroup                       `json:"group,omitempty"`
	ID                 *string                            `json:"id,omitempty"`
	IdpID              *string                            `json:"idp_id,omitempty"`
	Name               *string                            `json:"name,omitempty"`
	Object             *string                            `json:"object,omitempty"`
	OrganizationID     *string                            `json:"organization_id,omitempty"`
	PreviousAttributes *WorkOSGroupDataPreviousAttributes `json:"previous_attributes,omitempty"`
	RawAttributes      *WorkOSGroupDataRawAttributes      `json:"raw_attributes,omitempty"`
	UpdatedAt          *time.Time                         `json:"updated_at,omitempty"`
	User               *WorkOSUserData                    `json:"user,omitempty"`
}