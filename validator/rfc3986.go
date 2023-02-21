package validator

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"net/url"
)

// IsRFC3986 checks if the string is an RFC 3986 URI.
func IsRFC3986() rfc3986Validator {
	return rfc3986Validator{}
}

type rfc3986Validator struct {
}

// Description returns a plain text description of the validator's behavior, suitable for a practitioner to understand its impact.
func (v rfc3986Validator) Description(ctx context.Context) string {
	return "string must be in RFC 3986 format"
}

// MarkdownDescription returns a markdown formatted description of the validator's behavior, suitable for a practitioner to understand its impact.
func (v rfc3986Validator) MarkdownDescription(ctx context.Context) string {
	return "string must be in RFC 3986 format"
}

// ValidateString runs the main validation logic of the validator, reading configuration data out of `req` and updating `resp` with diagnostics.
func (v rfc3986Validator) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	// If the value is unknown or null, there is nothing to validate.
	if req.ConfigValue.IsUnknown() || req.ConfigValue.IsNull() {
		return
	}

	_, err := url.ParseRequestURI(req.ConfigValue.ValueString())

	if err != nil {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Invalid String Format",
			fmt.Sprintf("String must be in RFC 3986 format: %s", err),
		)
		return
	}
}
