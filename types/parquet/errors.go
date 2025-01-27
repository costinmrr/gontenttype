package parquet

import "errors"

func overrideError(err error) error {
	return errors.New("invalid parquet file: " + err.Error())
}
