// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type ReviewStep struct {
	Reviewers ReviewerQuantifier `json:"reviewers"`
	SkipIf    []Policy           `json:"skip_if,omitempty"`
}
