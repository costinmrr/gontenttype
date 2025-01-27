package gontenttype

type ContentType string

const (
	Unsupported ContentType = ""
	JSON        ContentType = "application/json"
	JSONLines   ContentType = "application/jsonl"
	XML         ContentType = "application/xml"
	CSV         ContentType = "text/csv"
	Parquet     ContentType = "application/vnd.apache.parquet"
)
