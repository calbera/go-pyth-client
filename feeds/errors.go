package feeds

import "errors"

// ErrFeedNotSupported is returned when the feed version is not supported.
var ErrFeedNotSupported = errors.New("feed version not supported")
