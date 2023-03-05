package grantkit

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	. "github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"abbey.so/terraform-provider-abbey/internal/abbey/provider"
	"abbey.so/terraform-provider-abbey/internal/abbey/resources/requestable"
	abbeyvalidator "abbey.so/terraform-provider-abbey/validator"
)

var (
	_ Resource              = (*resource)(nil)
	_ ResourceWithConfigure = (*resource)(nil)
)

func New() Resource {
	return &resource{data: nil}
}

type resource struct {
	data *provider.ResourceData
}

func (r *resource) Configure(_ context.Context, request ConfigureRequest, response *ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	providerData, ok := request.ProviderData.(*provider.ResourceData)
	if !ok {
		response.Diagnostics.AddError(
			"Unexpected Resource Configure type_",
			fmt.Sprintf("Got: %T. Please report this issue to the provider developers.", request.ProviderData),
		)
		return
	}

	r.data = providerData
}

func (r resource) Metadata(_ context.Context, _ MetadataRequest, response *MetadataResponse) {
	response.TypeName = provider.NewTypeName("grant_kit")
}

func (r resource) Schema(_ context.Context, _ SchemaRequest, response *SchemaResponse) {
	response.Schema = schema.Schema{
		Description: "The resource `grant_kit` allows you to automate access control to sensitive data." +
			"\n" +
			"This resource can be used to create access request workflows to help you with security and compliance. " +
			"You can also add OPA-based access policies for your target systems to control access by roles, attributes, " +
			"or time.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "The human-readable name of this resource.",
			},
			"description": schema.StringAttribute{
				Optional:    true,
				Description: "The text describing what this Grant Kit is used for and what it can do.",
			},
			"workflow": schema.SingleNestedAttribute{
				Optional: true,
				Description: "The workflow for _how_ someone gets access to a resource. " +
					"A workflow contains a list `steps` to be run sequentially.",
				Attributes: map[string]schema.Attribute{
					"steps": schema.ListNestedAttribute{
						Required: true,
						Description: "The chain of decisions that needs to determine if access to a resource is approved or denied. " +
							"Each step contains a list of `reviewers` that are responsible " +
							"for submitting an approve or deny decision. " +
							"A step may require approvals from `one_of` or `all_of` the reviewers.",
						Validators: []validator.List{listvalidator.SizeAtLeast(1)},
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"reviewers": schema.SingleNestedAttribute{
									Required: true,
									Description: "The list of Primary Identities, Secondary Identities, or Groups which are " +
										"responsible for submitting an approve or deny decision.",
									Attributes: map[string]schema.Attribute{
										"one_of": schema.ListAttribute{
											Optional:    true,
											ElementType: types.StringType,
											Description: "Require only one reviewer in the `reviewers` list to approve " +
												"in order to advance this step.",
											Validators: []validator.List{
												listvalidator.ExactlyOneOf(path.Expressions{
													path.MatchRelative().AtParent().AtName("all_of"),
												}...),
											},
										},
										"all_of": schema.ListAttribute{
											Optional:    true,
											ElementType: types.StringType,
											Description: "Require all reviewers in the `reviewers` list to approve " +
												"in order to advance this step.",
										},
									},
								},
								"skip_if": schema.ListNestedAttribute{
									Optional:    true,
									Description: "The condition that determines whether this step should be run.",
									Validators: []validator.List{
										listvalidator.SizeAtLeast(1),
									},
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"bundle": schema.StringAttribute{
												Optional: true,
												Description: "An RFC 3986 URI. Supports `github://` only. Schemes " +
													"such as `https://`, `file://`, and `s3://` to come in future releases. " +
													"You should use either `bundle` to contain your OPA Policies or supply them " +
													"directly in the `query` field.",
												Validators: []validator.String{
													stringvalidator.ExactlyOneOf(path.Expressions{
														path.MatchRelative().AtParent().AtName("query"),
													}...),
													abbeyvalidator.IsRFC3986(),
												},
											},
											"query": schema.StringAttribute{
												Optional: true,
												Description: "The UTF-8 text string containing Rego rules using the " +
													"Abbey OPA Framework." +
													"Rules should be written using `deny[msg] { ... }` for mandatory enforcement " +
													"and `warn[msg] { ... }` for advisory enforcement.",
											},
										},
									},
								},
							},
						},
					},
				},
			},
			// "policies": schema.SingleNestedAttribute{
			// 	Optional:    true,
			// 	Description: "The access policies that determine if the resource requester get access to this resource.",
			// 	Attributes: map[string]schema.Attribute{
			// 		"grant_if": schema.ListNestedAttribute{
			// 			Optional:    true,
			// 			Description: "Determines the conditions for which this resource should be granted access to the requester.",
			// 			Validators: []validator.List{
			// 				listvalidator.AtLeastOneOf(path.Expressions{
			// 					path.MatchRelative().AtParent().AtName("revoke_if"),
			// 				}...),
			// 				listvalidator.SizeAtLeast(1),
			// 			},
			// 			NestedObject: schema.NestedAttributeObject{
			// 				Attributes: map[string]schema.Attribute{
			// 					"bundle": schema.StringAttribute{
			// 						Optional: true,
			// 						Description: "An RFC 3986 URI. Supports `github://` only. Schemes " +
			// 							"such as `https://`, `file://`, and `s3://` to come in future releases. " +
			// 							"You should use either `bundle` to contain your OPA Policies or supply them " +
			// 							"directly in the `query` field.",
			// 						Validators: []validator.String{
			// 							stringvalidator.ExactlyOneOf(path.Expressions{
			// 								path.MatchRelative().AtParent().AtName("query"),
			// 							}...),
			// 							abbeyvalidator.IsRFC3986(),
			// 						},
			// 					},
			// 					"query": schema.StringAttribute{
			// 						Optional: true,
			// 						Description: "The UTF-8 text string containing Rego rules using the " +
			// 							"Abbey OPA Framework." +
			// 							"Rules should be written using `deny[msg] { ... }` for mandatory enforcement " +
			// 							"and `warn[msg] { ... }` for advisory enforcement.",
			// 					},
			// 				},
			// 			},
			// 		},
			// 		"revoke_if": schema.ListNestedAttribute{
			// 			Optional:    true,
			// 			Description: "Determines the conditions for which access to this resource should be revoked from the requester.",
			// 			Validators: []validator.List{
			// 				listvalidator.SizeAtLeast(1),
			// 			},
			// 			NestedObject: schema.NestedAttributeObject{
			// 				Attributes: map[string]schema.Attribute{
			// 					"bundle": schema.StringAttribute{
			// 						Optional: true,
			// 						Description: "An RFC 3986 URI. Supports `github://` only. Schemes " +
			// 							"such as `https://`, `file://`, and `s3://` to come in future releases. " +
			// 							"You should use either `bundle` to contain your OPA Policies or supply them " +
			// 							"directly in the `query` field.",
			// 						Validators: []validator.String{
			// 							stringvalidator.ExactlyOneOf(path.Expressions{
			// 								path.MatchRelative().AtParent().AtName("query"),
			// 							}...),
			// 							abbeyvalidator.IsRFC3986(),
			// 						},
			// 					},
			// 					"query": schema.StringAttribute{
			// 						Optional: true,
			// 						Description: "The UTF-8 text string containing Rego rules using the " +
			// 							"Abbey OPA Framework." +
			// 							"Rules should be written using `deny[msg] { ... }` for mandatory enforcement " +
			// 							"and `warn[msg] { ... }` for advisory enforcement.",
			// 					},
			// 				},
			// 			},
			// 		},
			// 	},
			// },
			"output": schema.SingleNestedAttribute{
				Required: true,
				Description: "The output represents how and where access changes should be made. " +
					"This generates HCL code in a Terraform file at the `location` with either `append` or `overwrite` behavior.",
				Attributes: map[string]schema.Attribute{
					"location": schema.StringAttribute{
						Required: true,
						Description: "An RFC 3986 URI. Supports `github://` only. Schemes " +
							"such as `https://`, `file://`, and `s3://` to come in future releases.",
						Validators: []validator.String{
							abbeyvalidator.IsRFC3986(),
						},
					},
					"append": schema.StringAttribute{
						Optional:    true,
						Description: "Appends this UTF-8 text string to the file at `location`.",
						Validators: []validator.String{
							stringvalidator.ExactlyOneOf(path.Expressions{
								path.MatchRelative().AtParent().AtName("overwrite"),
							}...),
						},
					},
					"overwrite": schema.StringAttribute{
						Optional:    true,
						Description: "Overwrites the file at `location` with this UTF-8 text string.",
					},
				},
			},
		},
	}
}

