// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"terraform/internal/sdk/pkg/models/shared"
	"time"
)

func (r *GrantKitResourceModel) ToCreateSDKType() *shared.GrantKitCreateParams {
	description := r.Description.ValueString()
	name := r.Name.ValueString()
	append1 := new(string)
	if !r.Output.Append.IsUnknown() && !r.Output.Append.IsNull() {
		*append1 = r.Output.Append.ValueString()
	} else {
		append1 = nil
	}
	location := r.Output.Location.ValueString()
	overwrite := new(string)
	if !r.Output.Overwrite.IsUnknown() && !r.Output.Overwrite.IsNull() {
		*overwrite = r.Output.Overwrite.ValueString()
	} else {
		overwrite = nil
	}
	output := shared.Output{
		Append:    append1,
		Location:  location,
		Overwrite: overwrite,
	}
	var policies *shared.Policies
	if r.Policies != nil {
		grantIf := make([]shared.Policy, 0)
		for _, grantIfItem := range r.Policies.GrantIf {
			bundle := new(string)
			if !grantIfItem.Bundle.IsUnknown() && !grantIfItem.Bundle.IsNull() {
				*bundle = grantIfItem.Bundle.ValueString()
			} else {
				bundle = nil
			}
			query := new(string)
			if !grantIfItem.Query.IsUnknown() && !grantIfItem.Query.IsNull() {
				*query = grantIfItem.Query.ValueString()
			} else {
				query = nil
			}
			grantIf = append(grantIf, shared.Policy{
				Bundle: bundle,
				Query:  query,
			})
		}
		revokeIf := make([]shared.Policy, 0)
		for _, revokeIfItem := range r.Policies.RevokeIf {
			bundle1 := new(string)
			if !revokeIfItem.Bundle.IsUnknown() && !revokeIfItem.Bundle.IsNull() {
				*bundle1 = revokeIfItem.Bundle.ValueString()
			} else {
				bundle1 = nil
			}
			query1 := new(string)
			if !revokeIfItem.Query.IsUnknown() && !revokeIfItem.Query.IsNull() {
				*query1 = revokeIfItem.Query.ValueString()
			} else {
				query1 = nil
			}
			revokeIf = append(revokeIf, shared.Policy{
				Bundle: bundle1,
				Query:  query1,
			})
		}
		policies = &shared.Policies{
			GrantIf:  grantIf,
			RevokeIf: revokeIf,
		}
	}
	var workflow *shared.GrantWorkflow
	if r.Workflow != nil {
		steps := make([]shared.Step, 0)
		for _, stepsItem := range r.Workflow.Steps {
			allOf := make([]string, 0)
			for _, allOfItem := range stepsItem.Reviewers.AllOf {
				allOf = append(allOf, allOfItem.ValueString())
			}
			oneOf := make([]string, 0)
			for _, oneOfItem := range stepsItem.Reviewers.OneOf {
				oneOf = append(oneOf, oneOfItem.ValueString())
			}
			reviewers := shared.Reviewers{
				AllOf: allOf,
				OneOf: oneOf,
			}
			skipIf := make([]shared.Policy, 0)
			for _, skipIfItem := range stepsItem.SkipIf {
				bundle2 := new(string)
				if !skipIfItem.Bundle.IsUnknown() && !skipIfItem.Bundle.IsNull() {
					*bundle2 = skipIfItem.Bundle.ValueString()
				} else {
					bundle2 = nil
				}
				query2 := new(string)
				if !skipIfItem.Query.IsUnknown() && !skipIfItem.Query.IsNull() {
					*query2 = skipIfItem.Query.ValueString()
				} else {
					query2 = nil
				}
				skipIf = append(skipIf, shared.Policy{
					Bundle: bundle2,
					Query:  query2,
				})
			}
			steps = append(steps, shared.Step{
				Reviewers: reviewers,
				SkipIf:    skipIf,
			})
		}
		workflow = &shared.GrantWorkflow{
			Steps: steps,
		}
	}
	out := shared.GrantKitCreateParams{
		Description: description,
		Name:        name,
		Output:      output,
		Policies:    policies,
		Workflow:    workflow,
	}
	return &out
}

func (r *GrantKitResourceModel) ToGetSDKType() *shared.GrantKitCreateParams {
	out := r.ToCreateSDKType()
	return out
}

