package requestable

import "encoding/json"

type enum struct {
	Type  string          `json:"type"`
	Value json.RawMessage `json:"value"`
}
