package jsonlines

import (
	"errors"
	"fmt"
)

var ErrEmptyContent = errors.New("content is empty")

func errorOnLine(lineNo int, err error) error {
	return errors.New("error on line " + fmt.Sprintf("%d", lineNo) + ": " + err.Error())
}
