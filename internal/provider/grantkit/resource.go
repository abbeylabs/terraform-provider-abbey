package grantkit

import (
	"abbey/v2/internal/shared/models/grant"
	"abbey/v2/internal/shared/models/grant_workflow"
	"abbey/v2/internal/shared/models/output"
	"abbey/v2/internal/shared/models/policy"
	"abbey/v2/internal/shared/models/request"
	"abbey/v2/internal/shared/models/review"
	"abbey/v2/internal/shared/models/reviewers"
	"abbey/v2/internal/shared/models/step"
	"abbey/v2/internal/utils"
	"context"
	"fmt"
	"github.com/go-provider-sdk/pkg/client"
	"github.com/go-provider-sdk/pkg/grantkits"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ensure we implement the needed interfaces
var _ resource.Resource = &GrantKitResource{}
var _ resource.ResourceWithImportState = &GrantKitResource{}

// constructor
func NewGrantKitResource() resource.Resource {
	return &GrantKitResource{}
}

// client wrapper
type GrantKitResource struct {
	client *client.Client
}

type GrantKitResourceModel struct {
	Id                    types.String                  `tfsdk:"id"`
	Name                  types.String                  `tfsdk:"name"`
	CurrentVersionId      types.String                  `tfsdk:"current_version_id"`
	Description           types.String                  `tfsdk:"description"`
	MaxGrantDurationInSec types.Float64                 `tfsdk:"max_grant_duration_in_sec"`
	Workflow              *grant_workflow.GrantWorkflow `tfsdk:"workflow"`
	Policies              []policy.Policy               `tfsdk:"policies"`
	Output                *output.Output                `tfsdk:"output"`
	Grants                []grant.Grant                 `tfsdk:"grants"`
	ResourceType          types.String                  `tfsdk:"resource_type"`
	Requests              []request.Request             `tfsdk:"requests"`
	CreatedAt             types.String                  `tfsdk:"created_at"`
	UpdatedAt             types.String                  `tfsdk:"updated_at"`
}

func (r *GrantKitResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_grant_kit"
}

func (r *GrantKitResource) Schema(_ context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "id",
				Computed:    true,
				Optional:    true,
			},

			"name": schema.StringAttribute{
				Description: "name",
				Required:    true,
			},

			"current_version_id": schema.StringAttribute{
				Description: "current_version_id",
				Computed:    true,
				Optional:    true,
			},

			"description": schema.StringAttribute{
				Description: "description",
				Required:    true,
			},

			"max_grant_duration_in_sec": schema.Float64Attribute{
				Description: "max_grant_duration_in_sec",
				Computed:    true,
				Optional:    true,
			},

			"workflow": schema.SingleNestedAttribute{
				Description: "workflow",
				Optional:    true,

				Attributes: map[string]schema.Attribute{
					"steps": schema.ListNestedAttribute{
						Description: "steps",
						Optional:    true,

						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"reviewers": schema.SingleNestedAttribute{
									Description: "reviewers",
									Optional:    true,

									Attributes: map[string]schema.Attribute{
										"one_of": schema.ListAttribute{
											Description: "one_of",
											Optional:    true,

											ElementType: types.StringType,
										},

										"all_of": schema.ListAttribute{
											Description: "all_of",
											Optional:    true,

											ElementType: types.StringType,
										},
									},
								},

								"skip_if": schema.ListNestedAttribute{
									Description: "skip_if",
									Optional:    true,

									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"bundle": schema.StringAttribute{
												Description: "bundle",
												Optional:    true,
											},

											"query": schema.StringAttribute{
												Description: "query",
												Optional:    true,
											},
										},
									},
								},
							},
						},
					},
				},
			},

			"policies": schema.ListNestedAttribute{
				Description: "policies",
				Optional:    true,

				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"bundle": schema.StringAttribute{
							Description: "bundle",
							Optional:    true,
						},

						"query": schema.StringAttribute{
							Description: "query",
							Optional:    true,
						},
					},
				},
			},

			"output": schema.SingleNestedAttribute{
				Description: "output",
				Required:    true,

				Attributes: map[string]schema.Attribute{
					"location": schema.StringAttribute{
						Description: "location",
						Required:    true,
					},

					"append": schema.StringAttribute{
						Description: "append",
						Optional:    true,
					},

					"overwrite": schema.StringAttribute{
						Description: "overwrite",
						Optional:    true,
					},
				},
			},

			"grants": schema.ListNestedAttribute{
				Description: "grants",
				Computed:    true,
				Optional:    true,

				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "id",
							Required:    true,
						},

						"grant_kit_id": schema.StringAttribute{
							Description: "grant_kit_id",
							Required:    true,
						},

						"grant_kit_version_id": schema.StringAttribute{
							Description: "grant_kit_version_id",
							Required:    true,
						},

						"user_id": schema.StringAttribute{
							Description: "user_id",
							Required:    true,
						},

						"request_id": schema.StringAttribute{
							Description: "request_id",
							Required:    true,
						},

						"organization_id": schema.StringAttribute{
							Description: "organization_id",
							Required:    true,
						},

						"deleted": schema.BoolAttribute{
							Description: "deleted",
							Required:    true,
						},

						"created_at": schema.StringAttribute{
							Description: "created_at",
							Required:    true,
						},

						"updated_at": schema.StringAttribute{
							Description: "updated_at",
							Required:    true,
						},
					},
				},
			},

			"resource_type": schema.StringAttribute{
				Description: "resource_type",
				Computed:    true,
				Optional:    true,
			},

			"requests": schema.ListNestedAttribute{
				Description: "requests",
				Computed:    true,
				Optional:    true,

				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "id",
							Required:    true,
						},

						"grant_kit_id": schema.StringAttribute{
							Description: "grant_kit_id",
							Required:    true,
						},

						"grant_kit_version_id": schema.StringAttribute{
							Description: "grant_kit_version_id",
							Required:    true,
						},

						"grant_kit_name": schema.StringAttribute{
							Description: "grant_kit_name",
							Optional:    true,
						},

						"reason": schema.StringAttribute{
							Description: "reason",
							Required:    true,
						},

						"user_id": schema.StringAttribute{
							Description: "user_id",
							Required:    true,
						},

						"status": schema.StringAttribute{
							Description: "status",
							Required:    true,
						},

						"reviews": schema.ListNestedAttribute{
							Description: "reviews",
							Optional:    true,

							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										Description: "id",
										Required:    true,
									},

									"user_id": schema.StringAttribute{
										Description: "user_id",
										Required:    true,
									},

									"user_email": schema.StringAttribute{
										Description: "user_email",
										Optional:    true,
									},

									"request_id": schema.StringAttribute{
										Description: "request_id",
										Required:    true,
									},

									"status": schema.StringAttribute{
										Description: "status",
										Required:    true,
									},

									"request_reason": schema.StringAttribute{
										Description: "request_reason",
										Required:    true,
									},

									"reason": schema.StringAttribute{
										Description: "reason",
										Required:    true,
									},

									"grant_kit_version_id": schema.StringAttribute{
										Description: "grant_kit_version_id",
										Required:    true,
									},

									"grant_kit_name": schema.StringAttribute{
										Description: "grant_kit_name",
										Required:    true,
									},

									"grant_id": schema.StringAttribute{
										Description: "grant_id",
										Required:    true,
									},

									"grant": schema.SingleNestedAttribute{
										Description: "grant",
										Optional:    true,

										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												Description: "id",
												Required:    true,
											},

											"grant_kit_id": schema.StringAttribute{
												Description: "grant_kit_id",
												Required:    true,
											},

											"grant_kit_version_id": schema.StringAttribute{
												Description: "grant_kit_version_id",
												Required:    true,
											},

											"user_id": schema.StringAttribute{
												Description: "user_id",
												Required:    true,
											},

											"request_id": schema.StringAttribute{
												Description: "request_id",
												Required:    true,
											},

											"organization_id": schema.StringAttribute{
												Description: "organization_id",
												Required:    true,
											},

											"deleted": schema.BoolAttribute{
												Description: "deleted",
												Required:    true,
											},

											"created_at": schema.StringAttribute{
												Description: "created_at",
												Required:    true,
											},

											"updated_at": schema.StringAttribute{
												Description: "updated_at",
												Required:    true,
											},
										},
									},

									"created_at": schema.StringAttribute{
										Description: "created_at",
										Required:    true,
									},

									"updated_at": schema.StringAttribute{
										Description: "updated_at",
										Required:    true,
									},

									"pull_request": schema.StringAttribute{
										Description: "pull_request",
										Required:    true,
									},
								},
							},
						},

						"grant_id": schema.StringAttribute{
							Description: "grant_id",
							Required:    true,
						},

						"created_at": schema.StringAttribute{
							Description: "created_at",
							Required:    true,
						},

						"updated_at": schema.StringAttribute{
							Description: "updated_at",
							Required:    true,
						},

						"pull_request": schema.StringAttribute{
							Description: "pull_request",
							Required:    true,
						},
					},
				},
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
		},
	}
}

