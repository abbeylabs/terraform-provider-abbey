package grant

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type Grant struct {
	Id                types.String `tfsdk:"id"`
	GrantKitId        types.String `tfsdk:"grant_kit_id"`
	GrantKitVersionId types.String `tfsdk:"grant_kit_version_id"`
	UserId            types.String `tfsdk:"user_id"`
	RequestId         types.String `tfsdk:"request_id"`
	OrganizationId    types.String `tfsdk:"organization_id"`
	Deleted           types.Bool   `tfsdk:"deleted"`
	CreatedAt         types.String `tfsdk:"created_at"`
	UpdatedAt         types.String `tfsdk:"updated_at"`
}
