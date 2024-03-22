package identities

type Identity struct {
	Id           *string `json:"id,omitempty"`
	CreatedAt    *string `json:"created_at,omitempty"`
	UpdatedAt    *string `json:"updated_at,omitempty"`
	AbbeyAccount *string `json:"abbey_account,omitempty" required:"true"`
	Source       *string `json:"source,omitempty" required:"true"`
	// Json encoded string. See documentation for details.
	Metadata *string `json:"metadata,omitempty" required:"true"`
}

func (i *Identity) SetId(id string) {
	i.Id = &id
}

func (i *Identity) GetId() *string {
	if i == nil {
		return nil
	}
	return i.Id
}

func (i *Identity) SetCreatedAt(createdAt string) {
	i.CreatedAt = &createdAt
}

func (i *Identity) GetCreatedAt() *string {
	if i == nil {
		return nil
	}
	return i.CreatedAt
}

func (i *Identity) SetUpdatedAt(updatedAt string) {
	i.UpdatedAt = &updatedAt
}

func (i *Identity) GetUpdatedAt() *string {
	if i == nil {
		return nil
	}
	return i.UpdatedAt
}

func (i *Identity) SetAbbeyAccount(abbeyAccount string) {
	i.AbbeyAccount = &abbeyAccount
}

func (i *Identity) GetAbbeyAccount() *string {
	if i == nil {
		return nil
	}
	return i.AbbeyAccount
}

func (i *Identity) SetSource(source string) {
	i.Source = &source
}

func (i *Identity) GetSource() *string {
	if i == nil {
		return nil
	}
	return i.Source
}

func (i *Identity) SetMetadata(metadata string) {
	i.Metadata = &metadata
}

func (i *Identity) GetMetadata() *string {
	if i == nil {
		return nil
	}
	return i.Metadata
}
