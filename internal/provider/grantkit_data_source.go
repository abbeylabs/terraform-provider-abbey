// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"abbey/v2/internal/sdk"
	"abbey/v2/internal/sdk/pkg/models/operations"
	"context"
	"fmt"

	"abbey/v2/internal/validators"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &GrantKitDataSource{}
var _ datasource.DataSourceWithConfigure = &GrantKitDataSource{}

func NewGrantKitDataSource() datasource.DataSource {
	return &GrantKitDataSource{}
}

// GrantKitDataSource is the data source implementation.
type GrantKitDataSource struct {
	client *sdk.SDK
}

// GrantKitDataSourceModel describes the data model.
type GrantKitDataSourceModel struct {
	CreatedAt        types.String   `tfsdk:"created_at"`
	CurrentVersionID types.String   `tfsdk:"current_version_id"`
	Description      types.String   `tfsdk:"description"`
	Grants           []Grant        `tfsdk:"grants"`
	ID               types.String   `tfsdk:"id"`
	Name             types.String   `tfsdk:"name"`
	Output           Output         `tfsdk:"output"`
	Policies         []Policy       `tfsdk:"policies"`
	Requests         []Request      `tfsdk:"requests"`
	UpdatedAt        types.String   `tfsdk:"updated_at"`
	Workflow         *GrantWorkflow `tfsdk:"workflow"`
}

// Metadata returns the data source type name.
func (r *GrantKitDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_grant_kit"
}

// Schema defines the schema for the data source.
func (r *GrantKitDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "GrantKit DataSource",

		Attributes: map[string]schema.Attribute{
			"created_at": schema.StringAttribute{
				Computed: true,
				Validators: []validator.String{
					validators.IsRFC3339(),
				},
			},
			"current_version_id": schema.StringAttribute{
				Computed: true,
			},
			"description": schema.StringAttribute{
				Computed: true,
			},
			"grants": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"created_at": schema.StringAttribute{
							Computed: true,
							Validators: []validator.String{
								validators.IsRFC3339(),
							},
						},
						"deleted": schema.BoolAttribute{
							Computed: true,
						},
						"grant_kit_id": schema.StringAttribute{
							Computed: true,
						},
						"grant_kit_version_id": schema.StringAttribute{
							Computed: true,
						},
						"id": schema.StringAttribute{
							Computed: true,
						},
						"organization_id": schema.StringAttribute{
							Computed: true,
						},
						"request_id": schema.StringAttribute{
							Computed: true,
						},
						"updated_at": schema.StringAttribute{
							Computed: true,
							Validators: []validator.String{
								validators.IsRFC3339(),
							},
						},
						"user_id": schema.StringAttribute{
							Computed: true,
						},
					},
				},
			},
			"id": schema.StringAttribute{
				Required:    true,
				Description: `The ID of the grant kit or resource to retrieve.`,
			},
			"name": schema.StringAttribute{
				Computed: true,
			},
			"output": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"append": schema.StringAttribute{
						Computed: true,
					},
					"location": schema.StringAttribute{
						Computed: true,
					},
					"overwrite": schema.StringAttribute{
						Computed: true,
					},
				},
			},
			"policies": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"bundle": schema.StringAttribute{
							Computed: true,
						},
						"query": schema.StringAttribute{
							Computed: true,
						},
					},
				},
			},
			"requests": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"created_at": schema.StringAttribute{
							Computed: true,
							Validators: []validator.String{
								validators.IsRFC3339(),
							},
						},
						"grant_id": schema.StringAttribute{
							Computed: true,
						},
						"grant_kit_id": schema.StringAttribute{
							Computed: true,
						},
						"grant_kit_name": schema.StringAttribute{
							Computed: true,
						},
						"grant_kit_version_id": schema.StringAttribute{
							Computed: true,
						},
						"id": schema.StringAttribute{
							Computed: true,
						},
						"pull_request": schema.StringAttribute{
							Computed: true,
						},
						"reason": schema.StringAttribute{
							Computed: true,
						},
						"reviews": schema.ListNestedAttribute{
							Computed: true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"created_at": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											validators.IsRFC3339(),
										},
									},
									"grant": schema.SingleNestedAttribute{
										Computed: true,
										Attributes: map[string]schema.Attribute{
											"created_at": schema.StringAttribute{
												Computed: true,
												Validators: []validator.String{
													validators.IsRFC3339(),
												},
											},
											"deleted": schema.BoolAttribute{
												Computed: true,
											},
											"grant_kit_id": schema.StringAttribute{
												Computed: true,
											},
											"grant_kit_version_id": schema.StringAttribute{
												Computed: true,
											},
											"id": schema.StringAttribute{
												Computed: true,
											},
											"organization_id": schema.StringAttribute{
												Computed: true,
											},
											"request_id": schema.StringAttribute{
												Computed: true,
											},
											"updated_at": schema.StringAttribute{
												Computed: true,
												Validators: []validator.String{
													validators.IsRFC3339(),
												},
											},
											"user_id": schema.StringAttribute{
												Computed: true,
											},
										},
										Description: `Success`,
									},
									"grant_id": schema.StringAttribute{
										Computed: true,
									},
									"grant_kit_name": schema.StringAttribute{
										Computed: true,
									},
									"grant_kit_version_id": schema.StringAttribute{
										Computed: true,
									},
									"id": schema.StringAttribute{
										Computed: true,
									},
									"pull_request": schema.StringAttribute{
										Computed: true,
									},
									"reason": schema.StringAttribute{
										Computed: true,
									},
									"request_id": schema.StringAttribute{
										Computed: true,
									},
									"status": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"Pending",
												"Denied",
												"Approved",
												"Canceled",
											),
										},
										Description: `must be one of [Pending, Denied, Approved, Canceled]`,
									},
									"updated_at": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											validators.IsRFC3339(),
										},
									},
									"user_email": schema.StringAttribute{
										Computed: true,
									},
									"user_id": schema.StringAttribute{
										Computed: true,
									},
								},
							},
						},
						"status": schema.StringAttribute{
							Computed: true,
							Validators: []validator.String{
								stringvalidator.OneOf(
									"Pending",
									"Denied",
									"Approved",
									"Canceled",
								),
							},
							Description: `must be one of [Pending, Denied, Approved, Canceled]`,
						},
						"updated_at": schema.StringAttribute{
							Computed: true,
							Validators: []validator.String{
								validators.IsRFC3339(),
							},
						},
						"user_id": schema.StringAttribute{
							Computed: true,
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
			"workflow": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"steps": schema.ListNestedAttribute{
						Computed: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"reviewers": schema.SingleNestedAttribute{
									Computed: true,
									Attributes: map[string]schema.Attribute{
										"all_of": schema.ListAttribute{
											Computed:    true,
											ElementType: types.StringType,
										},
										"one_of": schema.ListAttribute{
											Computed:    true,
											ElementType: types.StringType,
										},
									},
								},
								"skip_if": schema.ListNestedAttribute{
									Computed: true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"bundle": schema.StringAttribute{
												Computed: true,
											},
											"query": schema.StringAttribute{
												Computed: true,
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

func (r *GrantKitDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*sdk.SDK)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected DataSource Configure Type",
			fmt.Sprintf("Expected *sdk.SDK, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *GrantKitDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *GrantKitDataSourceModel
	var item types.Object

	resp.Diagnostics.Append(req.Config.Get(ctx, &item)...)
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