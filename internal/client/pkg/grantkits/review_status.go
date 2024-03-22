package grantkits

type ReviewStatus string

const (
	REVIEW_STATUS_PENDING  ReviewStatus = "Pending"
	REVIEW_STATUS_DENIED   ReviewStatus = "Denied"
	REVIEW_STATUS_APPROVED ReviewStatus = "Approved"
	REVIEW_STATUS_CANCELED ReviewStatus = "Canceled"
)
