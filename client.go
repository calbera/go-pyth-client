package client

import (
	"context"

	"github.com/calbera/go-pyth-client/hermes"
	"github.com/calbera/go-pyth-client/types"
)

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

var NewHermes = hermes.NewClient