func (r *GrantKitResourceModel) ToUpdateSDKType() *shared.GrantKitUpdateParams {
	description := r.Description.ValueString()
	name := r.Name.ValueString()
	append1 := new(string)
	if !r.Output.Append.IsUnknown() && !r.Output.Append.IsNull() {
		*append1 = r.Output.Append.ValueString()
	} else {
		append1 = nil
	}
	location := r.Output.Location.ValueString()
	overwrite := new(string)
	if !r.Output.Overwrite.IsUnknown() && !r.Output.Overwrite.IsNull() {
		*overwrite = r.Output.Overwrite.ValueString()
	} else {
		overwrite = nil
	}
	output := shared.Output{
		Append:    append1,
		Location:  location,
		Overwrite: overwrite,
	}
	var policies *shared.Policies
	if r.Policies != nil {
		grantIf := make([]shared.Policy, 0)
		for _, grantIfItem := range r.Policies.GrantIf {
			bundle := new(string)
			if !grantIfItem.Bundle.IsUnknown() && !grantIfItem.Bundle.IsNull() {
				*bundle = grantIfItem.Bundle.ValueString()
			} else {
				bundle = nil
			}
			query := new(string)
			if !grantIfItem.Query.IsUnknown() && !grantIfItem.Query.IsNull() {
				*query = grantIfItem.Query.ValueString()
			} else {
				query = nil
			}
			grantIf = append(grantIf, shared.Policy{
				Bundle: bundle,
				Query:  query,
			})
		}
		revokeIf := make([]shared.Policy, 0)
		for _, revokeIfItem := range r.Policies.RevokeIf {
			bundle1 := new(string)
			if !revokeIfItem.Bundle.IsUnknown() && !revokeIfItem.Bundle.IsNull() {
				*bundle1 = revokeIfItem.Bundle.ValueString()
			} else {
				bundle1 = nil
			}
			query1 := new(string)
			if !revokeIfItem.Query.IsUnknown() && !revokeIfItem.Query.IsNull() {
				*query1 = revokeIfItem.Query.ValueString()
			} else {
				query1 = nil
			}
			revokeIf = append(revokeIf, shared.Policy{
				Bundle: bundle1,
				Query:  query1,
			})
		}
		policies = &shared.Policies{
			GrantIf:  grantIf,
			RevokeIf: revokeIf,
		}
	}
	var workflow *shared.GrantWorkflow
	if r.Workflow != nil {
		steps := make([]shared.Step, 0)
		for _, stepsItem := range r.Workflow.Steps {
			allOf := make([]string, 0)
			for _, allOfItem := range stepsItem.Reviewers.AllOf {
				allOf = append(allOf, allOfItem.ValueString())
			}
			oneOf := make([]string, 0)
			for _, oneOfItem := range stepsItem.Reviewers.OneOf {
				oneOf = append(oneOf, oneOfItem.ValueString())
			}
			reviewers := shared.Reviewers{
				AllOf: allOf,
				OneOf: oneOf,
			}
			skipIf := make([]shared.Policy, 0)
			for _, skipIfItem := range stepsItem.SkipIf {
				bundle2 := new(string)
				if !skipIfItem.Bundle.IsUnknown() && !skipIfItem.Bundle.IsNull() {
					*bundle2 = skipIfItem.Bundle.ValueString()
				} else {
					bundle2 = nil
				}
				query2 := new(string)
				if !skipIfItem.Query.IsUnknown() && !skipIfItem.Query.IsNull() {
					*query2 = skipIfItem.Query.ValueString()
				} else {
					query2 = nil
				}
				skipIf = append(skipIf, shared.Policy{
					Bundle: bundle2,
					Query:  query2,
				})
			}
			steps = append(steps, shared.Step{
				Reviewers: reviewers,
				SkipIf:    skipIf,
			})
		}
		workflow = &shared.GrantWorkflow{
			Steps: steps,
		}
	}
	out := shared.GrantKitUpdateParams{
		Description: description,
		Name:        name,
		Output:      output,
		Policies:    policies,
		Workflow:    workflow,
	}
	return &out
}

func (r *GrantKitResourceModel) ToDeleteSDKType() *shared.GrantKitCreateParams {
	out := r.ToCreateSDKType()
	return out
}

