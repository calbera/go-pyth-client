package hermes

import (
	"sync"

	"github.com/calbera/go-pyth-client/types"
)

// Cached data returned from the `/v2/updates/price/stream` stream with mutex feature.
// Streaming data may return hundreds of ms after subscription, so the ready channel is used
// to block read operation until cached data gets populated.
type ssePriceData struct {
	mu            sync.RWMutex
	ready         chan struct{}
	broadcastOnce sync.Once
	latestPrice   map[string]*types.LatestPriceData
}

// JSON formatted price returned from the `v2/updates/price/latest` endpoint.
type price struct {
	Price       string `json:"price"`
	Conf        string `json:"conf"`
	Expo        int32  `json:"expo"`
	PublishTime int64  `json:"publish_time"`
}

// JSON response returned from the `v2/updates/price/latest` endpoint.
//
//nolint:revive // needed for JSON unmarshalling.
type latestPriceResponse struct {
	Binary struct {
		Data []string `json:"data"`
	} `json:"binary"`
	Parsed []struct {
		ID       string `json:"id"`
		Price    price  `json:"price"`
		EmaPrice price  `json:"ema_price"`
	} `json:"parsed"`
}
