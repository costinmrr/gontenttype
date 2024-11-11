package csv

import (
	"encoding/csv"
	"strings"
)

// IsCSV returns true if the content is a CSV.
func IsCSV(content string) error {
	if content == "" {
		return ErrEmptyContent
	}
	reader := csv.NewReader(strings.NewReader(content))
	_, err := reader.ReadAll()
	if err != nil {
		return err
	}

	return nil
}
