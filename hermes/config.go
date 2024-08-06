package hermes

import (
	"errors"
	"net/url"
	"time"

	"github.com/calbera/go-pyth-client/feeds"
)

var (
	// ErrInvalidHTTPTimeout is returned when the HTTP timeout is not greater than 0.
	ErrInvalidHTTPTimeout = errors.New("http timeout must be greater than 0")

	// ErrInvalidMaxRetries is returned when the max retries is less than 0.
	ErrInvalidMaxRetries = errors.New("max retries must be greater than or equal to 0")
)

type Config struct {
	// Offchain parameters
	APIEndpoint string
	HTTPTimeout time.Duration
	MaxRetries  int
	Ecosystem   string

	feedVersion feeds.Version

	// Onchain parameters
	UseMock bool // Uses the mock Pyth contract rather than the real one.
}

func (c *Config) Validate() error {
	_, err := url.Parse(c.APIEndpoint)
	if err != nil {
		return err
	}

	if c.HTTPTimeout <= 0 {
		return ErrInvalidHTTPTimeout
	}
	if c.MaxRetries < 0 {
		return ErrInvalidMaxRetries
	}

	if c.feedVersion, err = feeds.ToVersion(c.Ecosystem); err != nil {
		return err
	}

	return nil
}
