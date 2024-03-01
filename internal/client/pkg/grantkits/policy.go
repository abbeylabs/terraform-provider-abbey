package grantkits

type Policy struct {
	Bundle *string `json:"bundle,omitempty"`
	Query  *string `json:"query,omitempty"`
}

func (p *Policy) SetBundle(bundle string) {
	p.Bundle = &bundle
}

func (p *Policy) GetBundle() *string {
	if p == nil {
		return nil
	}
	return p.Bundle
}

func (p *Policy) SetQuery(query string) {
	p.Query = &query
}

func (p *Policy) GetQuery() *string {
	if p == nil {
		return nil
	}
	return p.Query
}
