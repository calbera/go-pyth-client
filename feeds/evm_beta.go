package feeds

// Pyth EVM Beta Price Feeds, reference: https://pyth.network/developers/price-feed-ids
//
// NOTE: To add more feeds, please follow the pattern in evm_stable.go.

var EVMBetaFeedsToIDs = map[string]string{
	"BTC/USD":  "f9c0172ba10dfa4d19088d94f5bf61d3b54d5bd7483a322a982e1373ee8ea31b",
	"ETH/USD":  "ca80ba6dc32e08d06f1aa886011eed1d77c77be9eb761cc10d72b7d0a2fd57a6",
	"ATOM/USD": "61226d39beea19d334f17c2febce27e12646d84675924ebb02b9cdaea68727e3",
	"TIA/USD":  "30998099c161c4f04b48569485cfab66256f36006810fe787f40fbc96e263438",
	"USDC/USD": "41f3625971ca2ed2263e78573fe5ce23e13d2558ed3f2e47ab0f84fb9e7ae722",
	// ...
}

var EVMBetaIDsToFeeds = map[string]string{
	"f9c0172ba10dfa4d19088d94f5bf61d3b54d5bd7483a322a982e1373ee8ea31b": "BTC/USD",
	"ca80ba6dc32e08d06f1aa886011eed1d77c77be9eb761cc10d72b7d0a2fd57a6": "ETH/USD",
	"61226d39beea19d334f17c2febce27e12646d84675924ebb02b9cdaea68727e3": "ATOM/USD",
	"30998099c161c4f04b48569485cfab66256f36006810fe787f40fbc96e263438": "TIA/USD",
	"41f3625971ca2ed2263e78573fe5ce23e13d2558ed3f2e47ab0f84fb9e7ae722": "USDC/USD",
	// ...
}
