package parquet

import (
	"strings"

	"github.com/parquet-go/parquet-go"
)

// IsParquet checks if the content is a parquet file.
// Parquet files start with a specific byte sequence known as the "magic number."
// The first four bytes are PAR1, and the last four bytes are also PAR1.
// We use the parquet-go library to validate the magic sequence and the footer.
// The footer contains the footer length, the file metadata (schema definition, row group information, and the column
// metadata), and the magic sequence.
// This should be enough to determine if the content is a parquet file.
func IsParquet(content string) error {
	_, err := parquet.OpenFile(strings.NewReader(content), int64(len(content)))
	if err != nil {
		return overrideError(err)
	}

	return nil
}