func (r *GrantKitResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *GrantKitResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var dataModel GrantKitResourceModel
	utils.PopulateModelData(ctx, &dataModel, resp.Diagnostics, req.State.Get)

	if resp.Diagnostics.HasError() {
		return
	}

	GrantKitIdOrName := dataModel.Id.ValueString()

	clientResponse, err := r.client.GrantKits.GetGrantKitById(ctx, GrantKitIdOrName)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unexpected error calling GrantKits.GetGrantKitById",
			err.Error(),
		)

		return
	}

	grantKit := clientResponse.Data

	dataModel.Id = utils.NullableString(grantKit.GetId())

	dataModel.Name = utils.NullableString(grantKit.GetName())

	dataModel.CurrentVersionId = utils.NullableString(grantKit.GetCurrentVersionId())

	dataModel.Description = utils.NullableString(grantKit.GetDescription())

	dataModel.MaxGrantDurationInSec = utils.NullableFloat64(grantKit.GetMaxGrantDurationInSec())

	if grantKit.Workflow != nil {
		dataModel.Workflow = utils.NullableObject(grantKit.Workflow, grant_workflow.GrantWorkflow{
			Steps: utils.MapList(grantKit.GetWorkflow().Steps, func(from grantkits.Step) step.Step {
				return step.Step{
					Reviewers: utils.NullableObject(from.Reviewers, reviewers.Reviewers{
						OneOf: utils.ToList(ctx, from.Reviewers.OneOf, types.StringType, &resp.Diagnostics),

						AllOf: utils.ToList(ctx, from.Reviewers.AllOf, types.StringType, &resp.Diagnostics),
					}),

					SkipIf: utils.MapList(from.SkipIf, func(from grantkits.Policy) policy.Policy {
						return policy.Policy{
							Bundle: utils.NullableString(from.GetBundle()),

							Query: utils.NullableString(from.GetQuery()),
						}
					}),
				}
			}),
		})
	}

	for index, item := range grantKit.Policies {
		dataModel.Policies[index] = policy.Policy{
			Bundle: utils.NullableString(item.GetBundle()),

			Query: utils.NullableString(item.GetQuery()),
		}
	}

	if grantKit.Output != nil {
		dataModel.Output = utils.NullableObject(grantKit.Output, output.Output{
			Location: utils.NullableString(grantKit.GetOutput().GetLocation()),

			Append: utils.NullableString(grantKit.GetOutput().GetAppend()),

			Overwrite: utils.NullableString(grantKit.GetOutput().GetOverwrite()),
		})
	}

	for index, item := range grantKit.Grants {
		dataModel.Grants[index] = grant.Grant{
			Id: utils.NullableString(item.GetId()),

			GrantKitId: utils.NullableString(item.GetGrantKitId()),

			GrantKitVersionId: utils.NullableString(item.GetGrantKitVersionId()),

			UserId: utils.NullableString(item.GetUserId()),

			RequestId: utils.NullableString(item.GetRequestId()),

			OrganizationId: utils.NullableString(item.GetOrganizationId()),

			Deleted: utils.NullableBool(item.GetDeleted()),

			CreatedAt: utils.NullableString(item.GetCreatedAt()),

			UpdatedAt: utils.NullableString(item.GetUpdatedAt()),
		}
	}

	dataModel.ResourceType = utils.NullableString(grantKit.GetResourceType())

	for index, item := range grantKit.Requests {
		dataModel.Requests[index] = request.Request{
			Id: utils.NullableString(item.GetId()),

			GrantKitId: utils.NullableString(item.GetGrantKitId()),

			GrantKitVersionId: utils.NullableString(item.GetGrantKitVersionId()),

			GrantKitName: utils.NullableString(item.GetGrantKitName()),

			Reason: utils.NullableString(item.GetReason()),

			UserId: utils.NullableString(item.GetUserId()),

			Status: types.StringValue(string(*item.GetStatus())),

			Reviews: utils.MapList(item.Reviews, func(from grantkits.Review) review.Review {
				return review.Review{
					Id: utils.NullableString(from.GetId()),

					UserId: utils.NullableString(from.GetUserId()),

					UserEmail: utils.NullableString(from.GetUserEmail()),

					RequestId: utils.NullableString(from.GetRequestId()),

					Status: types.StringValue(string(*from.GetStatus())),

					RequestReason: utils.NullableString(from.GetRequestReason()),

					Reason: utils.NullableString(from.GetReason()),

					GrantKitVersionId: utils.NullableString(from.GetGrantKitVersionId()),

					GrantKitName: utils.NullableString(from.GetGrantKitName()),

					GrantId: utils.NullableString(from.GetGrantId()),

					Grant: utils.NullableObject(from.Grant, grant.Grant{
						Id: utils.NullableString(from.Grant.GetId()),

						GrantKitId: utils.NullableString(from.Grant.GetGrantKitId()),

						GrantKitVersionId: utils.NullableString(from.Grant.GetGrantKitVersionId()),

						UserId: utils.NullableString(from.Grant.GetUserId()),

						RequestId: utils.NullableString(from.Grant.GetRequestId()),

						OrganizationId: utils.NullableString(from.Grant.GetOrganizationId()),

						Deleted: utils.NullableBool(from.Grant.GetDeleted()),

						CreatedAt: utils.NullableString(from.Grant.GetCreatedAt()),

						UpdatedAt: utils.NullableString(from.Grant.GetUpdatedAt()),
					}),

					CreatedAt: utils.NullableString(from.GetCreatedAt()),

					UpdatedAt: utils.NullableString(from.GetUpdatedAt()),

					PullRequest: utils.NullableString(from.GetPullRequest()),
				}
			}),

			GrantId: utils.NullableString(item.GetGrantId()),

			CreatedAt: utils.NullableString(item.GetCreatedAt()),

			UpdatedAt: utils.NullableString(item.GetUpdatedAt()),

			PullRequest: utils.NullableString(item.GetPullRequest()),
		}
	}

	dataModel.CreatedAt = utils.NullableString(grantKit.GetCreatedAt())

	dataModel.UpdatedAt = utils.NullableString(grantKit.GetUpdatedAt())

	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &dataModel)...)
}