func (r *GrantKitResourceModel) RefreshFromGetResponse(resp *shared.GrantKit) {
	r.CreatedAt = types.StringValue(resp.CreatedAt.Format(time.RFC3339))
	r.Description = types.StringValue(resp.Description)
	r.ID = types.StringValue(resp.ID)
	r.Name = types.StringValue(resp.Name)
	if resp.Output.Append != nil {
		r.Output.Append = types.StringValue(*resp.Output.Append)
	} else {
		r.Output.Append = types.StringNull()
	}
	r.Output.Location = types.StringValue(resp.Output.Location)
	if resp.Output.Overwrite != nil {
		r.Output.Overwrite = types.StringValue(*resp.Output.Overwrite)
	} else {
		r.Output.Overwrite = types.StringNull()
	}
	if r.Policies == nil {
		r.Policies = &Policies{}
	}
	if resp.Policies == nil {
		r.Policies = nil
	} else {
		r.Policies = &Policies{}
		r.Policies.GrantIf = nil
		for _, grantIfItem := range resp.Policies.GrantIf {
			var grantIf1 Policy
			if grantIfItem.Bundle != nil {
				grantIf1.Bundle = types.StringValue(*grantIfItem.Bundle)
			} else {
				grantIf1.Bundle = types.StringNull()
			}
			if grantIfItem.Query != nil {
				grantIf1.Query = types.StringValue(*grantIfItem.Query)
			} else {
				grantIf1.Query = types.StringNull()
			}
			r.Policies.GrantIf = append(r.Policies.GrantIf, grantIf1)
		}
		r.Policies.RevokeIf = nil
		for _, revokeIfItem := range resp.Policies.RevokeIf {
			var revokeIf1 Policy
			if revokeIfItem.Bundle != nil {
				revokeIf1.Bundle = types.StringValue(*revokeIfItem.Bundle)
			} else {
				revokeIf1.Bundle = types.StringNull()
			}
			if revokeIfItem.Query != nil {
				revokeIf1.Query = types.StringValue(*revokeIfItem.Query)
			} else {
				revokeIf1.Query = types.StringNull()
			}
			r.Policies.RevokeIf = append(r.Policies.RevokeIf, revokeIf1)
		}
	}
	r.UpdatedAt = types.StringValue(resp.UpdatedAt.Format(time.RFC3339))
	if resp.Version != nil {
		r.Version = types.Int64Value(*resp.Version)
	} else {
		r.Version = types.Int64Null()
	}
	if r.Workflow == nil {
		r.Workflow = &GrantWorkflow{}
	}
	if resp.Workflow == nil {
		r.Workflow = nil
	} else {
		r.Workflow = &GrantWorkflow{}
		r.Workflow.Steps = nil
		for _, stepsItem := range resp.Workflow.Steps {
			var steps1 Step
			steps1.Reviewers.AllOf = nil
			for _, v := range stepsItem.Reviewers.AllOf {
				steps1.Reviewers.AllOf = append(steps1.Reviewers.AllOf, types.StringValue(v))
			}
			steps1.Reviewers.OneOf = nil
			for _, v := range stepsItem.Reviewers.OneOf {
				steps1.Reviewers.OneOf = append(steps1.Reviewers.OneOf, types.StringValue(v))
			}
			steps1.SkipIf = nil
			for _, skipIfItem := range stepsItem.SkipIf {
				var skipIf1 Policy
				if skipIfItem.Bundle != nil {
					skipIf1.Bundle = types.StringValue(*skipIfItem.Bundle)
				} else {
					skipIf1.Bundle = types.StringNull()
				}
				if skipIfItem.Query != nil {
					skipIf1.Query = types.StringValue(*skipIfItem.Query)
				} else {
					skipIf1.Query = types.StringNull()
				}
				steps1.SkipIf = append(steps1.SkipIf, skipIf1)
			}
			r.Workflow.Steps = append(r.Workflow.Steps, steps1)
		}
	}
}

func (r *GrantKitResourceModel) RefreshFromCreateResponse(resp *shared.GrantKit) {
	r.RefreshFromGetResponse(resp)
}

func (r *GrantKitResourceModel) RefreshFromUpdateResponse(resp *shared.GrantKit) {
	r.RefreshFromGetResponse(resp)
}
