package grantkits

type Reviewers struct {
	OneOf []string `json:"one_of,omitempty"`
	AllOf []string `json:"all_of,omitempty"`
}

func (r *Reviewers) SetOneOf(oneOf []string) {
	r.OneOf = oneOf
}

func (r *Reviewers) GetOneOf() []string {
	if r == nil {
		return nil
	}
	return r.OneOf
}

func (r *Reviewers) SetAllOf(allOf []string) {
	r.AllOf = allOf
}

func (r *Reviewers) GetAllOf() []string {
	if r == nil {
		return nil
	}
	return r.AllOf
}
