package types

import "errors"

// Config.
var (
	// ErrInvalidHTTPTimeout is returned when the HTTP timeout is not greater than 0.
	ErrInvalidHTTPTimeout = errors.New("http timeout must be greater than 0")

	// ErrInvalidMaxRetries is returned when the max retries is less than 0.
	ErrInvalidMaxRetries = errors.New("max retries must be greater than or equal to 0")
)

// Feeds.
var (
	// ErrFeedNotSupported is returned when the feed version is not supported.
	ErrFeedNotSupported = errors.New("feed version not supported")
)
