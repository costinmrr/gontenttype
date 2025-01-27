package jsonlines

import (
	"bufio"
	"encoding/json"
	"strings"
)

// IsJSONLines returns true if the content is a JSON Lines.
// JSON Lines is a text format where each line is a valid JSON object. Read more at https://jsonlines.org.
func IsJSONLines(content string) error {
	if content == "" {
		return ErrEmptyContent
	}

	// Scan the content line by line, and try to unmarshal it. If it fails, it's not a JSON Lines content.
	// Return an error with the line number and the error message as soon as an error is encountered.
	scanner := bufio.NewScanner(strings.NewReader(content))
	lineNo := 1
	for scanner.Scan() {
		line := scanner.Text()
		err := json.Unmarshal([]byte(line), new(interface{}))
		if err != nil {
			return errorOnLine(lineNo, err)
		}
		lineNo++
	}

	return nil
}
