package hermes

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"github.com/calbera/go-pyth-client/types"
	sse "github.com/r3labs/sse/v2"
)

// Retry parameters.
const (
	initialBackoff = 1 * time.Second
	maxRetries     = 3
)

// Subscribe price feed from the streaming `v2/updates/price/stream` endpoint. Ensures this only
// happens once in the scope of runtime. Any further calls to this are unnecessary and no-ops.
func (c *Client) SubscribePriceStreaming(ctx context.Context, priceFeeds ...string) {
	c.subscribeOnce.Do(func() {
		client := sse.NewClient(c.buildBatchURLStream(priceFeeds...))

		subscribe := func() error {
			return client.SubscribeRawWithContext(ctx, func(msg *sse.Event) {
				c.handleSseEvent(msg)
			})
		}

		// Subscribe to the SSE using the context and the channel
		// Use goroutine since SSE Subscribe will block the current thread
		// see https://github.com/r3labs/sse/blob/master/README.md
		go c.subscribeWithRetries(ctx, subscribe)
	})
}

// Queries cached price feed update data, obtained from sse streaming endpoints.
// Returns the Pyth PriceFeed struct and the price feed update data for each pair.
func (c *Client) GetCachedLatestPriceUpdates(
	ctx context.Context, priceFeeds ...string,
) (map[string]*types.LatestPriceData, error) {
	// Validate parameters
	if len(priceFeeds) == 0 {
		return nil, fmt.Errorf("zero length of price feeds is an invalid input")
	}

	// Wait for the ready signal
	if err := c.waitForReady(ctx); err != nil {
		return nil, err
	}

	c.ssePriceCached.mu.RLock()
	defer c.ssePriceCached.mu.RUnlock()

	cachedUpdates := make(map[string]*types.LatestPriceData)
	for _, priceFeed := range priceFeeds {
		if _, ok := c.ssePriceCached.latestPrice[priceFeed]; !ok {
			return nil, fmt.Errorf("this price feed has not been subscribed to: %s", priceFeed)
		}
		cachedUpdates[priceFeed] = c.ssePriceCached.latestPrice[priceFeed]
	}
	return cachedUpdates, nil
}

// Handler of the sse streaming event.
func (c *Client) handleSseEvent(event *sse.Event) {
	// Decode the price from sse response to LatestPriceData.
	var priceResp latestPriceResponse
	if err := json.Unmarshal(event.Data, &priceResp); err != nil {
		c.logger.Error(
			"skipping msg, encountered an error when unmarshalling streaming data", "error", err,
		)
		return
	}

	// Unpack the price data and store in the local cache.
	c.ssePriceCached.mu.Lock()
	if err := c.resolveSsePrice(&priceResp); err != nil {
		c.logger.Error(
			"encountered an error when decoding price response from sse stream", "error", err,
		)
	}
	c.ssePriceCached.mu.Unlock()

	// Close the channel to effectively broadcast the first update to consumers
	c.ssePriceCached.broadcastOnce.Do(func() {
		close(c.ssePriceCached.ready) // Signal that the first update has occurred
	})
}

// waitForReady waits for the ready signal. Once the channel is closed,
// reads from the channel are non-blocking and very fast, essentially becoming a no-op.
func (c *Client) waitForReady(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-c.ssePriceCached.ready:
			return nil
		default:
			// If the channel is not closed, wait for a short duration before checking again
			// to avoid busy waiting
			// choose 100ms since 1st price data usually comes <1s after subscription
			time.Sleep(100 * time.Millisecond)
		}
	}
}

// subscribeWithRetries retries a subscription up to maxRetries with exponential backoff.
// If the task fails after maxRetries, it panics to indicate a critical failure.
func (c *Client) subscribeWithRetries(ctx context.Context, subscribe func() error) {
	backoff := initialBackoff
	retries := 0

	for retries < maxRetries {
		select {
		case <-ctx.Done():
			c.logger.Error("context cancelled while trying to subscribe to SSE stream")
			return
		default:
			err := subscribe()
			if err == nil {
				return
			}

			retries++
			c.logger.Error(
				"encountered an error when subscribing to SSE stream, now retrying...",
				"error", err, "num_retries", retries,
			)
			if retries >= maxRetries {
				panic(fmt.Sprintf("failed to subscribe to SSE stream after %d attempts: %v",
					retries, err))
			}

			// #nosec:G404 // fine.
			time.Sleep(backoff + time.Duration(rand.Intn(1000))*time.Millisecond)
			backoff *= 2
		}
	}
}
