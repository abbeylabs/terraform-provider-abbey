package demo

import (
	"abbey/v2/internal/utils"
	"context"
	"fmt"
	"github.com/go-provider-sdk/pkg/client"
	"github.com/go-provider-sdk/pkg/demo"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ensure we implement the needed interfaces
var _ resource.Resource = &DemoResource{}
var _ resource.ResourceWithImportState = &DemoResource{}

// constructor
func NewDemoResource() resource.Resource {
	return &DemoResource{}
}

// client wrapper
type DemoResource struct {
	client *client.Client
}

type DemoResourceModel struct {
	Email      types.String `tfsdk:"email"`
	Id         types.Int64  `tfsdk:"id"`
	UserId     types.String `tfsdk:"user_id"`
	CreatedAt  types.String `tfsdk:"created_at"`
	UpdatedAt  types.String `tfsdk:"updated_at"`
	Permission types.String `tfsdk:"permission"`
}

func (r *DemoResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_demo"
}

func (r *DemoResource) Schema(_ context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{

		Attributes: map[string]schema.Attribute{
			"email": schema.StringAttribute{
				Description: "The email of the user",
				Required:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},

			"id": schema.Int64Attribute{
				Description: "id",
				Computed:    true,
				Optional:    true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
			},

			"user_id": schema.StringAttribute{
				Description: "user_id",
				Computed:    true,
				Optional:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},

			"created_at": schema.StringAttribute{
				Description: "created_at",
				Computed:    true,
				Optional:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},

			"updated_at": schema.StringAttribute{
				Description: "updated_at",
				Computed:    true,
				Optional:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},

			"permission": schema.StringAttribute{
				Description: "permission",
				Required:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
		},
	}
}

func (r *DemoResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *DemoResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var dataModel DemoResourceModel
	utils.PopulateModelData(ctx, &dataModel, resp.Diagnostics, req.State.Get)

	if resp.Diagnostics.HasError() {
		return
	}

	params := demo.GetDemoRequestParams{}
	params.SetEmail(dataModel.Email.ValueString())

	if resp.Diagnostics.HasError() {
		return
	}

	clientResponse, err := r.client.Demo.GetDemo(ctx, params)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unexpected error calling Demo.GetDemo",
			err.Error(),
		)

		return
	}

	demo := clientResponse.Data

	dataModel.Id = utils.NullableInt64(demo.GetId())

	dataModel.UserId = utils.NullableString(demo.GetUserId())

	dataModel.CreatedAt = utils.NullableString(demo.GetCreatedAt())

	dataModel.UpdatedAt = utils.NullableString(demo.GetUpdatedAt())

	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &dataModel)...)
}

func (r *DemoResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var dataModel DemoResourceModel
	utils.PopulateModelData(ctx, &dataModel, resp.Diagnostics, req.Plan.Get)

	if resp.Diagnostics.HasError() {
		return
	}

	requestBody := demo.DemoParams{
		Permission: utils.Pointer(demo.Permission(dataModel.Permission.ValueString())),
		Email:      dataModel.Email.ValueStringPointer(),
	}

	clientResponse, err := r.client.Demo.CreateDemo(ctx, requestBody)

	if err != nil {
		resp.Diagnostics.AddError(
			"Error Creating Demo",
			err.Error(),
		)

		return
	}

	demo := clientResponse.Data
	dataModel.Id = utils.NullableInt64(demo.GetId())

	dataModel.UserId = utils.NullableString(demo.GetUserId())

	dataModel.CreatedAt = utils.NullableString(demo.GetCreatedAt())

	dataModel.UpdatedAt = utils.NullableString(demo.GetUpdatedAt())

	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &dataModel)...)
}

func (r *DemoResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var dataModel = &DemoResourceModel{}
	utils.PopulateModelData(ctx, &dataModel, resp.Diagnostics, req.State.Get)

	if resp.Diagnostics.HasError() {
		return
	}

	requestBody := demo.DemoParams{
		Permission: utils.Pointer(demo.Permission(dataModel.Permission.ValueString())),
		Email:      dataModel.Email.ValueStringPointer(),
	}

	_, err := r.client.Demo.DeleteDemo(ctx, requestBody)

	if err != nil {
		resp.Diagnostics.AddError(
			"Error Deleting Demo",
			err.Error(),
		)
	}
}

func (r *DemoResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
}

func (r *DemoResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