func (r *GrantKitResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var dataModel GrantKitResourceModel
	utils.PopulateModelData(ctx, &dataModel, resp.Diagnostics, req.Plan.Get)

	if resp.Diagnostics.HasError() {
		return
	}

	requestBody := grantkits.GrantKitCreateParams{
		Name:        dataModel.Name.ValueStringPointer(),
		Description: dataModel.Description.ValueStringPointer(),

		Workflow: utils.NullableTfStateObject(dataModel.Workflow, func(from *grant_workflow.GrantWorkflow) grantkits.GrantWorkflow {
			return grantkits.GrantWorkflow{
				Steps: utils.MapList(from.Steps, func(from step.Step) grantkits.Step {
					return grantkits.Step{

						Reviewers: utils.NullableTfStateObject(from.Reviewers, func(from *reviewers.Reviewers) grantkits.Reviewers {
							return grantkits.Reviewers{
								OneOf: utils.FromListToPrimitiveSlice[string](ctx, from.OneOf, &resp.Diagnostics),
								AllOf: utils.FromListToPrimitiveSlice[string](ctx, from.AllOf, &resp.Diagnostics),
							}
						}),
						SkipIf: utils.MapList(from.SkipIf, func(from policy.Policy) grantkits.Policy {
							return grantkits.Policy{
								Bundle: from.Bundle.ValueStringPointer(),
								Query:  from.Query.ValueStringPointer(),
							}
						}),
					}
				}),
			}
		}),
		Policies: utils.MapList(dataModel.Policies, func(from policy.Policy) grantkits.Policy {
			return grantkits.Policy{
				Bundle: from.Bundle.ValueStringPointer(),
				Query:  from.Query.ValueStringPointer(),
			}
		}),

		Output: utils.NullableTfStateObject(dataModel.Output, func(from *output.Output) grantkits.Output {
			return grantkits.Output{
				Location:  from.Location.ValueStringPointer(),
				Append:    from.Append.ValueStringPointer(),
				Overwrite: from.Overwrite.ValueStringPointer(),
			}
		}),
	}

	clientResponse, err := r.client.GrantKits.CreateGrantKit(ctx, requestBody)

	if err != nil {
		resp.Diagnostics.AddError(
			"Error Creating GrantKit",
			err.Error(),
		)

		return
	}

	grantKit := clientResponse.Data
	dataModel.Id = utils.NullableString(grantKit.GetId())

	dataModel.Name = utils.NullableString(grantKit.GetName())

	dataModel.CurrentVersionId = utils.NullableString(grantKit.GetCurrentVersionId())

	dataModel.Description = utils.NullableString(grantKit.GetDescription())

	dataModel.MaxGrantDurationInSec = utils.NullableFloat64(grantKit.GetMaxGrantDurationInSec())

	if grantKit.Workflow != nil {
		dataModel.Workflow = utils.NullableObject(grantKit.Workflow, grant_workflow.GrantWorkflow{
			Steps: utils.MapList(grantKit.GetWorkflow().Steps, func(from grantkits.Step) step.Step {
				return step.Step{
					Reviewers: utils.NullableObject(from.Reviewers, reviewers.Reviewers{
						OneOf: utils.ToList(ctx, from.Reviewers.OneOf, types.StringType, &resp.Diagnostics),

						AllOf: utils.ToList(ctx, from.Reviewers.AllOf, types.StringType, &resp.Diagnostics),
					}),

					SkipIf: utils.MapList(from.SkipIf, func(from grantkits.Policy) policy.Policy {
						return policy.Policy{
							Bundle: utils.NullableString(from.GetBundle()),

							Query: utils.NullableString(from.GetQuery()),
						}
					}),
				}
			}),
		})
	}

	for index, item := range grantKit.Policies {
		dataModel.Policies[index] = policy.Policy{
			Bundle: utils.NullableString(item.GetBundle()),

			Query: utils.NullableString(item.GetQuery()),
		}
	}

	if grantKit.Output != nil {
		dataModel.Output = utils.NullableObject(grantKit.Output, output.Output{
			Location: utils.NullableString(grantKit.GetOutput().GetLocation()),

			Append: utils.NullableString(grantKit.GetOutput().GetAppend()),

			Overwrite: utils.NullableString(grantKit.GetOutput().GetOverwrite()),
		})
	}

	for index, item := range grantKit.Grants {
		dataModel.Grants[index] = grant.Grant{
			Id: utils.NullableString(item.GetId()),

			GrantKitId: utils.NullableString(item.GetGrantKitId()),

			GrantKitVersionId: utils.NullableString(item.GetGrantKitVersionId()),

			UserId: utils.NullableString(item.GetUserId()),

			RequestId: utils.NullableString(item.GetRequestId()),

			OrganizationId: utils.NullableString(item.GetOrganizationId()),

			Deleted: utils.NullableBool(item.GetDeleted()),

			CreatedAt: utils.NullableString(item.GetCreatedAt()),

			UpdatedAt: utils.NullableString(item.GetUpdatedAt()),
		}
	}

	dataModel.ResourceType = utils.NullableString(grantKit.GetResourceType())

	for index, item := range grantKit.Requests {
		dataModel.Requests[index] = request.Request{
			Id: utils.NullableString(item.GetId()),

			GrantKitId: utils.NullableString(item.GetGrantKitId()),

			GrantKitVersionId: utils.NullableString(item.GetGrantKitVersionId()),

			GrantKitName: utils.NullableString(item.GetGrantKitName()),

			Reason: utils.NullableString(item.GetReason()),

			UserId: utils.NullableString(item.GetUserId()),

			Status: types.StringValue(string(*item.GetStatus())),

			Reviews: utils.MapList(item.Reviews, func(from grantkits.Review) review.Review {
				return review.Review{
					Id: utils.NullableString(from.GetId()),

					UserId: utils.NullableString(from.GetUserId()),

					UserEmail: utils.NullableString(from.GetUserEmail()),

					RequestId: utils.NullableString(from.GetRequestId()),

					Status: types.StringValue(string(*from.GetStatus())),

					RequestReason: utils.NullableString(from.GetRequestReason()),

					Reason: utils.NullableString(from.GetReason()),

					GrantKitVersionId: utils.NullableString(from.GetGrantKitVersionId()),

					GrantKitName: utils.NullableString(from.GetGrantKitName()),

					GrantId: utils.NullableString(from.GetGrantId()),

					Grant: utils.NullableObject(from.Grant, grant.Grant{
						Id: utils.NullableString(from.Grant.GetId()),

						GrantKitId: utils.NullableString(from.Grant.GetGrantKitId()),

						GrantKitVersionId: utils.NullableString(from.Grant.GetGrantKitVersionId()),

						UserId: utils.NullableString(from.Grant.GetUserId()),

						RequestId: utils.NullableString(from.Grant.GetRequestId()),

						OrganizationId: utils.NullableString(from.Grant.GetOrganizationId()),

						Deleted: utils.NullableBool(from.Grant.GetDeleted()),

						CreatedAt: utils.NullableString(from.Grant.GetCreatedAt()),

						UpdatedAt: utils.NullableString(from.Grant.GetUpdatedAt()),
					}),

					CreatedAt: utils.NullableString(from.GetCreatedAt()),

					UpdatedAt: utils.NullableString(from.GetUpdatedAt()),

					PullRequest: utils.NullableString(from.GetPullRequest()),
				}
			}),

			GrantId: utils.NullableString(item.GetGrantId()),

			CreatedAt: utils.NullableString(item.GetCreatedAt()),

			UpdatedAt: utils.NullableString(item.GetUpdatedAt()),

			PullRequest: utils.NullableString(item.GetPullRequest()),
		}
	}

	dataModel.CreatedAt = utils.NullableString(grantKit.GetCreatedAt())

	dataModel.UpdatedAt = utils.NullableString(grantKit.GetUpdatedAt())

	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &dataModel)...)
}

