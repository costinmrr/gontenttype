package xml

import (
	"bytes"
	"encoding/xml"
	"io"
	"strings"
)

// IsXML returns true if the content is an XML.
func IsXML(content string) error {
	// If the content is empty, it is not an XML.
	content = strings.TrimSpace(content)
	if content == "" {
		return ErrEmptyContent
	}

	decoder := xml.NewDecoder(bytes.NewReader([]byte(content)))
	var rootFound bool
	var depth int
	for {
		tok, err := decoder.Token()
		if err == io.EOF {
			// Reached end of input with no extra content
			break
		}
		if err != nil {
			// XML parsing error
			return err
		}
		// Check for extra content after XML root
		switch tok.(type) {
		case xml.StartElement:
			if !rootFound {
				// Found the root element
				rootFound = true
			}
			// Increase depth
			depth++
		case xml.EndElement:
			// Decrease depth
			depth--
			// If depth is 0, the root element was closed. Check for extra content
			if depth == 0 {
				if _, err := decoder.Token(); err != io.EOF {
					// Found content after root element was closed
					return ErrContentAfterRoot
				}
				// Valid XML with one root element and no extra content
				return nil
			}
		}
	}

	if !rootFound {
		return ErrRootNotFound
	}

	return nil
}
