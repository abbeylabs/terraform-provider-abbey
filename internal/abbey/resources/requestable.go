package resources

import (
	"bytes"
	"context"
	_ "embed"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"abbey.so/terraform-provider-abbey/internal/abbey/provider"
)

func Requestable() resource.Resource {
	return &requestable{}
}

var (
	_ resource.Resource              = (*requestable)(nil)
	_ resource.ResourceWithConfigure = (*requestable)(nil)
)

type requestable struct {
	data *provider.ResourceData
}

func (r *requestable) Metadata(
	_ context.Context,
	_ resource.MetadataRequest,
	response *resource.MetadataResponse,
) {
	response.TypeName = provider.NewTypeName("requestable")
}

//go:embed requestable_schema_append_description.md
var requestableAppendDescription string

func (r *requestable) Schema(
	_ context.Context,
	_ resource.SchemaRequest,
	response *resource.SchemaResponse,
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
				Required: true,
				Attributes: map[string]schema.Attribute{
					"builtin": schema.SingleNestedAttribute{
						Optional: true,
						Attributes: map[string]schema.Attribute{
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
				Required: true,
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

func (r *requestable) Configure(
	_ context.Context,
	request resource.ConfigureRequest,
	response *resource.ConfigureResponse,
) {
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

type requestableModel struct {
	Id       types.String `tfsdk:"id"`
	Name     types.String `tfsdk:"name"`
	Workflow types.Object `tfsdk:"workflow"`
	Grant    types.Object `tfsdk:"grant"`
}

type requestableView struct {
	Id       string           `json:"id"`
	Name     string           `json:"name"`
	Workflow *WorkflowEnum    `json:"workflow,omitempty"`
	Grant    requestableGrant `json:"grant"`
}

func (w *WorkflowEnum) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	var (
		key   string
		val   attr.Value
		diags diag.Diagnostics
	)

	switch v := w.Value.(type) {
	case *BuiltinWorkflowEnum:
		key = "builtin"
		val, diags = v.ToObjectValue(ctx)
		if diags.HasError() {
			return basetypes.ObjectValue{}, diags
		}
	case nil:
		return types.ObjectNull(map[string]attr.Type{
			"builtin": types.ObjectType{
				AttrTypes: map[string]attr.Type{
					"one_of": types.ObjectType{
						AttrTypes: map[string]attr.Type{
							"reviewers": types.ListType{
								ElemType: types.ObjectType{
									AttrTypes: map[string]attr.Type{
										"auth_id": types.StringType,
									},
								},
							},
						},
					},
				},
			},
		}), diags
	default:
		return basetypes.ObjectValue{}, diag.Diagnostics{
			diag.NewAttributeErrorDiagnostic(
				path.Root("workflow"),
				fmt.Sprintf("%T", v), "",
			),
		}
	}

	return types.ObjectValue(
		map[string]attr.Type{
			key: val.Type(ctx),
		},
		map[string]attr.Value{
			key: val,
		},
	)
}

func (u *UserQueryEnum) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	var (
		key   string
		val   attr.Value
		diags diag.Diagnostics
	)

	switch v := u.Value.(type) {
	case *UserQueryAuthId:
		key = "auth_id"
		val = types.StringValue(v.String())
		if diags.HasError() {
			return basetypes.ObjectValue{}, diags
		}
	default:
		return basetypes.ObjectValue{}, diag.Diagnostics{
			diag.NewAttributeErrorDiagnostic(path.Root("user_query"), "", ""),
		}
	}

	return types.ObjectValue(
		map[string]attr.Type{
			key: val.Type(ctx),
		},
		map[string]attr.Value{
			key: val,
		},
	)
}

func (b *BuiltinWorkflowEnum) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	var (
		key    string
		objVal attr.Value
		val    attr.Value
		diags  diag.Diagnostics
	)

	switch v := b.Value.(type) {
	case *BuiltinWorkflowOneOf:
		var reviewers attr.Value

		key = "one_of"

		reviewersVal := make([]attr.Value, 0, len(v.Reviewers))
		for _, re := range v.Reviewers {
			objVal, diags = re.ToObjectValue(ctx)
			if diags.HasError() {
				return basetypes.ObjectValue{}, diags
			}

			reviewersVal = append(reviewersVal, objVal)
		}

		reviewers, diags = basetypes.NewListValue(
			types.ObjectType{
				AttrTypes: map[string]attr.Type{
					"auth_id": types.StringType,
				},
			},
			reviewersVal,
		)
		if diags.HasError() {
			return basetypes.ObjectValue{}, diags
		}

		val, diags = types.ObjectValue(
			map[string]attr.Type{
				"reviewers": reviewers.Type(ctx),
			},
			map[string]attr.Value{
				"reviewers": reviewers,
			},
		)
		if diags.HasError() {
			return basetypes.ObjectValue{}, diags
		}
	default:
		return basetypes.ObjectValue{}, diag.Diagnostics{
			diag.NewAttributeErrorDiagnostic(
				path.Root("workflow"),
				"", "",
			),
		}
	}

	return types.ObjectValue(
		map[string]attr.Type{
			key: val.Type(ctx),
		},
		map[string]attr.Value{
			key: val,
		},
	)
}

type requestableGrant struct {
	Type  string `json:"type"`
	Value any    `json:"value"`
}

func (g *requestableGrant) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	var (
		key   string
		val   attr.Value
		diags diag.Diagnostics
	)

	switch v := g.Value.(type) {
	case requestableGrantGenerate:
		key = "generate"
		val, diags = v.ToObjectValue(ctx)
		if diags.HasError() {
			return basetypes.ObjectValue{}, diags
		}
	default:
		return basetypes.ObjectValue{}, diag.Diagnostics{
			diag.NewAttributeErrorDiagnostic(
				path.Root("grant"),
				"", "",
			),
		}
	}

	return types.ObjectValue(
		map[string]attr.Type{
			key: val.Type(ctx),
		},
		map[string]attr.Value{
			key: val,
		},
	)
}

func (g *requestableGrant) UnmarshalJSON(b []byte) error {
	var (
		value any
		enum  struct {
			Type  string          `json:"type"`
			Value json.RawMessage `json:"value"`
		}
	)

	if err := json.Unmarshal(b, &enum); err != nil {
		return err
	}

	switch enum.Type {
	case "Generate":
		var x requestableGrantGenerate

		if err := json.Unmarshal(enum.Value, &x); err != nil {
			return err
		}

		value = x
	default:
		return fmt.Errorf("unknown grant type: %s", enum.Type)
	}

	g.Type = enum.Type
	g.Value = value

	return nil
}

type requestableGrantGenerate struct {
	Type  string `json:"type"`
	Value any    `json:"value"`
}

func (g *requestableGrantGenerate) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	var (
		diags  diag.Diagnostics
		objVal attr.Value
		key    string
	)

	switch v := g.Value.(type) {
	case requestableGrantGenerateGithub:
		key = "github"
		objVal, diags = v.ToObjectValue(ctx)
		if diags.HasError() {
			return basetypes.ObjectValue{}, diags
		}
	default:
		return basetypes.ObjectValue{}, diag.Diagnostics{
			diag.NewAttributeErrorDiagnostic(
				path.Root("grant").AtName("generate"),
				"", "",
			),
		}
	}

	return types.ObjectValue(
		map[string]attr.Type{
			key: objVal.Type(ctx),
		},
		map[string]attr.Value{
			key: objVal,
		},
	)
}

