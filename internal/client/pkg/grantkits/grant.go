package grantkits

type Grant struct {
	Id                *string `json:"id,omitempty" required:"true"`
	GrantKitId        *string `json:"grant_kit_id,omitempty" required:"true"`
	GrantKitVersionId *string `json:"grant_kit_version_id,omitempty" required:"true"`
	UserId            *string `json:"user_id,omitempty" required:"true"`
	RequestId         *string `json:"request_id,omitempty" required:"true"`
	OrganizationId    *string `json:"organization_id,omitempty" required:"true"`
	Deleted           *bool   `json:"deleted,omitempty" required:"true"`
	CreatedAt         *string `json:"created_at,omitempty" required:"true"`
	UpdatedAt         *string `json:"updated_at,omitempty" required:"true"`
}

func (g *Grant) SetId(id string) {
	g.Id = &id
}

func (g *Grant) GetId() *string {
	if g == nil {
		return nil
	}
	return g.Id
}

func (g *Grant) SetGrantKitId(grantKitId string) {
	g.GrantKitId = &grantKitId
}

func (g *Grant) GetGrantKitId() *string {
	if g == nil {
		return nil
	}
	return g.GrantKitId
}

func (g *Grant) SetGrantKitVersionId(grantKitVersionId string) {
	g.GrantKitVersionId = &grantKitVersionId
}

func (g *Grant) GetGrantKitVersionId() *string {
	if g == nil {
		return nil
	}
	return g.GrantKitVersionId
}

func (g *Grant) SetUserId(userId string) {
	g.UserId = &userId
}

func (g *Grant) GetUserId() *string {
	if g == nil {
		return nil
	}
	return g.UserId
}

func (g *Grant) SetRequestId(requestId string) {
	g.RequestId = &requestId
}

func (g *Grant) GetRequestId() *string {
	if g == nil {
		return nil
	}
	return g.RequestId
}

func (g *Grant) SetOrganizationId(organizationId string) {
	g.OrganizationId = &organizationId
}

func (g *Grant) GetOrganizationId() *string {
	if g == nil {
		return nil
	}
	return g.OrganizationId
}

func (g *Grant) SetDeleted(deleted bool) {
	g.Deleted = &deleted
}

func (g *Grant) GetDeleted() *bool {
	if g == nil {
		return nil
	}
	return g.Deleted
}

func (g *Grant) SetCreatedAt(createdAt string) {
	g.CreatedAt = &createdAt
}

func (g *Grant) GetCreatedAt() *string {
	if g == nil {
		return nil
	}
	return g.CreatedAt
}

func (g *Grant) SetUpdatedAt(updatedAt string) {
	g.UpdatedAt = &updatedAt
}

func (g *Grant) GetUpdatedAt() *string {
	if g == nil {
		return nil
	}
	return g.UpdatedAt
}
