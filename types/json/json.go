package json

import "encoding/json"

// IsJSON returns true if the content is a JSON.
func IsJSON(content string) error {
	err := json.Unmarshal([]byte(content), new(interface{}))
	if err != nil {
		return err
	}

	return nil
}
