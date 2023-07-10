// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"time"
)

// GrantKit - Created
type GrantKit struct {
	CreatedAt        time.Time      `json:"created_at"`
	CurrentVersionID string         `json:"current_version_id"`
	Description      string         `json:"description"`
	ID               string         `json:"id"`
	Name             string         `json:"name"`
	Output           Output         `json:"output"`
	Policies         []Policy       `json:"policies,omitempty"`
	UpdatedAt        time.Time      `json:"updated_at"`
	Workflow         *GrantWorkflow `json:"workflow,omitempty"`
}
