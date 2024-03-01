package grantkits

type Request struct {
	Id                *string        `json:"id,omitempty" required:"true"`
	GrantKitId        *string        `json:"grant_kit_id,omitempty" required:"true"`
	GrantKitVersionId *string        `json:"grant_kit_version_id,omitempty" required:"true"`
	GrantKitName      *string        `json:"grant_kit_name,omitempty"`
	Reason            *string        `json:"reason,omitempty" required:"true"`
	UserId            *string        `json:"user_id,omitempty" required:"true"`
	Status            *RequestStatus `json:"status,omitempty" required:"true"`
	Reviews           []Review       `json:"reviews,omitempty"`
	GrantId           *string        `json:"grant_id,omitempty" required:"true"`
	CreatedAt         *string        `json:"created_at,omitempty" required:"true"`
	UpdatedAt         *string        `json:"updated_at,omitempty" required:"true"`
	PullRequest       *string        `json:"pull_request,omitempty" required:"true"`
}

func (r *Request) SetId(id string) {
	r.Id = &id
}

func (r *Request) GetId() *string {
	if r == nil {
		return nil
	}
	return r.Id
}

func (r *Request) SetGrantKitId(grantKitId string) {
	r.GrantKitId = &grantKitId
}

func (r *Request) GetGrantKitId() *string {
	if r == nil {
		return nil
	}
	return r.GrantKitId
}

func (r *Request) SetGrantKitVersionId(grantKitVersionId string) {
	r.GrantKitVersionId = &grantKitVersionId
}

func (r *Request) GetGrantKitVersionId() *string {
	if r == nil {
		return nil
	}
	return r.GrantKitVersionId
}

func (r *Request) SetGrantKitName(grantKitName string) {
	r.GrantKitName = &grantKitName
}

func (r *Request) GetGrantKitName() *string {
	if r == nil {
		return nil
	}
	return r.GrantKitName
}

func (r *Request) SetReason(reason string) {
	r.Reason = &reason
}

func (r *Request) GetReason() *string {
	if r == nil {
		return nil
	}
	return r.Reason
}

func (r *Request) SetUserId(userId string) {
	r.UserId = &userId
}

func (r *Request) GetUserId() *string {
	if r == nil {
		return nil
	}
	return r.UserId
}

func (r *Request) SetStatus(status RequestStatus) {
	r.Status = &status
}

func (r *Request) GetStatus() *RequestStatus {
	if r == nil {
		return nil
	}
	return r.Status
}

func (r *Request) SetReviews(reviews []Review) {
	r.Reviews = reviews
}

func (r *Request) GetReviews() []Review {
	if r == nil {
		return nil
	}
	return r.Reviews
}

func (r *Request) SetGrantId(grantId string) {
	r.GrantId = &grantId
}

func (r *Request) GetGrantId() *string {
	if r == nil {
		return nil
	}
	return r.GrantId
}

func (r *Request) SetCreatedAt(createdAt string) {
	r.CreatedAt = &createdAt
}

func (r *Request) GetCreatedAt() *string {
	if r == nil {
		return nil
	}
	return r.CreatedAt
}

func (r *Request) SetUpdatedAt(updatedAt string) {
	r.UpdatedAt = &updatedAt
}

func (r *Request) GetUpdatedAt() *string {
	if r == nil {
		return nil
	}
	return r.UpdatedAt
}

func (r *Request) SetPullRequest(pullRequest string) {
	r.PullRequest = &pullRequest
}

func (r *Request) GetPullRequest() *string {
	if r == nil {
		return nil
	}
	return r.PullRequest
}
