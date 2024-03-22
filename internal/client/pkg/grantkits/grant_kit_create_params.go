package grantkits

type GrantKitCreateParams struct {
	Name        *string        `json:"name,omitempty" required:"true"`
	Description *string        `json:"description,omitempty" required:"true"`
	Workflow    *GrantWorkflow `json:"workflow,omitempty"`
	Policies    []Policy       `json:"policies,omitempty"`
	Output      *Output        `json:"output,omitempty" required:"true"`
}

func (g *GrantKitCreateParams) SetName(name string) {
	g.Name = &name
}

func (g *GrantKitCreateParams) GetName() *string {
	if g == nil {
		return nil
	}
	return g.Name
}

func (g *GrantKitCreateParams) SetDescription(description string) {
	g.Description = &description
}

func (g *GrantKitCreateParams) GetDescription() *string {
	if g == nil {
		return nil
	}
	return g.Description
}

func (g *GrantKitCreateParams) SetWorkflow(workflow GrantWorkflow) {
	g.Workflow = &workflow
}

func (g *GrantKitCreateParams) GetWorkflow() *GrantWorkflow {
	if g == nil {
		return nil
	}
	return g.Workflow
}

func (g *GrantKitCreateParams) SetPolicies(policies []Policy) {
	g.Policies = policies
}

func (g *GrantKitCreateParams) GetPolicies() []Policy {
	if g == nil {
		return nil
	}
	return g.Policies
}

func (g *GrantKitCreateParams) SetOutput(output Output) {
	g.Output = &output
}

func (g *GrantKitCreateParams) GetOutput() *Output {
	if g == nil {
		return nil
	}
	return g.Output
}
