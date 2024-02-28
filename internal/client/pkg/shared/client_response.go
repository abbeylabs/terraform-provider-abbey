package shared

type ClientResponse[T any] struct {
	Data     T
	Metadata ClientResponseMetadata
}

type ClientResponseMetadata struct {
	Headers    map[string]string
	StatusCode int
}
