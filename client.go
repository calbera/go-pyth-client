package client

import (
	"context"
	"time"

	"github.com/calbera/go-pyth-client/benchmarks"
	"github.com/calbera/go-pyth-client/bindings/apyth"
	"github.com/calbera/go-pyth-client/hermes"
	"github.com/calbera/go-pyth-client/types"
)

// Benchmarks is the interface that wraps the methods of the benchmarks.Client.
type Benchmarks interface {
	GetHistoricalPriceUpdatesSync(
		ctx context.Context, timestamp time.Time, priceFeeds ...string,
	) (map[string]*apyth.PythStructsPriceFeed, error)
	Shutdown()
}

// NewBenchmarks creates a new Pyth Benchmarks client.
var NewBenchmarks = benchmarks.NewClient

// Hermes is the interface that wraps the methods of the hermes.Client.
type Hermes interface {
	GetCachedLatestPriceUpdates(
		ctx context.Context, priceFeeds ...string,
	) (map[string]*types.LatestPriceData, error)
	GetLatestPriceUpdatesAsync(
		ctx context.Context, priceFeeds ...string,
	) (map[string]*types.LatestPriceData, error)
	GetLatestPriceUpdatesSync(
		ctx context.Context, priceFeeds ...string,
	) (map[string]*types.LatestPriceData, error)
	SubscribePriceStreaming(ctx context.Context, priceFeeds ...string)
	Shutdown()
}

// NewHermes creates a new Pyth Hermes client.
var NewHermes = hermes.NewClient
