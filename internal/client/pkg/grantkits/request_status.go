package grantkits

type RequestStatus string

const (
	REQUEST_STATUS_PENDING  RequestStatus = "Pending"
	REQUEST_STATUS_DENIED   RequestStatus = "Denied"
	REQUEST_STATUS_APPROVED RequestStatus = "Approved"
	REQUEST_STATUS_CANCELED RequestStatus = "Canceled"
)
