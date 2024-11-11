package xml

import "errors"

var (
	ErrEmptyContent     = errors.New("empty content")
	ErrSecondRootFound  = errors.New("found a second root element")
	ErrContentAfterRoot = errors.New("found content after root element was closed")
	ErrRootNotFound     = errors.New("root element not found")
)