func (r resource) Create(ctx context.Context, request CreateRequest, response *CreateResponse) {
	var plan Model

	response.Diagnostics.Append(request.Plan.Get(ctx, &plan)...)
	if response.Diagnostics.HasError() {
		return
	}

	view, diags := plan.ToRequestableView(ctx)
	response.Diagnostics.Append(diags...)
	if response.Diagnostics.HasError() {
		return
	}

	body, err := json.Marshal(view)
	if err != nil {
		response.Diagnostics.
			AddError("Unknown", fmt.Sprintf("Failed to serialize generate config: %v.", err))
		return
	}

	req, err := http.NewRequestWithContext(
		ctx, http.MethodPost,
		fmt.Sprintf("%s/v1/requestables", r.data.Host),
		bytes.NewReader(body),
	)
	if err != nil {
		response.Diagnostics.
			AddError("Unknown", fmt.Sprintf("Failed to build request: %v.", err))
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", r.data.Token))

	resp, err := r.data.Client.Do(req)
	if err != nil {
		response.Diagnostics.
			AddError("Unknown", fmt.Sprintf("Failed to execute request: %v", err))
		return
	}

	defer func() {
		if err := resp.Body.Close(); err != nil {
			response.Diagnostics.AddWarning("Unknown", fmt.Sprintf("Failed to close response body: %v", err))
		}
	}()

	if resp.StatusCode != http.StatusCreated {
		response.Diagnostics.
			AddError("Unknown", fmt.Sprintf("Expected status %d, got %s.", http.StatusCreated, resp.Status))
		return
	}

	var newView requestable.View

	err = json.NewDecoder(resp.Body).Decode(&newView)
	if err != nil {
		response.Diagnostics.
			AddError("Unknown", fmt.Sprintf("Failed to deserialize response body: %v", err))
		return
	}

	state, diags_ := ModelFromRequestableView(newView)
	response.Diagnostics.Append(diags_...)
	if response.Diagnostics.HasError() {
		return
	}

	response.Diagnostics.Append(response.State.Set(ctx, state)...)
}

func (r resource) Read(ctx context.Context, request ReadRequest, response *ReadResponse) {
	var model Model

	response.Diagnostics.Append(request.State.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	response.Diagnostics.Append(response.State.Set(ctx, model)...)
}

func (r resource) Update(ctx context.Context, request UpdateRequest, response *UpdateResponse) {
}

func (r resource) Delete(ctx context.Context, request DeleteRequest, response *DeleteResponse) {
}
