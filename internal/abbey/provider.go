package abbey

import (
	"context"
	"net/http"
	"os"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"

	abbeyprovider "abbey.so/terraform-provider-abbey/internal/abbey/provider"
	"abbey.so/terraform-provider-abbey/internal/abbey/resources"
	identityresource "abbey.so/terraform-provider-abbey/internal/abbey/resources/identity"
)

const (
	envKeyToken = "ABBEY_TOKEN"
)

type provider_ struct {
	version string

	// The following fields are populated by [provider.Provider.Configure].

	host string
}

var _ provider.Provider = (*provider_)(nil)

func New(version string, defaultHost string) func() provider.Provider {
	return func() provider.Provider {
		return &provider_{
			version: version,
			host:    defaultHost,
		}
	}
}

func (p *provider_) Metadata(
	_ context.Context,
	_ provider.MetadataRequest,
	response *provider.MetadataResponse,
) {
	response.TypeName = "abbey"
	response.Version = p.version
}

func (p *provider_) Schema(
	_ context.Context,
	_ provider.SchemaRequest,
	response *provider.SchemaResponse,
) {
	response.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"host": schema.StringAttribute{
				Optional: true,
			},
		},
	}
}

type Config struct {
	Host types.String `tfsdk:"host"`
}

func (p *provider_) Configure(
	ctx context.Context,
	request provider.ConfigureRequest,
	response *provider.ConfigureResponse,
) {
	var config Config

	response.Diagnostics.Append(request.Config.Get(ctx, &config)...)
	if response.Diagnostics.HasError() {
		return
	}

	host := p.host
	token := os.Getenv(envKeyToken)

	if !config.Host.IsNull() {
		host = config.Host.ValueString()
	}

	response.ResourceData = &abbeyprovider.ResourceData{
		Client: http.DefaultClient,
		Host:   host,
		Token:  token,
	}
}

func (p *provider_) DataSources(context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{}
}

func (p *provider_) Resources(context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		resources.Requestable,
		identityresource.New,
	}
}
