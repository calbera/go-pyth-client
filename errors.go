package client

import "errors"

var (
	// ErrFeedNotSupported is returned when the feed version is not supported.
	ErrFeedNotSupported = errors.New("feed version not supported")
)
