package hermes

import (
	"net/url"
	"time"

	"github.com/calbera/go-pyth-client/feeds"
	"github.com/calbera/go-pyth-client/types"
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
		return types.ErrInvalidHTTPTimeout
	}
	if c.MaxRetries < 0 {
		return types.ErrInvalidMaxRetries
	}

	if c.feedVersion, err = feeds.ToVersion(c.Ecosystem); err != nil {
		return err
	}

	return nil
}
