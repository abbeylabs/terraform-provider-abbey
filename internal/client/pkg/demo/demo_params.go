package demo

type DemoParams struct {
	Permission *Permission `json:"permission,omitempty" required:"true"`
	Email      *string     `json:"email,omitempty" required:"true"`
}

func (d *DemoParams) SetPermission(permission Permission) {
	d.Permission = &permission
}

func (d *DemoParams) GetPermission() *Permission {
	if d == nil {
		return nil
	}
	return d.Permission
}

func (d *DemoParams) SetEmail(email string) {
	d.Email = &email
}

func (d *DemoParams) GetEmail() *string {
	if d == nil {
		return nil
	}
	return d.Email
}

type Permission string

const (
	PERMISSION_READ_WRITE Permission = "read_write"
)
