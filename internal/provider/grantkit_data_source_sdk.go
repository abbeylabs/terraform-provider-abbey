// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"abbey/v2/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"time"
)

func (r *GrantKitDataSourceModel) RefreshFromGetResponse(resp *shared.GrantKit) {
	r.CreatedAt = types.StringValue(resp.CreatedAt.Format(time.RFC3339))
	r.CurrentVersionID = types.StringValue(resp.CurrentVersionID)
	r.Description = types.StringValue(resp.Description)
	r.Grants = nil
	for _, grantsItem := range resp.Grants {
		var grants1 Grant
		grants1.CreatedAt = types.StringValue(grantsItem.CreatedAt.Format(time.RFC3339))
		grants1.Deleted = types.BoolValue(grantsItem.Deleted)
		grants1.GrantKitID = types.StringValue(grantsItem.GrantKitID)
		grants1.GrantKitVersionID = types.StringValue(grantsItem.GrantKitVersionID)
		grants1.ID = types.StringValue(grantsItem.ID)
		grants1.OrganizationID = types.StringValue(grantsItem.OrganizationID)
		grants1.RequestID = types.StringValue(grantsItem.RequestID)
		grants1.UpdatedAt = types.StringValue(grantsItem.UpdatedAt.Format(time.RFC3339))
		grants1.UserID = types.StringValue(grantsItem.UserID)
		r.Grants = append(r.Grants, grants1)
	}
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
	r.Policies = nil
	for _, policiesItem := range resp.Policies {
		var policies1 Policy
		if policiesItem.Bundle != nil {
			policies1.Bundle = types.StringValue(*policiesItem.Bundle)
		} else {
			policies1.Bundle = types.StringNull()
		}
		if policiesItem.Query != nil {
			policies1.Query = types.StringValue(*policiesItem.Query)
		} else {
			policies1.Query = types.StringNull()
		}
		r.Policies = append(r.Policies, policies1)
	}
	r.Requests = nil
	for _, requestsItem := range resp.Requests {
		var requests1 Request
		requests1.CreatedAt = types.StringValue(requestsItem.CreatedAt.Format(time.RFC3339))
		requests1.GrantID = types.StringValue(requestsItem.GrantID)
		requests1.GrantKitID = types.StringValue(requestsItem.GrantKitID)
		if requestsItem.GrantKitName != nil {
			requests1.GrantKitName = types.StringValue(*requestsItem.GrantKitName)
		} else {
			requests1.GrantKitName = types.StringNull()
		}
		requests1.GrantKitVersionID = types.StringValue(requestsItem.GrantKitVersionID)
		requests1.ID = types.StringValue(requestsItem.ID)
		requests1.PullRequest = types.StringValue(requestsItem.PullRequest)
		requests1.Reason = types.StringValue(requestsItem.Reason)
		requests1.Reviews = nil
		for _, reviewsItem := range requestsItem.Reviews {
			var reviews1 Review
			reviews1.CreatedAt = types.StringValue(reviewsItem.CreatedAt.Format(time.RFC3339))
			if reviews1.Grant == nil {
				reviews1.Grant = &Grant{}
			}
			if reviewsItem.Grant == nil {
				reviews1.Grant = nil
			} else {
				reviews1.Grant = &Grant{}
				reviews1.Grant.CreatedAt = types.StringValue(reviewsItem.Grant.CreatedAt.Format(time.RFC3339))
				reviews1.Grant.Deleted = types.BoolValue(reviewsItem.Grant.Deleted)
				reviews1.Grant.GrantKitID = types.StringValue(reviewsItem.Grant.GrantKitID)
				reviews1.Grant.GrantKitVersionID = types.StringValue(reviewsItem.Grant.GrantKitVersionID)
				reviews1.Grant.ID = types.StringValue(reviewsItem.Grant.ID)
				reviews1.Grant.OrganizationID = types.StringValue(reviewsItem.Grant.OrganizationID)
				reviews1.Grant.RequestID = types.StringValue(reviewsItem.Grant.RequestID)
				reviews1.Grant.UpdatedAt = types.StringValue(reviewsItem.Grant.UpdatedAt.Format(time.RFC3339))
				reviews1.Grant.UserID = types.StringValue(reviewsItem.Grant.UserID)
			}
			reviews1.GrantID = types.StringValue(reviewsItem.GrantID)
			reviews1.GrantKitName = types.StringValue(reviewsItem.GrantKitName)
			reviews1.GrantKitVersionID = types.StringValue(reviewsItem.GrantKitVersionID)
			reviews1.ID = types.StringValue(reviewsItem.ID)
			reviews1.PullRequest = types.StringValue(reviewsItem.PullRequest)
			reviews1.Reason = types.StringValue(reviewsItem.Reason)
			reviews1.RequestID = types.StringValue(reviewsItem.RequestID)
			reviews1.RequestReason = types.StringValue(reviewsItem.RequestReason)
			reviews1.Status = types.StringValue(string(reviewsItem.Status))
			reviews1.UpdatedAt = types.StringValue(reviewsItem.UpdatedAt.Format(time.RFC3339))
			if reviewsItem.UserEmail != nil {
				reviews1.UserEmail = types.StringValue(*reviewsItem.UserEmail)
			} else {
				reviews1.UserEmail = types.StringNull()
			}
			reviews1.UserID = types.StringValue(reviewsItem.UserID)
			requests1.Reviews = append(requests1.Reviews, reviews1)
		}
		requests1.Status = types.StringValue(string(requestsItem.Status))
		requests1.UpdatedAt = types.StringValue(requestsItem.UpdatedAt.Format(time.RFC3339))
		requests1.UserID = types.StringValue(requestsItem.UserID)
		r.Requests = append(r.Requests, requests1)
	}
	r.UpdatedAt = types.StringValue(resp.UpdatedAt.Format(time.RFC3339))
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
			if steps1.Reviewers == nil {
				steps1.Reviewers = &Reviewers{}
			}
			if stepsItem.Reviewers == nil {
				steps1.Reviewers = nil
			} else {
				steps1.Reviewers = &Reviewers{}
				steps1.Reviewers.AllOf = nil
				for _, v := range stepsItem.Reviewers.AllOf {
					steps1.Reviewers.AllOf = append(steps1.Reviewers.AllOf, types.StringValue(v))
				}
				steps1.Reviewers.OneOf = nil
				for _, v := range stepsItem.Reviewers.OneOf {
					steps1.Reviewers.OneOf = append(steps1.Reviewers.OneOf, types.StringValue(v))
				}
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
