package provider

import (
	"context"
	"github.com/go-provider-sdk/pkg/client"
	"github.com/go-provider-sdk/pkg/clientconfig"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"abbey/v2/internal/provider/grantkit"

	"abbey/v2/internal/provider/identity"

	"abbey/v2/internal/provider/demo"
)

// Ensure Provider satisfies various provider interfaces.
var _ provider.Provider = &Provider{}

type Provider struct {
	version string
}

type abbeyProviderModel struct {
	ServerUrl  types.String `tfsdk:"server_url"`
	BearerAuth types.String `tfsdk:"bearer_auth"`
}

func (p *Provider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "abbey"
	resp.Version = "0.2.7"
}

func (p *Provider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"server_url": schema.StringAttribute{
				Optional:    true,
				Required:    false,
				Sensitive:   false,
				Description: "The API host.",
			},
			"bearer_auth": schema.StringAttribute{
				Required:    true,
				Sensitive:   true,
				Description: "The authentication token.",
			},
		},
	}
}

func (p *Provider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var dataModel abbeyProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &dataModel)...)

	if resp.Diagnostics.HasError() {
		return
	}

	serverUrl := dataModel.ServerUrl.ValueString()
	if serverUrl == "" {
		serverUrl = "https://api.abbey.io/v1"
	}

	if dataModel.BearerAuth.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("bearer_auth"),
			"Unknown BearerAuth",
			"Cannot create API client without bearer_auth.",
		)
		return
	}

	config := clientconfig.NewConfig()
	config.SetBaseUrl(serverUrl)
	config.SetAccessToken(dataModel.BearerAuth.ValueString())
	apiClient := client.NewClient(config)

	// Example of setting the client in resp
	resp.DataSourceData = apiClient
	resp.ResourceData = apiClient
}

func (p *Provider) Resources(ctx context.Context) []func() resource.Resource {
	resources := []func() resource.Resource{}
	resources = append(resources, grantkit.NewGrantKitResource)
	resources = append(resources, identity.NewIdentityResource)
	resources = append(resources, demo.NewDemoResource)
	return resources
}

func (p *Provider) DataSources(ctx context.Context) []func() datasource.DataSource {
	dataSources := []func() datasource.DataSource{}
	return dataSources
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &Provider{
			version: version,
		}
	}
}
