package identity

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	. "github.com/hashicorp/terraform-plugin-framework/resource"
	. "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"abbey.so/terraform-provider-abbey/internal/abbey/provider"
)

func New() Resource {
	return &resource{data: nil}
}

type resource struct {
	data *provider.ResourceData
}

var (
	_ Resource              = (*resource)(nil)
	_ ResourceWithConfigure = (*resource)(nil)
)

func (r *resource) Metadata(
	_ context.Context,
	_ MetadataRequest,
	response *MetadataResponse,
) {
	response.TypeName = provider.NewTypeName("identity")
}

func (r *resource) Schema(
	_ context.Context,
	_ SchemaRequest,
	response *SchemaResponse,
) {
	response.Schema = Schema{
		Attributes: map[string]Attribute{
			"id":         StringAttribute{Computed: true},
			"name":       StringAttribute{Required: true},
			"created_at": StringAttribute{Computed: true},
			"linked":     StringAttribute{Optional: true},
		},
	}
}

type model struct {
	Id        types.String `tfsdk:"id"`
	CreatedAt types.String `tfsdk:"created_at"`
	Name      types.String `tfsdk:"name"`
	Linked    types.String `tfsdk:"linked"`
}

func modelFromView(v view) model {
	return model{
		Id:        basetypes.NewStringValue(v.Id),
		CreatedAt: basetypes.NewStringValue(v.CreatedAt.Format(time.RFC3339)),
		Name:      basetypes.NewStringValue(v.Name),
		Linked:    basetypes.NewStringValue(string(v.Linked)),
	}
}

type view struct {
	Id        string          `json:"id"`
	CreatedAt time.Time       `json:"created_at"`
	Name      string          `json:"name"`
	Linked    json.RawMessage `json:"linked"`
}

func (r *resource) Create(
	ctx context.Context,
	request CreateRequest,
	response *CreateResponse,
) {
	var m model

	response.Diagnostics.Append(request.Plan.Get(ctx, &m)...)
	if response.Diagnostics.HasError() {
		return
	}

	body := new(bytes.Buffer)
	requestBody := struct {
		Name   string          `json:"name"`
		Linked json.RawMessage `json:"linked"`
	}{
		Name:   m.Name.ValueString(),
		Linked: []byte(m.Linked.ValueString()),
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
		fmt.Sprintf("%s/v1/identities", r.data.Host),
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

	var v view

	err = json.NewDecoder(resp.Body).Decode(&v)
	if err != nil {
		response.Diagnostics.AddError(
			"Unknown",
			fmt.Sprintf("Failed to deserialize response body: %v", err),
		)
		return
	}

	response.Diagnostics.Append(response.State.Set(ctx, modelFromView(v))...)
}

func (r *resource) Read(
	ctx context.Context,
	request ReadRequest,
	response *ReadResponse,
) {
	var m model
	response.Diagnostics.Append(request.State.Get(ctx, &m)...)
	if response.Diagnostics.HasError() {
		return
	}

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		fmt.Sprintf("%s/v1/identities/%s", r.data.Host, m.Id.ValueString()),
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

	var v view
	err = json.NewDecoder(resp.Body).Decode(&v)
	if err != nil {
		response.Diagnostics.AddError(
			"Unknown",
			fmt.Sprintf("Failed to deserialize response body: %v", err),
		)
		return
	}

	response.Diagnostics.Append(response.State.Set(ctx, modelFromView(v))...)
}

func (r *resource) Update(
	ctx context.Context,
	request UpdateRequest,
	response *UpdateResponse,
) {
	response.Diagnostics.AddWarning(
		"Update Operation Not Implemented Yet",
		"You can workaround this by destroying your target and re-applying your configuration.\nTo destroy this resource, run:\n\n\t```\nterraform destroy -target <your resource id>```")
}

func (r *resource) Delete(
	ctx context.Context,
	request DeleteRequest,
	response *DeleteResponse,
) {
	var state model

	response.Diagnostics.Append(request.State.Get(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodDelete,
		fmt.Sprintf("%s/v1/identities/%s", r.data.Host, state.Id.ValueString()),
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
			"Unexpected Resource Configure type_",
			fmt.Sprintf("Got: %T. Please report this issue to the provider developers.", request.ProviderData),
		)
		return
	}

	r.data = providerData
}
