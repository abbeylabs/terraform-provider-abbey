package demo

type Demo struct {
	Id        *int64  `json:"id,omitempty"`
	UserId    *string `json:"user_id,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
}

func (d *Demo) SetId(id int64) {
	d.Id = &id
}

func (d *Demo) GetId() *int64 {
	if d == nil {
		return nil
	}
	return d.Id
}

func (d *Demo) SetUserId(userId string) {
	d.UserId = &userId
}

func (d *Demo) GetUserId() *string {
	if d == nil {
		return nil
	}
	return d.UserId
}

func (d *Demo) SetCreatedAt(createdAt string) {
	d.CreatedAt = &createdAt
}

func (d *Demo) GetCreatedAt() *string {
	if d == nil {
		return nil
	}
	return d.CreatedAt
}

func (d *Demo) SetUpdatedAt(updatedAt string) {
	d.UpdatedAt = &updatedAt
}

func (d *Demo) GetUpdatedAt() *string {
	if d == nil {
		return nil
	}
	return d.UpdatedAt
}
