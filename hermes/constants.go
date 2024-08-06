package hermes

// Pyth contract methods used for (un)packing price updates.
const (
	queryPriceFeed = `queryPriceFeed`
)

// API endpoints used.
const (
	latestPriceAPI = "/v2/updates/price/latest?"
	priceStreamAPI = "/v2/updates/price/stream?"
)
