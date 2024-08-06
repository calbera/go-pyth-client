package types

import (
	"github.com/calbera/go-pyth-client/bindings/apyth"
)

// Actionable data returned from the `v2/updates/price/latest` endpoint for 1 price feed ID.
type LatestPriceData struct {
	PriceFeed *apyth.PythStructsPriceFeed

	// Always contains only 1 []byte for this price feed's update data.
	UpdateData []byte
}
