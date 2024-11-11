package gontenttype

import (
	"github.com/costinmrr/gontenttype/types/csv"
	"github.com/costinmrr/gontenttype/types/json"
	"github.com/costinmrr/gontenttype/types/xml"
)

func GetContentType(content string) ContentType {
	err := json.IsJSON(content)
	if err == nil {
		return JSON
	}

	err = xml.IsXML(content)
	if err == nil {
		return XML
	}

	err = csv.IsCSV(content)
	if err == nil {
		return CSV
	}

	return Unsupported
}