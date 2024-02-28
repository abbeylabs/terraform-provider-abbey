package grantkits

type Review struct {
	Id                *string       `json:"id,omitempty" required:"true"`
	UserId            *string       `json:"user_id,omitempty" required:"true"`
	UserEmail         *string       `json:"user_email,omitempty"`
	RequestId         *string       `json:"request_id,omitempty" required:"true"`
	Status            *ReviewStatus `json:"status,omitempty" required:"true"`
	RequestReason     *string       `json:"request_reason,omitempty" required:"true"`
	Reason            *string       `json:"reason,omitempty" required:"true"`
	GrantKitVersionId *string       `json:"grant_kit_version_id,omitempty" required:"true"`
	GrantKitName      *string       `json:"grant_kit_name,omitempty" required:"true"`
	GrantId           *string       `json:"grant_id,omitempty" required:"true"`
	Grant             *Grant        `json:"grant,omitempty"`
	CreatedAt         *string       `json:"created_at,omitempty" required:"true"`
	UpdatedAt         *string       `json:"updated_at,omitempty" required:"true"`
	PullRequest       *string       `json:"pull_request,omitempty" required:"true"`
}

func (r *Review) SetId(id string) {
	r.Id = &id
}

func (r *Review) GetId() *string {
	if r == nil {
		return nil
	}
	return r.Id
}

func (r *Review) SetUserId(userId string) {
	r.UserId = &userId
}

func (r *Review) GetUserId() *string {
	if r == nil {
		return nil
	}
	return r.UserId
}

func (r *Review) SetUserEmail(userEmail string) {
	r.UserEmail = &userEmail
}

func (r *Review) GetUserEmail() *string {
	if r == nil {
		return nil
	}
	return r.UserEmail
}

func (r *Review) SetRequestId(requestId string) {
	r.RequestId = &requestId
}

func (r *Review) GetRequestId() *string {
	if r == nil {
		return nil
	}
	return r.RequestId
}

func (r *Review) SetStatus(status ReviewStatus) {
	r.Status = &status
}

func (r *Review) GetStatus() *ReviewStatus {
	if r == nil {
		return nil
	}
	return r.Status
}

func (r *Review) SetRequestReason(requestReason string) {
	r.RequestReason = &requestReason
}

func (r *Review) GetRequestReason() *string {
	if r == nil {
		return nil
	}
	return r.RequestReason
}

func (r *Review) SetReason(reason string) {
	r.Reason = &reason
}

func (r *Review) GetReason() *string {
	if r == nil {
		return nil
	}
	return r.Reason
}

func (r *Review) SetGrantKitVersionId(grantKitVersionId string) {
	r.GrantKitVersionId = &grantKitVersionId
}

func (r *Review) GetGrantKitVersionId() *string {
	if r == nil {
		return nil
	}
	return r.GrantKitVersionId
}

func (r *Review) SetGrantKitName(grantKitName string) {
	r.GrantKitName = &grantKitName
}

func (r *Review) GetGrantKitName() *string {
	if r == nil {
		return nil
	}
	return r.GrantKitName
}

func (r *Review) SetGrantId(grantId string) {
	r.GrantId = &grantId
}

func (r *Review) GetGrantId() *string {
	if r == nil {
		return nil
	}
	return r.GrantId
}

func (r *Review) SetGrant(grant Grant) {
	r.Grant = &grant
}

func (r *Review) GetGrant() *Grant {
	if r == nil {
		return nil
	}
	return r.Grant
}

func (r *Review) SetCreatedAt(createdAt string) {
	r.CreatedAt = &createdAt
}

func (r *Review) GetCreatedAt() *string {
	if r == nil {
		return nil
	}
	return r.CreatedAt
}

func (r *Review) SetUpdatedAt(updatedAt string) {
	r.UpdatedAt = &updatedAt
}

func (r *Review) GetUpdatedAt() *string {
	if r == nil {
		return nil
	}
	return r.UpdatedAt
}

func (r *Review) SetPullRequest(pullRequest string) {
	r.PullRequest = &pullRequest
}

func (r *Review) GetPullRequest() *string {
	if r == nil {
		return nil
	}
	return r.PullRequest
}
