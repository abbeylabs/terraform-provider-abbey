// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"abbey/v2/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"time"
)

func (r *IdentityResourceModel) ToCreateSDKType() *shared.IdentityParams {
	linked := r.Linked.ValueString()
	name := r.Name.ValueString()
	out := shared.IdentityParams{
		Linked: linked,
		Name:   name,
	}
	return &out
}

func (r *IdentityResourceModel) ToGetSDKType() *shared.IdentityParams {
	out := r.ToCreateSDKType()
	return out
}

func (r *IdentityResourceModel) ToDeleteSDKType() *shared.IdentityParams {
	out := r.ToCreateSDKType()
	return out
}

func (r *IdentityResourceModel) RefreshFromGetResponse(resp *shared.Identity) {
	r.CreatedAt = types.StringValue(resp.CreatedAt.Format(time.RFC3339))
	r.ID = types.StringValue(resp.ID)
	r.Linked = types.StringValue(resp.Linked)
	r.Name = types.StringValue(resp.Name)
}

func (r *IdentityResourceModel) RefreshFromCreateResponse(resp *shared.Identity) {
	r.RefreshFromGetResponse(resp)
}
