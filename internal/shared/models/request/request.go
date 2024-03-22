package request

import (
	"abbey/v2/internal/shared/models/review"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type Request struct {
	Id                types.String    `tfsdk:"id"`
	GrantKitId        types.String    `tfsdk:"grant_kit_id"`
	GrantKitVersionId types.String    `tfsdk:"grant_kit_version_id"`
	GrantKitName      types.String    `tfsdk:"grant_kit_name"`
	Reason            types.String    `tfsdk:"reason"`
	UserId            types.String    `tfsdk:"user_id"`
	Status            types.String    `tfsdk:"status"`
	Reviews           []review.Review `tfsdk:"reviews"`
	GrantId           types.String    `tfsdk:"grant_id"`
	CreatedAt         types.String    `tfsdk:"created_at"`
	UpdatedAt         types.String    `tfsdk:"updated_at"`
	PullRequest       types.String    `tfsdk:"pull_request"`
}
