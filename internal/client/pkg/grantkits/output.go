package grantkits

type Output struct {
	Location  *string `json:"location,omitempty" required:"true"`
	Append    *string `json:"append,omitempty"`
	Overwrite *string `json:"overwrite,omitempty"`
}

func (o *Output) SetLocation(location string) {
	o.Location = &location
}

func (o *Output) GetLocation() *string {
	if o == nil {
		return nil
	}
	return o.Location
}

func (o *Output) SetAppend(append string) {
	o.Append = &append
}

func (o *Output) GetAppend() *string {
	if o == nil {
		return nil
	}
	return o.Append
}

func (o *Output) SetOverwrite(overwrite string) {
	o.Overwrite = &overwrite
}

func (o *Output) GetOverwrite() *string {
	if o == nil {
		return nil
	}
	return o.Overwrite
}
