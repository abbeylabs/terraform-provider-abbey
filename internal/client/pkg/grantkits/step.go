package grantkits

type Step struct {
	Reviewers *Reviewers `json:"reviewers,omitempty"`
	SkipIf    []Policy   `json:"skip_if,omitempty"`
}

func (s *Step) SetReviewers(reviewers Reviewers) {
	s.Reviewers = &reviewers
}

func (s *Step) GetReviewers() *Reviewers {
	if s == nil {
		return nil
	}
	return s.Reviewers
}

func (s *Step) SetSkipIf(skipIf []Policy) {
	s.SkipIf = skipIf
}

func (s *Step) GetSkipIf() []Policy {
	if s == nil {
		return nil
	}
	return s.SkipIf
}