func (r *GrantKitResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var dataModel = &GrantKitResourceModel{}
	utils.PopulateModelData(ctx, &dataModel, resp.Diagnostics, req.State.Get)

	if resp.Diagnostics.HasError() {
		return
	}

	GrantKitIdOrName := dataModel.Id.ValueString()

	_, err := r.client.GrantKits.DeleteGrantKit(ctx, GrantKitIdOrName)

	if err != nil {
		resp.Diagnostics.AddError(
			"Error Deleting GrantKit",
			err.Error(),
		)
	}
}

func (r *GrantKitResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var stateModel = &GrantKitResourceModel{}
	var dataModel = &GrantKitResourceModel{}
	utils.PopulateModelData(ctx, &stateModel, resp.Diagnostics, req.State.Get)
	utils.PopulateModelData(ctx, &dataModel, resp.Diagnostics, req.Plan.Get)

	if resp.Diagnostics.HasError() {
		return
	}

	GrantKitIdOrName := stateModel.Id.ValueString()

	requestBody := grantkits.GrantKitUpdateParams{
		Name:        dataModel.Name.ValueStringPointer(),
		Description: dataModel.Description.ValueStringPointer(),

		Workflow: utils.NullableTfStateObject(dataModel.Workflow, func(from *grant_workflow.GrantWorkflow) grantkits.GrantWorkflow {
			return grantkits.GrantWorkflow{
				Steps: utils.MapList(from.Steps, func(from step.Step) grantkits.Step {
					return grantkits.Step{

						Reviewers: utils.NullableTfStateObject(from.Reviewers, func(from *reviewers.Reviewers) grantkits.Reviewers {
							return grantkits.Reviewers{
								OneOf: utils.FromListToPrimitiveSlice[string](ctx, from.OneOf, &resp.Diagnostics),
								AllOf: utils.FromListToPrimitiveSlice[string](ctx, from.AllOf, &resp.Diagnostics),
							}
						}),
						SkipIf: utils.MapList(from.SkipIf, func(from policy.Policy) grantkits.Policy {
							return grantkits.Policy{
								Bundle: from.Bundle.ValueStringPointer(),
								Query:  from.Query.ValueStringPointer(),
							}
						}),
					}
				}),
			}
		}),

		Output: utils.NullableTfStateObject(dataModel.Output, func(from *output.Output) grantkits.Output {
			return grantkits.Output{
				Location:  from.Location.ValueStringPointer(),
				Append:    from.Append.ValueStringPointer(),
				Overwrite: from.Overwrite.ValueStringPointer(),
			}
		}),
		Policies: utils.MapList(dataModel.Policies, func(from policy.Policy) grantkits.Policy {
			return grantkits.Policy{
				Bundle: from.Bundle.ValueStringPointer(),
				Query:  from.Query.ValueStringPointer(),
			}
		}),
	}

	clientResponse, err := r.client.GrantKits.UpdateGrantKit(ctx, GrantKitIdOrName, requestBody)

	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating GrantKit",
			err.Error(),
		)

		return
	}
	grantKit := clientResponse.Data
	dataModel.Id = utils.NullableString(grantKit.GetId())

	dataModel.Name = utils.NullableString(grantKit.GetName())

	dataModel.CurrentVersionId = utils.NullableString(grantKit.GetCurrentVersionId())

	dataModel.Description = utils.NullableString(grantKit.GetDescription())

	dataModel.MaxGrantDurationInSec = utils.NullableFloat64(grantKit.GetMaxGrantDurationInSec())

	if grantKit.Workflow != nil {
		dataModel.Workflow = utils.NullableObject(grantKit.Workflow, grant_workflow.GrantWorkflow{
			Steps: utils.MapList(grantKit.GetWorkflow().Steps, func(from grantkits.Step) step.Step {
				return step.Step{
					Reviewers: utils.NullableObject(from.Reviewers, reviewers.Reviewers{
						OneOf: utils.ToList(ctx, from.Reviewers.OneOf, types.StringType, &resp.Diagnostics),

						AllOf: utils.ToList(ctx, from.Reviewers.AllOf, types.StringType, &resp.Diagnostics),
					}),

					SkipIf: utils.MapList(from.SkipIf, func(from grantkits.Policy) policy.Policy {
						return policy.Policy{
							Bundle: utils.NullableString(from.GetBundle()),

							Query: utils.NullableString(from.GetQuery()),
						}
					}),
				}
			}),
		})
	}

	for index, item := range grantKit.Policies {
		dataModel.Policies[index] = policy.Policy{
			Bundle: utils.NullableString(item.GetBundle()),

			Query: utils.NullableString(item.GetQuery()),
		}
	}

	if grantKit.Output != nil {
		dataModel.Output = utils.NullableObject(grantKit.Output, output.Output{
			Location: utils.NullableString(grantKit.GetOutput().GetLocation()),

			Append: utils.NullableString(grantKit.GetOutput().GetAppend()),

			Overwrite: utils.NullableString(grantKit.GetOutput().GetOverwrite()),
		})
	}

	for index, item := range grantKit.Grants {
		dataModel.Grants[index] = grant.Grant{
			Id: utils.NullableString(item.GetId()),

			GrantKitId: utils.NullableString(item.GetGrantKitId()),

			GrantKitVersionId: utils.NullableString(item.GetGrantKitVersionId()),

			UserId: utils.NullableString(item.GetUserId()),

			RequestId: utils.NullableString(item.GetRequestId()),

			OrganizationId: utils.NullableString(item.GetOrganizationId()),

			Deleted: utils.NullableBool(item.GetDeleted()),

			CreatedAt: utils.NullableString(item.GetCreatedAt()),

			UpdatedAt: utils.NullableString(item.GetUpdatedAt()),
		}
	}

	dataModel.ResourceType = utils.NullableString(grantKit.GetResourceType())

	for index, item := range grantKit.Requests {
		dataModel.Requests[index] = request.Request{
			Id: utils.NullableString(item.GetId()),

			GrantKitId: utils.NullableString(item.GetGrantKitId()),

			GrantKitVersionId: utils.NullableString(item.GetGrantKitVersionId()),

			GrantKitName: utils.NullableString(item.GetGrantKitName()),

			Reason: utils.NullableString(item.GetReason()),

			UserId: utils.NullableString(item.GetUserId()),

			Status: types.StringValue(string(*item.GetStatus())),

			Reviews: utils.MapList(item.Reviews, func(from grantkits.Review) review.Review {
				return review.Review{
					Id: utils.NullableString(from.GetId()),

					UserId: utils.NullableString(from.GetUserId()),

					UserEmail: utils.NullableString(from.GetUserEmail()),

					RequestId: utils.NullableString(from.GetRequestId()),

					Status: types.StringValue(string(*from.GetStatus())),

					RequestReason: utils.NullableString(from.GetRequestReason()),

					Reason: utils.NullableString(from.GetReason()),

					GrantKitVersionId: utils.NullableString(from.GetGrantKitVersionId()),

					GrantKitName: utils.NullableString(from.GetGrantKitName()),

					GrantId: utils.NullableString(from.GetGrantId()),

					Grant: utils.NullableObject(from.Grant, grant.Grant{
						Id: utils.NullableString(from.Grant.GetId()),

						GrantKitId: utils.NullableString(from.Grant.GetGrantKitId()),

						GrantKitVersionId: utils.NullableString(from.Grant.GetGrantKitVersionId()),

						UserId: utils.NullableString(from.Grant.GetUserId()),

						RequestId: utils.NullableString(from.Grant.GetRequestId()),

						OrganizationId: utils.NullableString(from.Grant.GetOrganizationId()),

						Deleted: utils.NullableBool(from.Grant.GetDeleted()),

						CreatedAt: utils.NullableString(from.Grant.GetCreatedAt()),

						UpdatedAt: utils.NullableString(from.Grant.GetUpdatedAt()),
					}),

					CreatedAt: utils.NullableString(from.GetCreatedAt()),

					UpdatedAt: utils.NullableString(from.GetUpdatedAt()),

					PullRequest: utils.NullableString(from.GetPullRequest()),
				}
			}),

			GrantId: utils.NullableString(item.GetGrantId()),

			CreatedAt: utils.NullableString(item.GetCreatedAt()),

			UpdatedAt: utils.NullableString(item.GetUpdatedAt()),

			PullRequest: utils.NullableString(item.GetPullRequest()),
		}
	}

	dataModel.CreatedAt = utils.NullableString(grantKit.GetCreatedAt())

	dataModel.UpdatedAt = utils.NullableString(grantKit.GetUpdatedAt())

	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &dataModel)...)
}

func (r *GrantKitResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
