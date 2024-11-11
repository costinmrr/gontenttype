package gontenttype

type ContentType string

const (
	Unsupported ContentType = ""
	JSON        ContentType = "application/json"
	XML         ContentType = "application/xml"
	CSV         ContentType = "text/csv"
)
