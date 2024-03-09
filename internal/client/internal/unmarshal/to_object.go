package unmarshal

import (
	"encoding/json"
)

func ToObject(source []byte, target any) error {
	err := json.Unmarshal(source, target)
	if err != nil {
		return err
	}
	return nil
}
