// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package validators

import (
	"abbey/v2/internal/sdk/pkg/types"
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/helpers/validatordiag"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ validator.String = DateValidator{}

type DateValidator struct {
}

func (validator DateValidator) Description(ctx context.Context) string {
	return "value must be a string in YYYY-MM-DD format"
}

func (validator DateValidator) MarkdownDescription(ctx context.Context) string {
	return validator.Description(ctx)
}

func (validator DateValidator) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	// Only validate the attribute configuration value if it is known.
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}

	if _, err := types.NewDateFromString(req.ConfigValue.ValueString()); err != nil {
		resp.Diagnostics.Append(validatordiag.InvalidAttributeTypeDiagnostic(
			req.Path,
			validator.MarkdownDescription(ctx),
			req.ConfigValue.ValueString(),
		))
		return
	}
}

// IsDate returns an AttributeValidator which ensures that any configured
// attribute value:
//
//   - Is a String.
//   - Is in YYYY-MM-DD Format.
//
// Null (unconfigured) and unknown (known after apply) values are skipped.
func IsValidDate() validator.String {
	return DateValidator{}
}
