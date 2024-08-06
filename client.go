package client

import (
	"context"
	"time"

	"github.com/calbera/go-pyth-client/benchmarks"
	"github.com/calbera/go-pyth-client/bindings/apyth"
	"github.com/calbera/go-pyth-client/hermes"
	"github.com/calbera/go-pyth-client/types"
)

// Benchmarks is the interface of the Benchmarks client.
type Benchmarks interface {
	// Queries the `v1/updates/price/{timestamp}` endpoint for all price feed IDs together. Takes
	// the price feed keys (uses corresponding Pyth feed ID). Returns the Pyth PriceFeed struct
	// and the price feed update data for each pair.
	GetHistoricalPriceUpdatesSync(
		ctx context.Context, timestamp time.Time, priceFeeds ...string,
	) (map[string]*apyth.PythStructsPriceFeed, error)

	// Shutdown gracefully shuts down the Pyth Benchmarks client.
	Shutdown()
}

// NewBenchmarks creates a new Pyth Benchmarks client.
var NewBenchmarks = benchmarks.NewClient

// Hermes is the interface of the Hermes client.
type Hermes interface {
	// Queries the `v2/updates/price/latest` endpoint for each price feed individually, in
	// parallel. Takes the price feed keys (uses corresponding Pyth feed ID). Returns the Pyth
	// PriceFeed struct and the price feed update data for each pair.
	GetLatestPriceUpdatesAsync(
		ctx context.Context, priceFeeds ...string,
	) (map[string]*types.LatestPriceData, error)

	// Queries the `v2/updates/price/latest` endpoint for all price feed ID together. Takes the
	// price feed keys (uses corresponding Pyth feed ID). Returns the Pyth PriceFeed struct and the
	// price feed update data for each pair.
	GetLatestPriceUpdatesSync(
		ctx context.Context, priceFeeds ...string,
	) (map[string]*types.LatestPriceData, error)

	// Subscribe price feed from the streaming `v2/updates/price/stream` endpoint. Ensures this
	// only happens once in the scope of runtime. Any further calls to this are unnecessary and
	// no-ops.
	SubscribePriceStreaming(ctx context.Context, priceFeeds ...string)

	// Queries cached price feed update data, obtained from the SSE streaming endpoint.
	// Returns the Pyth PriceFeed struct and the price feed update data for each pair.
	GetCachedLatestPriceUpdates(
		ctx context.Context, priceFeeds ...string,
	) (map[string]*types.LatestPriceData, error)

	// Gracefully shuts down the Pyth Hermes client.
	Shutdown()
}

// NewHermes creates a new Pyth Hermes client.
var NewHermes = hermes.NewClient
