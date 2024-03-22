package grant_workflow

import (
	"abbey/v2/internal/shared/models/step"
)

type GrantWorkflow struct {
	Steps []step.Step `tfsdk:"steps"`
}
