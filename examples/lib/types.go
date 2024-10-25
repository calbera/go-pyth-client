package lib

import (
	"math"
	"math/big"
)

// Prices are represented by int64s.
type Price int64

// The min/max possible price of an order.
const (
	MinPrice Price = 0
	MaxPrice Price = math.MaxInt64
)

// PriceCmp implements a standard comparator function for the Price type.
func PriceCmp(a, b Price) int {
	if a > b {
		return 1
	} else if a == b {
		return 0
	} else {
		return -1
	}
}

// PriceCmpUnsafe implements a standard comparator function for the Price type.
//
// NOTE: unsafe from under/overflows. Ensure the inputs are valid Price [MinPrice, MaxPrice].
func PriceCmpUnsafe(a, b Price) int {
	return int(a - b)
}

// PriceUpdate represents a price update received by the oracle.
type PriceUpdate struct {
	PairIndex      uint64
	Price          Price
	Conf           Price
	TimeStamp      uint64
	PythUpdateData [][]byte
	PythUpdateFee  *big.Int
}

// PriceUpdates is a map of price updates, keyed by pair index.
type PriceUpdates map[uint64]*PriceUpdate

// FeedOption is an enum for the different types of price feeds that can be queried.
type FeedOption int

const (
	Feed_ZERO       FeedOption = iota // 0 => no price feed to query
	Feed_SINGULAR                     // 1 => only 1 price feed to query for the final price
	Feed_TRIANGULAR                   // 2 => use triangular arbitrage to get price from target
	// pair (at index 0) and quote pair (at index 1)
)
