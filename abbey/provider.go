package abbey

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"abbey_requestable": resourceRequestable(),
		},
		DataSourcesMap: map[string]*schema.Resource{},
	}
}
