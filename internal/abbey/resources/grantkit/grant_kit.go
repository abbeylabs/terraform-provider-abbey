package grantkit

import (
	. "abbey.so/terraform-provider-abbey/internal/abbey/tf"
)

type GrantKit struct {
	Id          String `json:"id"          tfsdk:"id"`
	Name        String `json:"name"        tfsdk:"name"`
	Description String `json:"description" tfsdk:"description"`
}
