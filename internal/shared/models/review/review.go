package review

import (
	"abbey/v2/internal/shared/models/grant"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type Review struct {
	Id                types.String `tfsdk:"id"`
	UserId            types.String `tfsdk:"user_id"`
	UserEmail         types.String `tfsdk:"user_email"`
	RequestId         types.String `tfsdk:"request_id"`
	Status            types.String `tfsdk:"status"`
	RequestReason     types.String `tfsdk:"request_reason"`
	Reason            types.String `tfsdk:"reason"`
	GrantKitVersionId types.String `tfsdk:"grant_kit_version_id"`
	GrantKitName      types.String `tfsdk:"grant_kit_name"`
	GrantId           types.String `tfsdk:"grant_id"`
	Grant             *grant.Grant `tfsdk:"grant"`
	CreatedAt         types.String `tfsdk:"created_at"`
	UpdatedAt         types.String `tfsdk:"updated_at"`
	PullRequest       types.String `tfsdk:"pull_request"`
}
