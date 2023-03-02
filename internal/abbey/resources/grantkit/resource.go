package grantkit

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	. "github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"abbey.so/terraform-provider-abbey/internal/abbey/provider"
	. "abbey.so/terraform-provider-abbey/internal/abbey/tf"
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

func (self *resource) Configure(ctx context.Context, request ConfigureRequest, response *ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	providerData, ok := request.ProviderData.(*provider.ResourceData)
	if !ok {
		response.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Got: %T. Please report this issue to the provider developers.", request.ProviderData),
		)
		return
	}

	self.data = providerData
}

func (self resource) Metadata(ctx context.Context, request MetadataRequest, response *MetadataResponse) {
	response.TypeName = provider.NewTypeName("grant_kit")
}

func (self resource) Schema(ctx context.Context, request SchemaRequest, response *SchemaResponse) {
	response.Schema = schema.Schema{
		Description: "The resource `grant_kit` allows you to automate access control to sensitive data." +
			"\n" +
			"This resource can be used to create access request workflows to help you with security and compliance. " +
			"You can also add OPA-based access policies for your target systems to control access by roles, attributes, " +
			"or time.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				CustomType: StringType,
				Computed:   true,
			},
			"name": schema.StringAttribute{
				CustomType:  StringType,
				Required:    true,
				Description: "The human-readable name of this resource.",
			},
			"description": schema.StringAttribute{
				CustomType:  StringType,
				Required:    true,
				Description: "The text describing what this Grant Kit is used for and what it can do.",
			},
			// "workflow": schema.SingleNestedAttribute{
			// 	Optional: true,
			// 	Description: "The workflow for _how_ someone gets access to a resource. " +
			// 		"A workflow contains a list `steps` to be run sequentially.",
			// 	Attributes: map[string]schema.Attribute{
			// 		"steps": schema.ListNestedAttribute{
			// 			Required: true,
			// 			Description: "The chain of decisions that needs to determine if access to a resource is approved or denied. " +
			// 				"Each step contains a list of `reviewers` that are responsible " +
			// 				"for submitting an approve or deny decision. " +
			// 				"A step may require approvals from `one_of` or `all_of` the reviewers.",
			// 			Validators: []validator.List{listvalidator.SizeAtLeast(1)},
			// 			NestedObject: schema.NestedAttributeObject{
			// 				Attributes: map[string]schema.Attribute{
			// 					"reviewers": schema.SingleNestedAttribute{
			// 						Required: true,
			// 						Description: "The list of Primary Identities, Secondary Identities, or Groups which are " +
			// 							"responsible for submitting an approve or deny decision.",
			// 						Attributes: map[string]schema.Attribute{
			// 							"one_of": schema.ListAttribute{
			// 								Optional:    true,
			// 								ElementType: types.StringType,
			// 								Description: "Require only one reviewer in the `reviewers` list to approve " +
			// 									"in order to advance this step.",
			// 								Validators: []validator.List{
			// 									listvalidator.ExactlyOneOf(path.Expressions{
			// 										path.MatchRelative().AtParent().AtName("all_of"),
			// 									}...),
			// 								},
			// 							},
			// 							"all_of": schema.ListAttribute{
			// 								Optional:    true,
			// 								ElementType: types.StringType,
			// 								Description: "Require all reviewers in the `reviewers` list to approve " +
			// 									"in order to advance this step.",
			// 							},
			// 						},
			// 					},
			// 					"skip_if": schema.ListNestedAttribute{
			// 						Optional:    true,
			// 						Description: "The condition that determines whether this step should be run.",
			// 						Validators: []validator.List{
			// 							listvalidator.SizeAtLeast(1),
			// 						},
			// 						NestedObject: schema.NestedAttributeObject{
			// 							Attributes: map[string]schema.Attribute{
			// 								"bundle": schema.StringAttribute{
			// 									Optional: true,
			// 									Description: "An RFC 3986 URI. Supports `github://` only. Schemes " +
			// 										"such as `https://`, `file://`, and `s3://` to come in future releases. " +
			// 										"You should use either `bundle` to contain your OPA Policies or supply them " +
			// 										"directly in the `query` field.",
			// 									Validators: []validator.String{
			// 										stringvalidator.ExactlyOneOf(path.Expressions{
			// 											path.MatchRelative().AtParent().AtName("query"),
			// 										}...),
			// 										abbeyvalidator.IsRFC3986(),
			// 									},
			// 								},
			// 								"query": schema.StringAttribute{
			// 									Optional: true,
			// 									Description: "The UTF-8 text string containing Rego rules using the " +
			// 										"Abbey OPA Framework." +
			// 										"Rules should be written using `deny[msg] { ... }` for mandatory enforcement " +
			// 										"and `warn[msg] { ... }` for advisory enforcement.",
			// 								},
			// 							},
			// 						},
			// 					},
			// 				},
			// 			},
			// 		},
			// 	},
			// },
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
			// "output": schema.SingleNestedAttribute{
			// 	Required: true,
			// 	Description: "The output represents how and where access changes should be made. " +
			// 		"This generates HCL code in a Terraform file at the `location` with either `append` or `overwrite` behavior.",
			// 	Attributes: map[string]schema.Attribute{
			// 		"location": schema.StringAttribute{
			// 			Required: true,
			// 			Description: "An RFC 3986 URI. Supports `github://` only. Schemes " +
			// 				"such as `https://`, `file://`, and `s3://` to come in future releases.",
			// 			Validators: []validator.String{
			// 				abbeyvalidator.IsRFC3986(),
			// 			},
			// 		},
			// 		"append": schema.StringAttribute{
			// 			Optional:    true,
			// 			Description: "Appends this UTF-8 text string to the file at `location`.",
			// 			Validators: []validator.String{
			// 				stringvalidator.ExactlyOneOf(path.Expressions{
			// 					path.MatchRelative().AtParent().AtName("overwrite"),
			// 				}...),
			// 			},
			// 		},
			// 		"overwrite": schema.StringAttribute{
			// 			Optional:    true,
			// 			Description: "Overwrites the file at `location` with this UTF-8 text string.",
			// 		},
			// 	},
			// },
		},
	}
}

func (self resource) Create(ctx context.Context, request CreateRequest, response *CreateResponse) {
	var plan GrantKit

	response.Diagnostics.Append(request.Plan.Get(ctx, &plan)...)
	if response.Diagnostics.HasError() {
		return
	}

	mockData := `
		{
			"id": "mockId",
			"name": "name",
			"description": "description"
		}
	`

	var model GrantKit

	if err := json.NewDecoder(strings.NewReader(mockData)).Decode(&model); err != nil {
		response.Diagnostics.AddError("Bad decode", err.Error())
		return
	}

	model.Name = plan.Name

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}

func (self resource) Read(ctx context.Context, request ReadRequest, response *ReadResponse) {
	var model GrantKit

	response.Diagnostics.Append(request.State.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	response.Diagnostics.Append(response.State.Set(ctx, model)...)
}

func (self resource) Update(ctx context.Context, request UpdateRequest, response *UpdateResponse) {
}

func (self resource) Delete(ctx context.Context, request DeleteRequest, response *DeleteResponse) {
}
