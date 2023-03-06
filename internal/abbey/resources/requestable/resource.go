package requestable

import (
	"bytes"
	"context"
	_ "embed"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	. "github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/moznion/go-optional"

	. "abbey.so/terraform-provider-abbey/internal/abbey/entity"
	"abbey.so/terraform-provider-abbey/internal/abbey/provider"
)

func New() Resource {
	return &resource{data: nil}
}

var (
	_ Resource              = (*resource)(nil)
	_ ResourceWithConfigure = (*resource)(nil)
)

type resource struct {
	data *provider.ResourceData
}

func (r *resource) Metadata(
	_ context.Context,
	_ MetadataRequest,
	response *MetadataResponse,
) {
	response.TypeName = provider.NewTypeName("requestable")
}

//go:embed schema_append_description.md
var requestableAppendDescription string

func (r *resource) Schema(
	_ context.Context,
	_ SchemaRequest,
	response *SchemaResponse,
) {
	response.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Required: true,
			},
			"workflow": schema.SingleNestedAttribute{
				Required:   true,
				CustomType: WorkflowType{},
				Attributes: map[string]schema.Attribute{
					"builtin": schema.SingleNestedAttribute{
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"all_of": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"reviewers": schema.ListNestedAttribute{
										Required: true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"auth_id": schema.StringAttribute{
													Optional:            true,
													Sensitive:           true,
													MarkdownDescription: "The authentication identifier of the reviewer in Abbey Labs. It may be email, phone number, or username.",
												},
											},
										},
									},
								},
							},
							"one_of": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"reviewers": schema.ListNestedAttribute{
										Required: true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"auth_id": schema.StringAttribute{
													Optional:            true,
													Sensitive:           true,
													MarkdownDescription: "The authentication identifier of the reviewer in Abbey Labs. It may be email, phone number, or username.",
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
			"grant": schema.SingleNestedAttribute{
				Required:   true,
				CustomType: GrantType{},
				Attributes: map[string]schema.Attribute{
					"generate": schema.SingleNestedAttribute{
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"github": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"repo": schema.StringAttribute{
										Required:            true,
										MarkdownDescription: "The full repository name including the owner, e.g. `abbeylabs/iiac`.",
										Validators: []validator.String{
											stringvalidator.LengthAtLeast(3),
										},
									},
									"path": schema.StringAttribute{
										Required:            true,
										MarkdownDescription: "Access code block will be generated to this file in the repo.",
										Validators: []validator.String{
											stringvalidator.LengthAtLeast(1),
										},
									},
									"append": schema.StringAttribute{
										Optional:            true,
										MarkdownDescription: requestableAppendDescription,
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

func (r *resource) Configure(
	_ context.Context,
	request ConfigureRequest,
	response *ConfigureResponse,
) {
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

	r.data = providerData
}

type Model struct {
	Id       types.String `tfsdk:"id"`
	Name     types.String `tfsdk:"name"`
	Workflow WorkflowTf   `tfsdk:"workflow"`
	Grant    GrantTf      `tfsdk:"grant"`
}

func (m Model) ToView() *View {
	var (
		workflow *Workflow
		grant    *Grant
	)

	if !m.Workflow.IsNull() && !m.Grant.IsUnknown() {
		workflow = &m.Workflow.Workflow
	}

	if !m.Grant.IsNull() && !m.Grant.IsUnknown() {
		grant = &m.Grant.Grant
	}

	return &View{
		Id:       m.Id.ValueString(),
		Name:     m.Name.ValueString(),
		Workflow: workflow,
		Grant:    grant,
		Policies: nil,
	}
}

type View struct {
	Id       string            `json:"id,omitempty"`
	Name     string            `json:"name,omitempty"`
	Workflow *Workflow         `json:"workflow,omitempty"`
	Grant    *Grant            `json:"grant,omitempty"`
	Policies Option[PolicySet] `json:"policies,omitempty"`
}

func (v View) ToModel() *Model {
	var (
		workflow WorkflowTf
		grant    GrantTf
	)

	if v.Workflow != nil {
		workflow = NewWorkflow(*v.Workflow)
	}

	if v.Grant != nil {
		grant = NewGrant(*v.Grant)
	}

	return &Model{
		Id:       types.StringValue(v.Id),
		Name:     types.StringValue(v.Name),
		Workflow: workflow,
		Grant:    grant,
	}
}

func (r *resource) Create(
	ctx context.Context,
	request CreateRequest,
	response *CreateResponse,
) {
	var model Model

	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	view := model.ToView()

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

	var newView View

	err = json.NewDecoder(resp.Body).Decode(&newView)
	if err != nil {
		response.Diagnostics.
			AddError("Unknown", fmt.Sprintf("Failed to deserialize response body: %v", err))
		return
	}

	newModel := newView.ToModel()

	response.Diagnostics.Append(response.State.Set(ctx, newModel)...)
}

func (r *resource) Read(
	ctx context.Context,
	request ReadRequest,
	response *ReadResponse,
) {
	var model Model

	response.Diagnostics.Append(request.State.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		fmt.Sprintf("%s/v1/requestables/%s", r.data.Host, model.Id.ValueString()),
		nil,
	)
	if err != nil {
		response.Diagnostics.AddError(
			"Unknown",
			fmt.Sprintf("Failed to build request: %v", err),
		)
		return
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", r.data.Token))

	resp, err := r.data.Client.Do(req)
	if err != nil {
		response.Diagnostics.AddError(
			"Unknown",
			fmt.Sprintf("Failed to execute request: %v", err),
		)
		return
	}

	if resp.StatusCode == http.StatusNotFound {
		response.State.RemoveResource(ctx)
		return
	}

	if resp.StatusCode != http.StatusOK {
		response.Diagnostics.AddError(
			"Unknown",
			fmt.Sprintf("Expected status %d, got %s.", http.StatusOK, resp.Status),
		)
		return
	}

	var view View

	err = json.NewDecoder(resp.Body).Decode(&view)
	if err != nil {
		response.Diagnostics.AddError(
			"Unknown",
			fmt.Sprintf("Failed to deserialize response body: %v", err),
		)
		return
	}

	response.Diagnostics.Append(response.State.Set(ctx, view.ToModel())...)
}

func (r *resource) Update(
	context.Context,
	UpdateRequest,
	*UpdateResponse,
) {
	// TODO implement me
	panic("implement me")
}

func (r *resource) Delete(
	ctx context.Context,
	request DeleteRequest,
	response *DeleteResponse,
) {
	var model Model

	response.Diagnostics.Append(request.State.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodDelete,
		fmt.Sprintf("%s/v1/requestables/%s", r.data.Host, model.Id.ValueString()),
		nil,
	)
	if err != nil {
		response.Diagnostics.AddError(
			"Unknown",
			fmt.Sprintf("Failed to build request: %v", err),
		)
		return
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", r.data.Token))

	resp, err := r.data.Client.Do(req)
	if err != nil {
		response.Diagnostics.AddError(
			"Unknown",
			fmt.Sprintf("Failed to execute request: %v", err),
		)
		return
	}

	if resp.StatusCode != http.StatusNoContent {
		response.Diagnostics.AddError(
			"Unknown",
			fmt.Sprintf("Expected status %d, got %s.", http.StatusNoContent, resp.Status),
		)
		return
	}

	response.State.RemoveResource(ctx)
}
