// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"abbey/internal/sdk"
	"context"
	"fmt"

	"abbey/internal/sdk/pkg/models/operations"
	"abbey/internal/validators"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
) // Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &GrantKitResource{}
var _ resource.ResourceWithImportState = &GrantKitResource{}

func NewGrantKitResource() resource.Resource {
	return &GrantKitResource{}
}

// GrantKitResource defines the resource implementation.
type GrantKitResource struct {
	client *sdk.SDK
}

// GrantKitResourceModel describes the resource data model.
type GrantKitResourceModel struct {
	CreatedAt   types.String   `tfsdk:"created_at"`
	Description types.String   `tfsdk:"description"`
	ID          types.String   `tfsdk:"id"`
	Name        types.String   `tfsdk:"name"`
	Output      Output         `tfsdk:"output"`
	Policies    *Policies      `tfsdk:"policies"`
	UpdatedAt   types.String   `tfsdk:"updated_at"`
	Version     types.Int64    `tfsdk:"version"`
	Workflow    *GrantWorkflow `tfsdk:"workflow"`
}

func (r *GrantKitResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_grant_kit"
}

func (r *GrantKitResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "GrantKit Resource",

		Attributes: map[string]schema.Attribute{
			"created_at": schema.StringAttribute{
				Computed: true,
				Validators: []validator.String{
					validators.IsRFC3339(),
				},
			},
			"description": schema.StringAttribute{
				Required: true,
			},
			"id": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Required: true,
			},
			"output": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"append": schema.StringAttribute{
						Computed: true,
						Optional: true,
					},
					"location": schema.StringAttribute{
						Required: true,
					},
					"overwrite": schema.StringAttribute{
						Computed: true,
						Optional: true,
					},
				},
			},
			"policies": schema.SingleNestedAttribute{
				Computed: true,
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"grant_if": schema.ListNestedAttribute{
						Computed: true,
						Optional: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"bundle": schema.StringAttribute{
									Computed: true,
									Optional: true,
								},
								"query": schema.StringAttribute{
									Computed: true,
									Optional: true,
								},
							},
						},
					},
					"revoke_if": schema.ListNestedAttribute{
						Computed: true,
						Optional: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"bundle": schema.StringAttribute{
									Computed: true,
									Optional: true,
								},
								"query": schema.StringAttribute{
									Computed: true,
									Optional: true,
								},
							},
						},
					},
				},
			},
			"updated_at": schema.StringAttribute{
				Computed: true,
				Validators: []validator.String{
					validators.IsRFC3339(),
				},
			},
			"version": schema.Int64Attribute{
				Computed: true,
			},
			"workflow": schema.SingleNestedAttribute{
				Computed: true,
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"steps": schema.ListNestedAttribute{
						Required: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"reviewers": schema.SingleNestedAttribute{
									Required: true,
									Attributes: map[string]schema.Attribute{
										"all_of": schema.ListAttribute{
											Computed:    true,
											Optional:    true,
											ElementType: types.StringType,
										},
										"one_of": schema.ListAttribute{
											Computed:    true,
											Optional:    true,
											ElementType: types.StringType,
										},
									},
								},
								"skip_if": schema.ListNestedAttribute{
									Computed: true,
									Optional: true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"bundle": schema.StringAttribute{
												Computed: true,
												Optional: true,
											},
											"query": schema.StringAttribute{
												Computed: true,
												Optional: true,
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func (r *GrantKitResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*sdk.SDK)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *sdk.SDK, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *GrantKitResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *GrantKitResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	request := *data.ToCreateSDKType()
	res, err := r.client.GrantKits.CreateGrantKit(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode != 201 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if res.GrantKit == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromCreateResponse(res.GrantKit)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GrantKitResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *GrantKitResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	grantKitIDOrName := data.ID.ValueString()
	request := operations.GetGrantKitByIDRequest{
		GrantKitIDOrName: grantKitIDOrName,
	}
	res, err := r.client.GrantKits.GetGrantKitByID(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if res.GrantKit == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromGetResponse(res.GrantKit)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GrantKitResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *GrantKitResourceModel
	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	grantKitUpdateParams := *data.ToUpdateSDKType()
	grantKitIDOrName := data.ID.ValueString()
	request := operations.UpdateGrantKitRequest{
		GrantKitUpdateParams: grantKitUpdateParams,
		GrantKitIDOrName:     grantKitIDOrName,
	}
	res, err := r.client.GrantKits.UpdateGrantKit(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if res.GrantKit == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromUpdateResponse(res.GrantKit)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GrantKitResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *GrantKitResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	grantKitIDOrName := data.ID.ValueString()
	request := operations.DeleteGrantKitRequest{
		GrantKitIDOrName: grantKitIDOrName,
	}
	res, err := r.client.GrantKits.DeleteGrantKit(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}

}

func (r *GrantKitResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
