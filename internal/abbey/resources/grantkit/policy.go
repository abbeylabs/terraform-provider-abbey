package grantkit

import (
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type Policy struct {
	Bundle types.String `tfsdk:"bundle"`
	Query  types.String `tfsdk:"query"`
}

func PolicyAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"bundle": types.StringType,
		"query":  types.StringType,
	}
}

func PolicyType() attr.Type {
	return types.ObjectType{AttrTypes: PolicyAttrTypes()}
}
