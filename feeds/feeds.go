package feeds

import "github.com/calbera/go-pyth-client/types"

// Enum for Pyth supported price feed ecosystems.
type Version int8

const (
	InvalidFeed Version = iota - 1
	EVMStable
	EVMBeta
	// ...
)

// ToVersion converts a ecosystem string to a Version enum.
func ToVersion(ecosystem string) (Version, error) {
	switch ecosystem {
	case "EVM-Stable", "EVM Stable", "evm stable", "evm-stable":
		return EVMStable, nil
	case "EVM-Beta", "EVM Beta", "evm beta", "evm-beta":
		return EVMBeta, nil
	default:
		return InvalidFeed, types.ErrFeedNotSupported
	}
}

// MustGetPriceFeedID converts a ticker string to a price feed ID string. Panics if the ticker is
// not supported.
func MustGetPriceFeedID(version Version, ticker string) string {
	switch version {
	case EVMStable:
		return EVMStableFeedsToIDs[ticker]
	case EVMBeta:
		return EVMBetaFeedsToIDs[ticker]
	default:
		panic(types.ErrFeedNotSupported)
	}
}

// MustGetPriceFeed converts a price feed ID string to a ticker string. Panics if the price feed ID
// is not supported.
func MustGetPriceFeed(version Version, id string) string {
	hexID := AddHexPrefix(id)
	switch version {
	case EVMStable:
		return EVMStableIDsToFeeds[hexID]
	case EVMBeta:
		return EVMBetaIDsToFeeds[hexID]
	default:
		panic(types.ErrFeedNotSupported)
	}
}
