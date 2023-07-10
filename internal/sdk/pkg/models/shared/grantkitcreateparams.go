// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type GrantKitCreateParams struct {
	Description string         `json:"description"`
	Name        string         `json:"name"`
	Output      Output         `json:"output"`
	Policies    []Policy       `json:"policies,omitempty"`
	Workflow    *GrantWorkflow `json:"workflow,omitempty"`
}
