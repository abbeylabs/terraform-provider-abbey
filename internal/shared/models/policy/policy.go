package policy

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type Policy struct {
	Bundle types.String `tfsdk:"bundle"`
	Query  types.String `tfsdk:"query"`
}
