package step

import (
	"abbey/v2/internal/shared/models/policy"
	"abbey/v2/internal/shared/models/reviewers"
)

type Step struct {
	Reviewers *reviewers.Reviewers `tfsdk:"reviewers"`
	SkipIf    []policy.Policy      `tfsdk:"skip_if"`
}
