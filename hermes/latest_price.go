package hermes

import (
	"context"
	"encoding/json"
	"sync"

	"github.com/calbera/go-pyth-client/feeds"
	"github.com/calbera/go-pyth-client/types"
	"golang.org/x/sync/errgroup"
)

// Queries the `v2/updates/price/latest` endpoint for all price feed ID together. Takes the price
// feed keys (uses corresponding Pyth feed ID). Returns the Pyth PriceFeed struct and the price
// feed update data for each pair.
func (c *Client) GetLatestPriceUpdatesSync(
	_ context.Context, priceFeeds ...string,
) (map[string]*types.LatestPriceData, error) {
	// Validate parameters.
	if len(priceFeeds) == 0 {
		return nil, nil
	}

	// Build and fire the request.
	resp, err := c.client.Get(c.buildBatchURLLatestPrice(priceFeeds...))
	if err != nil {
		return nil, err
	}

	// Parse the response.
	var priceResp latestPriceResponse
	if err = json.NewDecoder(resp.Body).Decode(&priceResp); err != nil {
		return nil, err
	}

	latestPriceData := make(map[string]*types.LatestPriceData, len(priceResp.Parsed))
	err = c.resolveMany(&priceResp, latestPriceData)
	return latestPriceData, err
}

// Queries the `v2/updates/price/latest` endpoint for each price feed individually, in parallel.
// Takes the price feed keys (uses corresponding Pyth feed ID). Returns the Pyth PriceFeed struct
// and the price feed update data for each pair.
func (c *Client) GetLatestPriceUpdatesAsync(
	ctx context.Context, priceFeeds ...string,
) (map[string]*types.LatestPriceData, error) {
	// Validate parameters.
	if len(priceFeeds) == 0 {
		return nil, nil
	}

	// Initialize errgroup and results to run the requests in parallel.
	var (
		g, _    = errgroup.WithContext(ctx)
		results = sync.Map{}
	)

	// Fetch the price data results in parallel.
	g.SetLimit(len(priceFeeds))
	for _, priceFeed := range priceFeeds {
		g.Go(func() error {
			lpd, err := c.fetchIndividualPriceData(
				feeds.MustGetPriceFeedID(c.cfg.feedVersion, priceFeed),
			)
			if err != nil {
				return err
			}

			results.Store(priceFeed, lpd)

			return nil
		})
	}

	// Wait for all requests to finish.
	if err := g.Wait(); err != nil {
		return nil, err
	}

	// Resolve the results into the prices return map.
	prices := make(map[string]*types.LatestPriceData, len(priceFeeds))
	for _, priceFeed := range priceFeeds {
		if lpd, ok := results.Load(priceFeed); ok {
			//nolint:revive // is the only type in the map.
			prices[priceFeed] = lpd.(*types.LatestPriceData)
		}
	}
	return prices, nil
}

// A task for fetching price for each feed ID with the `v2/updates/price/latest` endpoint.
func (c *Client) fetchIndividualPriceData(feedID string) (*types.LatestPriceData, error) {
	// Build and fire the individual price request.
	url := c.cfg.APIEndpoint + "/v2/updates/price/latest?ids[]=" + feedID
	resp, err := c.client.Get(url)
	if err != nil {
		return nil, err
	}

	// Parse the response.
	var priceResp latestPriceResponse
	if err = json.NewDecoder(resp.Body).Decode(&priceResp); err != nil {
		return nil, err
	}
	return c.resolveOne(&priceResp)
}
