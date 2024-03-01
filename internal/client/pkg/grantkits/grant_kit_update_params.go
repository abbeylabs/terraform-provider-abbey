package grantkits

type GrantKitUpdateParams struct {
	// The name of the connection
	Name        *string        `json:"name,omitempty" required:"true"`
	Description *string        `json:"description,omitempty" required:"true"`
	Workflow    *GrantWorkflow `json:"workflow,omitempty"`
	Output      *Output        `json:"output,omitempty" required:"true"`
	Policies    []Policy       `json:"policies,omitempty"`
}

func (g *GrantKitUpdateParams) SetName(name string) {
	g.Name = &name
}

func (g *GrantKitUpdateParams) GetName() *string {
	if g == nil {
		return nil
	}
	return g.Name
}

func (g *GrantKitUpdateParams) SetDescription(description string) {
	g.Description = &description
}

func (g *GrantKitUpdateParams) GetDescription() *string {
	if g == nil {
		return nil
	}
	return g.Description
}

func (g *GrantKitUpdateParams) SetWorkflow(workflow GrantWorkflow) {
	g.Workflow = &workflow
}

func (g *GrantKitUpdateParams) GetWorkflow() *GrantWorkflow {
	if g == nil {
		return nil
	}
	return g.Workflow
}

func (g *GrantKitUpdateParams) SetOutput(output Output) {
	g.Output = &output
}

func (g *GrantKitUpdateParams) GetOutput() *Output {
	if g == nil {
		return nil
	}
	return g.Output
}

func (g *GrantKitUpdateParams) SetPolicies(policies []Policy) {
	g.Policies = policies
}

func (g *GrantKitUpdateParams) GetPolicies() []Policy {
	if g == nil {
		return nil
	}
	return g.Policies
}
