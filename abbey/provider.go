package abbey

import (
	"context"
	"os"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Meta struct {
	Host  string
	Token string
}

func Provider() *schema.Provider {
	return &schema.Provider{
		ConfigureContextFunc: func(ctx context.Context, d *schema.ResourceData) (any, diag.Diagnostics) {
			var diags diag.Diagnostics

			host := d.Get("host").(string)
			token := os.Getenv("ABBEY_TOKEN")

			return Meta{Host: host, Token: token}, diags
		},
		DataSourcesMap: map[string]*schema.Resource{},
		ResourcesMap: map[string]*schema.Resource{
			"abbey_requestable": resourceRequestable(),
		},
		Schema: map[string]*schema.Schema{
			"host": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "https://api.abbey.so",
			},
		},
	}
}
