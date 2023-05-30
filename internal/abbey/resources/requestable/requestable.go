package requestable

import (
	_ "embed"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/moznion/go-optional"

	. "abbey.io/terraform-provider-abbey/internal/abbey/entity"
)

type Model struct {
	Id       types.String `tfsdk:"id"`
	Name     types.String `tfsdk:"name"`
	Workflow WorkflowTf   `tfsdk:"workflow"`
	Grant    GrantTf      `tfsdk:"grant"`
}

func (m Model) ToView() *View {
	var (
		workflow *Workflow
		grant    *Grant
	)

	if !m.Workflow.IsNull() && !m.Grant.IsUnknown() {
		workflow = &m.Workflow.Workflow
	}

	if !m.Grant.IsNull() && !m.Grant.IsUnknown() {
		grant = &m.Grant.Grant
	}

	return &View{
		Id:          m.Id.ValueString(),
		Name:        m.Name.ValueString(),
		Description: "",
		Workflow:    workflow,
		Grant:       grant,
		Policies:    nil,
	}
}

type View struct {
	Id          string            `json:"id,omitempty"`
	Name        string            `json:"name,omitempty"`
	Description string            `json:"description"`
	Workflow    *Workflow         `json:"workflow,omitempty"`
	Grant       *Grant            `json:"grant,omitempty"`
	Policies    Option[PolicySet] `json:"policies,omitempty"`
}

func (v View) ToModel() *Model {
	var (
		workflow WorkflowTf
		grant    GrantTf
	)

	if v.Workflow != nil {
		workflow = NewWorkflow(*v.Workflow)
	}

	if v.Grant != nil {
		grant = NewGrant(*v.Grant)
	}

	return &Model{
		Id:       types.StringValue(v.Id),
		Name:     types.StringValue(v.Name),
		Workflow: workflow,
		Grant:    grant,
	}
}
