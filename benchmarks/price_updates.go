package benchmarks

import (
	"context"
	"encoding/json"
	"strconv"
	"strings"
	"time"

	"github.com/calbera/go-pyth-client/bindings/apyth"
	"github.com/calbera/go-pyth-client/feeds"
)

func (c *Client) GetHistoricalPriceUpdatesSync(
	_ context.Context, timestamp time.Time, priceFeeds ...string,
) (map[string]*apyth.PythStructsPriceFeed, error) {
	// Validate parameters.
	if len(priceFeeds) == 0 {
		return nil, nil
	}

	// Build and fire the request.
	resp, err := c.client.Get(c.buildBatchURL(timestamp, priceFeeds...))
	if err != nil {
		return nil, err
	}

	// Parse the response.
	var priceResp priceResponse
	if err = json.NewDecoder(resp.Body).Decode(&priceResp); err != nil {
		return nil, err
	}

	priceResults := make(map[string]*apyth.PythStructsPriceFeed, len(priceResp.Parsed))
	if err = resolveMany(&priceResp, priceResults); err != nil {
		return nil, err
	}
	return priceResults, nil
}

// Builds the API endpoint for querying multiple feeds on `v1/updates/price/{timestamp}`.
func (c *Client) buildBatchURL(timestamp time.Time, priceFeeds ...string) string {
	// Batch the price feed IDs into a single query string.
	urlComponents := make([]string, len(priceFeeds)+3)
	urlComponents[0] = c.cfg.APIEndpoint
	urlComponents[1] = priceUpdateAPI
	urlComponents[2] = strconv.FormatInt(timestamp.Unix(), 10) + "?"
	for i, priceFeed := range priceFeeds {
		urlComponents[i+3] = "ids=" +
			feeds.MustGetPriceFeedID(feeds.EVMStable, priceFeed) + "&"
	}
	return strings.Join(urlComponents, "")
}
