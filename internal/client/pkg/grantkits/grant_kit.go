package grantkits

type GrantKit struct {
	Id                    *string        `json:"id,omitempty" required:"true"`
	Name                  *string        `json:"name,omitempty" required:"true"`
	CurrentVersionId      *string        `json:"current_version_id,omitempty" required:"true"`
	Description           *string        `json:"description,omitempty" required:"true"`
	MaxGrantDurationInSec *float64       `json:"max_grant_duration_in_sec,omitempty"`
	Workflow              *GrantWorkflow `json:"workflow,omitempty"`
	Policies              []Policy       `json:"policies,omitempty"`
	Output                *Output        `json:"output,omitempty" required:"true"`
	Grants                []Grant        `json:"grants,omitempty" required:"true"`
	ResourceType          *string        `json:"resource_type,omitempty" required:"true"`
	Requests              []Request      `json:"requests,omitempty" required:"true"`
	CreatedAt             *string        `json:"created_at,omitempty" required:"true"`
	UpdatedAt             *string        `json:"updated_at,omitempty" required:"true"`
}

func (g *GrantKit) SetId(id string) {
	g.Id = &id
}

func (g *GrantKit) GetId() *string {
	if g == nil {
		return nil
	}
	return g.Id
}

func (g *GrantKit) SetName(name string) {
	g.Name = &name
}

func (g *GrantKit) GetName() *string {
	if g == nil {
		return nil
	}
	return g.Name
}

func (g *GrantKit) SetCurrentVersionId(currentVersionId string) {
	g.CurrentVersionId = &currentVersionId
}

func (g *GrantKit) GetCurrentVersionId() *string {
	if g == nil {
		return nil
	}
	return g.CurrentVersionId
}

func (g *GrantKit) SetDescription(description string) {
	g.Description = &description
}

func (g *GrantKit) GetDescription() *string {
	if g == nil {
		return nil
	}
	return g.Description
}

func (g *GrantKit) SetMaxGrantDurationInSec(maxGrantDurationInSec float64) {
	g.MaxGrantDurationInSec = &maxGrantDurationInSec
}

func (g *GrantKit) GetMaxGrantDurationInSec() *float64 {
	if g == nil {
		return nil
	}
	return g.MaxGrantDurationInSec
}

func (g *GrantKit) SetWorkflow(workflow GrantWorkflow) {
	g.Workflow = &workflow
}

func (g *GrantKit) GetWorkflow() *GrantWorkflow {
	if g == nil {
		return nil
	}
	return g.Workflow
}

func (g *GrantKit) SetPolicies(policies []Policy) {
	g.Policies = policies
}

func (g *GrantKit) GetPolicies() []Policy {
	if g == nil {
		return nil
	}
	return g.Policies
}

func (g *GrantKit) SetOutput(output Output) {
	g.Output = &output
}

func (g *GrantKit) GetOutput() *Output {
	if g == nil {
		return nil
	}
	return g.Output
}

func (g *GrantKit) SetGrants(grants []Grant) {
	g.Grants = grants
}

func (g *GrantKit) GetGrants() []Grant {
	if g == nil {
		return nil
	}
	return g.Grants
}

func (g *GrantKit) SetResourceType(resourceType string) {
	g.ResourceType = &resourceType
}

func (g *GrantKit) GetResourceType() *string {
	if g == nil {
		return nil
	}
	return g.ResourceType
}

func (g *GrantKit) SetRequests(requests []Request) {
	g.Requests = requests
}

func (g *GrantKit) GetRequests() []Request {
	if g == nil {
		return nil
	}
	return g.Requests
}

func (g *GrantKit) SetCreatedAt(createdAt string) {
	g.CreatedAt = &createdAt
}

func (g *GrantKit) GetCreatedAt() *string {
	if g == nil {
		return nil
	}
	return g.CreatedAt
}

func (g *GrantKit) SetUpdatedAt(updatedAt string) {
	g.UpdatedAt = &updatedAt
}

func (g *GrantKit) GetUpdatedAt() *string {
	if g == nil {
		return nil
	}
	return g.UpdatedAt
}
