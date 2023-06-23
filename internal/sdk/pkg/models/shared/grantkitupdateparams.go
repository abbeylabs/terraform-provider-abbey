// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type GrantKitUpdateParams struct {
	Description string `json:"description"`
	// The name of the connection
	Name     string         `json:"name"`
	Output   Output         `json:"output"`
	Policies *Policies      `json:"policies,omitempty"`
	Workflow *GrantWorkflow `json:"workflow,omitempty"`
}
