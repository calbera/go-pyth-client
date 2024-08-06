package benchmarks

import (
	"net/url"
	"time"

	"github.com/calbera/go-pyth-client/types"
)

type Config struct {
	APIEndpoint string
	HTTPTimeout time.Duration
	MaxRetries  int
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

	return nil
}
