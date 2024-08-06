package feeds

// Enum for Pyth supported price feed ecosystems.
type Version int8

const (
	InvalidFeed Version = iota - 1
	EVMStable
	EVMBeta
	// ...
)

func ToVersion(ecosystem string) (Version, error) {
	switch ecosystem {
	case "EVM-Stable", "EVM Stable", "evm stable", "evm-stable":
		return EVMStable, nil
	case "EVM-Beta", "EVM Beta", "evm beta", "evm-beta":
		return EVMBeta, nil
	default:
		return InvalidFeed, ErrFeedNotSupported
	}
}

func MustGetPriceFeedID(version Version, ticker string) string {
	switch version {
	case EVMStable:
		return EVMStableFeedsToIDs[ticker]
	case EVMBeta:
		return EVMBetaFeedsToIDs[ticker]
	default:
		panic(ErrFeedNotSupported)
	}
}

func MustGetPriceFeed(version Version, id string) string {
	switch version {
	case EVMStable:
		return EVMStableIDsToFeeds[id]
	case EVMBeta:
		return EVMBetaIDsToFeeds[id]
	default:
		panic(ErrFeedNotSupported)
	}
}

// Reference: https://pyth.network/developers/price-feed-ids

// Pyth EVM Stable Price Feeds.
var (
	EVMStableFeedsToIDs = map[string]string{
		"BTC/USD":  "e62df6c8b4a85fe1a67db44dc12de5db330f7ac66b72dc658afedf0f4a415b43",
		"ETH/USD":  "ff61491a931112ddf1bd8147cd1b641375f79f5825126d665480874634fd0ace",
		"ATOM/USD": "b00b60f88b03a6a625a8d1c048c3f66653edf217439983d037e7222c4e612819",
		"TIA/USD":  "09f7c1d7dfbb7df2b8fe3d3d87ee94a2259d212da4f30c1f0540d066dfa44723",
		"USDC/USD": "eaa020c61cc479712813461ce153894a96a6c00b21ed0cfc2798d1f9a9e9c94a",
		// ...
	}

	EVMStableIDsToFeeds = map[string]string{
		"e62df6c8b4a85fe1a67db44dc12de5db330f7ac66b72dc658afedf0f4a415b43": "BTC/USD",
		"ff61491a931112ddf1bd8147cd1b641375f79f5825126d665480874634fd0ace": "ETH/USD",
		"b00b60f88b03a6a625a8d1c048c3f66653edf217439983d037e7222c4e612819": "ATOM/USD",
		"09f7c1d7dfbb7df2b8fe3d3d87ee94a2259d212da4f30c1f0540d066dfa44723": "TIA/USD",
		"eaa020c61cc479712813461ce153894a96a6c00b21ed0cfc2798d1f9a9e9c94a": "USDC/USD",
		// ...
	}
)

// Pyth EVM Beta Price Feeds.
var (
	EVMBetaFeedsToIDs = map[string]string{
		"BTC/USD":  "f9c0172ba10dfa4d19088d94f5bf61d3b54d5bd7483a322a982e1373ee8ea31b",
		"ETH/USD":  "ca80ba6dc32e08d06f1aa886011eed1d77c77be9eb761cc10d72b7d0a2fd57a6",
		"ATOM/USD": "61226d39beea19d334f17c2febce27e12646d84675924ebb02b9cdaea68727e3",
		"TIA/USD":  "30998099c161c4f04b48569485cfab66256f36006810fe787f40fbc96e263438",
		"USDC/USD": "41f3625971ca2ed2263e78573fe5ce23e13d2558ed3f2e47ab0f84fb9e7ae722",
		// ...
	}

	EVMBetaIDsToFeeds = map[string]string{
		"f9c0172ba10dfa4d19088d94f5bf61d3b54d5bd7483a322a982e1373ee8ea31b": "BTC/USD",
		"ca80ba6dc32e08d06f1aa886011eed1d77c77be9eb761cc10d72b7d0a2fd57a6": "ETH/USD",
		"61226d39beea19d334f17c2febce27e12646d84675924ebb02b9cdaea68727e3": "ATOM/USD",
		"30998099c161c4f04b48569485cfab66256f36006810fe787f40fbc96e263438": "TIA/USD",
		"41f3625971ca2ed2263e78573fe5ce23e13d2558ed3f2e47ab0f84fb9e7ae722": "USDC/USD",
		// ...
	}
)
