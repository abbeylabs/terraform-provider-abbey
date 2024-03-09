package identity

import (
	"abbey/v2/internal/utils"
	"context"
	"fmt"
	"github.com/go-provider-sdk/pkg/client"
	"github.com/go-provider-sdk/pkg/identities"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ensure we implement the needed interfaces
var _ resource.Resource = &IdentityResource{}
var _ resource.ResourceWithImportState = &IdentityResource{}

// constructor
func NewIdentityResource() resource.Resource {
	return &IdentityResource{}
}

// client wrapper
type IdentityResource struct {
	client *client.Client
}

type IdentityResourceModel struct {
	Id           types.String `tfsdk:"id"`
	CreatedAt    types.String `tfsdk:"created_at"`
	UpdatedAt    types.String `tfsdk:"updated_at"`
	AbbeyAccount types.String `tfsdk:"abbey_account"`
	Source       types.String `tfsdk:"source"`
	Metadata     types.String `tfsdk:"metadata"`
}

func (r *IdentityResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_identity"
}

func (r *IdentityResource) Schema(_ context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "id",
				Computed:    true,
				Optional:    true,
			},

			"created_at": schema.StringAttribute{
				Description: "created_at",
				Computed:    true,
				Optional:    true,
			},

			"updated_at": schema.StringAttribute{
				Description: "updated_at",
				Computed:    true,
				Optional:    true,
			},

			"abbey_account": schema.StringAttribute{
				Description: "abbey_account",
				Required:    true,
			},

			"source": schema.StringAttribute{
				Description: "source",
				Required:    true,
			},

			"metadata": schema.StringAttribute{
				Description: "Json encoded string. See documentation for details.",
				Required:    true,
			},
		},
	}
}

func (r *IdentityResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	apiClient, ok := req.ProviderData.(*client.Client)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *client.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = apiClient
}

func (r *IdentityResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var dataModel IdentityResourceModel
	utils.PopulateModelData(ctx, &dataModel, resp.Diagnostics, req.State.Get)

	if resp.Diagnostics.HasError() {
		return
	}

	IdentityId := dataModel.Id.ValueString()

	clientResponse, err := r.client.Identities.GetIdentity(ctx, IdentityId)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unexpected error calling Identities.GetIdentity",
			err.Error(),
		)

		return
	}

	identity := clientResponse.Data

	dataModel.Id = utils.NullableString(identity.GetId())

	dataModel.CreatedAt = utils.NullableString(identity.GetCreatedAt())

	dataModel.UpdatedAt = utils.NullableString(identity.GetUpdatedAt())

	dataModel.AbbeyAccount = utils.NullableString(identity.GetAbbeyAccount())

	dataModel.Source = utils.NullableString(identity.GetSource())

	dataModel.Metadata = utils.NullableString(identity.GetMetadata())

	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &dataModel)...)
}

func (r *IdentityResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var dataModel IdentityResourceModel
	utils.PopulateModelData(ctx, &dataModel, resp.Diagnostics, req.Plan.Get)

	if resp.Diagnostics.HasError() {
		return
	}

	requestBody := identities.IdentityParams{
		AbbeyAccount: dataModel.AbbeyAccount.ValueStringPointer(),
		Source:       dataModel.Source.ValueStringPointer(),
		Metadata:     dataModel.Metadata.ValueStringPointer(),
	}

	clientResponse, err := r.client.Identities.CreateIdentity(ctx, requestBody)

	if err != nil {
		resp.Diagnostics.AddError(
			"Error Creating Identity",
			err.Error(),
		)

		return
	}

	identity := clientResponse.Data
	dataModel.Id = utils.NullableString(identity.GetId())

	dataModel.CreatedAt = utils.NullableString(identity.GetCreatedAt())

	dataModel.UpdatedAt = utils.NullableString(identity.GetUpdatedAt())

	dataModel.AbbeyAccount = utils.NullableString(identity.GetAbbeyAccount())

	dataModel.Source = utils.NullableString(identity.GetSource())

	dataModel.Metadata = utils.NullableString(identity.GetMetadata())

	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &dataModel)...)
}

func (r *IdentityResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var dataModel = &IdentityResourceModel{}
	utils.PopulateModelData(ctx, &dataModel, resp.Diagnostics, req.State.Get)

	if resp.Diagnostics.HasError() {
		return
	}

	IdentityId := dataModel.Id.ValueString()

	_, err := r.client.Identities.DeleteIdentity(ctx, IdentityId)

	if err != nil {
		resp.Diagnostics.AddError(
			"Error Deleting Identity",
			err.Error(),
		)
	}
}

func (r *IdentityResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var stateModel = &IdentityResourceModel{}
	var dataModel = &IdentityResourceModel{}
	utils.PopulateModelData(ctx, &stateModel, resp.Diagnostics, req.State.Get)
	utils.PopulateModelData(ctx, &dataModel, resp.Diagnostics, req.Plan.Get)

	if resp.Diagnostics.HasError() {
		return
	}

	IdentityId := stateModel.Id.ValueString()

	requestBody := identities.IdentityParams{
		AbbeyAccount: dataModel.AbbeyAccount.ValueStringPointer(),
		Source:       dataModel.Source.ValueStringPointer(),
		Metadata:     dataModel.Metadata.ValueStringPointer(),
	}

	clientResponse, err := r.client.Identities.UpdateIdentity(ctx, IdentityId, requestBody)

	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating Identity",
			err.Error(),
		)

		return
	}
	identity := clientResponse.Data
	dataModel.Id = utils.NullableString(identity.GetId())

	dataModel.CreatedAt = utils.NullableString(identity.GetCreatedAt())

	dataModel.UpdatedAt = utils.NullableString(identity.GetUpdatedAt())

	dataModel.AbbeyAccount = utils.NullableString(identity.GetAbbeyAccount())

	dataModel.Source = utils.NullableString(identity.GetSource())

	dataModel.Metadata = utils.NullableString(identity.GetMetadata())

	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &dataModel)...)
}

func (r *IdentityResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
