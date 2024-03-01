package identities

type IdentityParams struct {
	AbbeyAccount *string `json:"abbey_account,omitempty" required:"true"`
	Source       *string `json:"source,omitempty" required:"true"`
	// Json encoded string. See documentation for details.
	Metadata *string `json:"metadata,omitempty" required:"true"`
}

func (i *IdentityParams) SetAbbeyAccount(abbeyAccount string) {
	i.AbbeyAccount = &abbeyAccount
}

func (i *IdentityParams) GetAbbeyAccount() *string {
	if i == nil {
		return nil
	}
	return i.AbbeyAccount
}

func (i *IdentityParams) SetSource(source string) {
	i.Source = &source
}

func (i *IdentityParams) GetSource() *string {
	if i == nil {
		return nil
	}
	return i.Source
}

func (i *IdentityParams) SetMetadata(metadata string) {
	i.Metadata = &metadata
}

func (i *IdentityParams) GetMetadata() *string {
	if i == nil {
		return nil
	}
	return i.Metadata
}
