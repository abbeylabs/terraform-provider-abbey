package grantkits

type GrantWorkflow struct {
	Steps []Step `json:"steps,omitempty"`
}

func (g *GrantWorkflow) SetSteps(steps []Step) {
	g.Steps = steps
}

func (g *GrantWorkflow) GetSteps() []Step {
	if g == nil {
		return nil
	}
	return g.Steps
}
