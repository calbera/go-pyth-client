package benchmarks

import (
	"github.com/hashicorp/go-retryablehttp"
)

// Client is a client for the Pyth Benchmarks API (https://benchmarks.pyth.network/docs)
type Client struct {
	// Config for Pyth and HTTP calls.
	cfg *Config

	// HTTP client that handles retries with a default retry policy.
	client *retryablehttp.Client

	// The logger to handle logs
	logger retryablehttp.LeveledLogger
}

// NewClient creates a client for the Pyth Benchmarks API.
func NewClient(cfg *Config, logger retryablehttp.LeveledLogger) (*Client, error) {
	// Ensure the given configuration is valid.
	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	// Setup and configure the retryable HTTP client.
	httpClient := retryablehttp.NewClient()
	httpClient.HTTPClient.Timeout = cfg.HTTPTimeout
	httpClient.Logger = logger
	httpClient.RetryMax = cfg.MaxRetries

	return &Client{
		cfg:    cfg,
		client: httpClient,
		logger: logger,
	}, nil
}

// Shutdown gracefully shuts down the Pythnet Hermes client.
func (c *Client) Shutdown() {
	c.client.HTTPClient.CloseIdleConnections()
}
