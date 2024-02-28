package output

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type Output struct {
	Location  types.String `tfsdk:"location"`
	Append    types.String `tfsdk:"append"`
	Overwrite types.String `tfsdk:"overwrite"`
}
