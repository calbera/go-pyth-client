package hermes

import (
	"strings"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/hashicorp/go-retryablehttp"

	"github.com/calbera/go-pyth-client/bindings/apyth"
	"github.com/calbera/go-pyth-client/feeds"
	"github.com/calbera/go-pyth-client/types"
)

// Client is a client for the Pyth Hermes API (https://hermes.pyth.network/docs)
type Client struct {
	// Config for Pyth and HTTP calls.
	cfg *Config

	// HTTP client that handles retries with a default retry policy.
	client *retryablehttp.Client

	// The logger to handle logs
	logger retryablehttp.LeveledLogger

	// ABI of the Pyth contract, useful for (en/de)coding responses.
	pythABI *abi.ABI

	// The cached price feed from the `/v2/updates/price/stream` stream
	ssePriceCached *ssePriceData

	// The subscription of the `/v2/updates/price/stream` stream should only happen once
	subscribeOnce sync.Once
}

// NewClient creates a client for the Pyth Hermes API.
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

	// Build the ABI of the Pyth contract for (en/de)coding responses.
	var pythABI abi.ABI
	if err := pythABI.UnmarshalJSON([]byte(apyth.ContractMetaData.ABI)); err != nil {
		return nil, err
	}

	// Initialize the cached sse price data struct
	ssePrice := &ssePriceData{
		latestPrice: make(map[string]*types.LatestPriceData),
		ready:       make(chan struct{}),
	}

	return &Client{
		cfg:            cfg,
		client:         retryablehttp.NewClient(),
		logger:         logger,
		pythABI:        &pythABI,
		ssePriceCached: ssePrice,
	}, nil
}

// Shutdown gracefully shuts down the Pyth Hermes client.
func (c *Client) Shutdown() {
	c.client.HTTPClient.CloseIdleConnections()
}

// Builds the API endpoint for querying multiple feeds on `v2/updates/price/latest`.
func (c *Client) buildBatchURLLatestPrice(priceFeeds ...string) string {
	return c.buildBatchURL(latestPriceAPI, priceFeeds...)
}

// Builds the API endpoint for querying multiple feeds on `v2/updates/price/stream`.
func (c *Client) buildBatchURLStream(priceFeeds ...string) string {
	return c.buildBatchURL(priceStreamAPI, priceFeeds...)
}

// Builds the API endpoint for querying multiple feeds on `v2/updates/price/latest`.
func (c *Client) buildBatchURL(apiName string, priceFeeds ...string) string {
	// Batch the price feed IDs into a single query string.
	urlComponents := make([]string, len(priceFeeds)+2)
	urlComponents[0] = c.cfg.APIEndpoint
	urlComponents[1] = apiName
	for i, priceFeed := range priceFeeds {
		urlComponents[i+2] = "ids[]=" +
			feeds.MustGetPriceFeedID(c.cfg.feedVersion, priceFeed) + "&"
	}

	return strings.Join(urlComponents, "")
}
