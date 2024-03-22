package reviewers

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type Reviewers struct {
	OneOf types.List `tfsdk:"one_of"`
	AllOf types.List `tfsdk:"all_of"`
}