func (g *requestableGrantGenerate) UnmarshalJSON(b []byte) error {
	var (
		value any
		enum  struct {
			Type  string          `json:"type"`
			Value json.RawMessage `json:"value"`
		}
	)

	if err := json.Unmarshal(b, &enum); err != nil {
		return err
	}

	switch enum.Type {
	case "Github":
		var x requestableGrantGenerateGithub

		if err := json.Unmarshal(enum.Value, &x); err != nil {
			return err
		}

		value = x
	default:
		return fmt.Errorf("unknown generate type: %s", enum.Type)
	}

	g.Type = enum.Type
	g.Value = value

	return nil
}

type requestableGrantGenerateGithub struct {
	Repo   string `tfsdk:"repo" json:"repo"`
	Path   string `tfsdk:"path" json:"path"`
	Append string `tfsdk:"append" json:"append"`
}

func (g *requestableGrantGenerateGithub) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		map[string]attr.Type{
			"repo":   types.StringType,
			"path":   types.StringType,
			"append": types.StringType,
		},
		g,
	)
}

func (r *requestable) Create(
	ctx context.Context,
	request resource.CreateRequest,
	response *resource.CreateResponse,
) {
	processGenerateGithub := func(val attr.Value) json.RawMessage {
		var config requestableGrantGenerateGithub

		obj := val.(types.Object)
		diags := obj.As(ctx, &config, basetypes.ObjectAsOptions{})
		response.Diagnostics.Append(diags...)

		bs, err := json.Marshal(config)
		if err != nil {
			response.Diagnostics.AddError(
				"Unknown",
				fmt.Sprintf("Failed to serialize generate GitHub config: %v", err),
			)
			return nil
		}

		return bs
	}
	processGenerate := func(val attr.Value) json.RawMessage {
		attrs := val.(types.Object).Attributes()

		if len(attrs) != 1 {
			response.Diagnostics.AddError(
				"InvalidInput",
				fmt.Sprintf("Expected 1 generate strategy, got %d.", len(attrs)),
			)
			return nil
		}

		var config struct {
			Type  string          `json:"type"`
			Value json.RawMessage `json:"value"`
		}

		for type_, value := range attrs {
			switch type_ {
			case "github":
				config.Type = "Github"
				config.Value = processGenerateGithub(value)
				if response.Diagnostics.HasError() {
					return nil
				}
			default:
				response.Diagnostics.AddError(
					"InvalidInput",
					fmt.Sprintf("Unknown grant strategy: %s.", type_),
				)
				return nil
			}
		}

		bs, err := json.Marshal(config)
		if err != nil {
			response.Diagnostics.AddError(
				"Unknown",
				fmt.Sprintf("Failed to serialize generate config: %v", err),
			)
			return nil
		}

		return bs
	}
	processUserQuery := func(val attr.Value) *UserQueryEnum {
		attrs := val.(types.Object).Attributes()
		if len(attrs) != 1 {
			response.Diagnostics.AddError(
				"InvalidInput",
				fmt.Sprintf("Expected 1 builtin workflow type, got %d.", len(attrs)),
			)
			return nil
		}

		var userQuery UserQueryEnum

		for type_, value := range attrs {
			switch type_ {
			case "auth_id":
				x := UserQueryAuthId(value.(basetypes.StringValue).ValueString())
				userQuery.Type = UserQueryTypeAuthId
				userQuery.Value = &x
				if response.Diagnostics.HasError() {
					return nil
				}
			default:
				response.Diagnostics.AddError(
					"InvalidInput",
					fmt.Sprintf("Unknown user query type: %s.", type_),
				)
				return nil
			}
		}

		return &userQuery
	}
	processBuiltinWorkflowOneOf := func(val attr.Value) *BuiltinWorkflowOneOf {
		attrs := val.(types.Object).Attributes()
		if len(attrs) != 1 {
			response.Diagnostics.AddError(
				"InvalidInput",
				fmt.Sprintf("Expected 1 builtin workflow type, got %d.", len(attrs)),
			)
			return nil
		}

		var workflow BuiltinWorkflowOneOf

		reviewerValues := attrs["reviewers"].(basetypes.ListValue).Elements()
		for _, reviewerValue := range reviewerValues {
			reviewer := processUserQuery(reviewerValue)
			if response.Diagnostics.HasError() {
				return nil
			}

			workflow.Reviewers = append(workflow.Reviewers, *reviewer)
		}

		return &workflow
	}
	processBuiltin := func(val attr.Value) *BuiltinWorkflowEnum {
		attrs := val.(types.Object).Attributes()
		if len(attrs) != 1 {
			response.Diagnostics.AddError(
				"InvalidInput",
				fmt.Sprintf("Expected 1 builtin workflow type, got %d.", len(attrs)),
			)
			return nil
		}

		var builtinWorkflow BuiltinWorkflowEnum

		for type_, value := range attrs {
			switch type_ {
			case "one_of":
				builtinWorkflow.Type = BuiltinWorkflowTypeOneOf
				builtinWorkflow.Value = processBuiltinWorkflowOneOf(value)
				if response.Diagnostics.HasError() {
					return nil
				}
			default:
				response.Diagnostics.AddError(
					"InvalidInput",
					fmt.Sprintf("Unknown builtin workflow type: %s.", type_),
				)
				return nil
			}
		}

		return &builtinWorkflow
	}

	var model requestableModel

	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	workflowAttrs := model.Grant.Attributes()
	if len(workflowAttrs) != 1 {
		response.Diagnostics.AddError(
			"InvalidInput",
			fmt.Sprintf("Expected 1 workflow type, got %d.", len(workflowAttrs)),
		)
		return
	}

	var workflow WorkflowEnum

	for type_, value := range model.Workflow.Attributes() {
		switch type_ {
		case "builtin":
			workflow.Type = WorkflowTypeBuiltin
			workflow.Value = processBuiltin(value)
			if response.Diagnostics.HasError() {
				return
			}
		default:
			response.Diagnostics.AddError(
				"InvalidInput",
				fmt.Sprintf("Unknown workflow type: %s.", type_),
			)
			return
		}

		break
	}

	attrs := model.Grant.Attributes()
	if len(attrs) != 1 {
		response.Diagnostics.AddError(
			"InvalidInput",
			fmt.Sprintf("Expected 1 grant strategy, got %d.", len(attrs)),
		)
		return
	}

	type Grant struct {
		Type  string          `json:"type"`
		Value json.RawMessage `json:"value"`
	}

	var grant Grant

	for type_, value := range attrs {
		switch type_ {
		case "generate":
			grant.Type = "Generate"
			grant.Value = processGenerate(value)
			if response.Diagnostics.HasError() {
				return
			}
		default:
			response.Diagnostics.AddError(
				"InvalidInput",
				fmt.Sprintf("Unknown grant strategy: %s.", type_),
			)
			return
		}

		break
	}

	body := new(bytes.Buffer)
	requestBody := struct {
		Name     string       `json:"name"`
		Workflow WorkflowEnum `json:"workflow"`
		Grant    Grant        `json:"grant"`
	}{
		Name:     model.Name.ValueString(),
		Workflow: workflow,
		Grant:    grant,
	}

	err := json.NewEncoder(body).Encode(requestBody)
	if err != nil {
		response.Diagnostics.AddError(
			"Unknown",
			fmt.Sprintf("Failed to serialize generate config: %v", err),
		)
		return
	}

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		fmt.Sprintf("%s/v1/requestables", r.data.Host),
		body,
	)
	if err != nil {
		response.Diagnostics.AddError(
			"Unknown",
			fmt.Sprintf("Failed to build request: %v", err),
		)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", r.data.Token))

	resp, err := r.data.Client.Do(req)
	if err != nil {
		response.Diagnostics.AddError(
			"Unknown",
			fmt.Sprintf("Failed to execute request: %v", err),
		)
		return
	}

	if resp.StatusCode != http.StatusCreated {
		response.Diagnostics.AddError(
			"Unknown",
			fmt.Sprintf("Expected status %d, got %s.", http.StatusCreated, resp.Status),
		)
		return
	}

	var view requestableView

	err = json.NewDecoder(resp.Body).Decode(&view)
	if err != nil {
		response.Diagnostics.AddError(
			"Unknown",
			fmt.Sprintf("Failed to deserialize response body: %v", err),
		)
		return
	}

	var (
		workflowObjValue basetypes.ObjectValue
		diags            diag.Diagnostics
	)

	if view.Workflow != nil {
		workflowObjValue, diags = view.Workflow.ToObjectValue(ctx)
		response.Diagnostics.Append(diags...)
		if response.Diagnostics.HasError() {
			return
		}
	}

	grantObjValue, diags := view.Grant.ToObjectValue(ctx)
	response.Diagnostics.Append(diags...)
	if response.Diagnostics.HasError() {
		return
	}

	model.Id = types.StringValue(view.Id)
	model.Name = types.StringValue(view.Name)
	model.Workflow = workflowObjValue
	model.Grant = grantObjValue

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}

func (r *requestable) Read(
	ctx context.Context,
	request resource.ReadRequest,
	response *resource.ReadResponse,
) {
	var model requestableModel

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

	var view requestableView

	err = json.NewDecoder(resp.Body).Decode(&view)
	if err != nil {
		response.Diagnostics.AddError(
			"Unknown",
			fmt.Sprintf("Failed to deserialize response body: %v", err),
		)
		return
	}

	workflowObjValue, diags := view.Workflow.ToObjectValue(ctx)
	response.Diagnostics.Append(diags...)
	if response.Diagnostics.HasError() {
		return
	}

	objValue, diags := view.Grant.ToObjectValue(ctx)
	response.Diagnostics.Append(diags...)
	if response.Diagnostics.HasError() {
		return
	}

	model.Id = types.StringValue(view.Id)
	model.Name = types.StringValue(view.Name)
	model.Workflow = workflowObjValue
	model.Grant = objValue

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}

func (r *requestable) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	// TODO implement me
	panic("implement me")
}

func (r *requestable) Delete(
	ctx context.Context,
	request resource.DeleteRequest,
	response *resource.DeleteResponse,
) {
	var model requestableModel

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
