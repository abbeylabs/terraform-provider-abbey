package demo

type GetDemoRequestParams struct {
	Email *string `queryParam:"email" required:"true"`
}

func (params *GetDemoRequestParams) SetEmail(email string) {
	params.Email = &email
}
