package query

import (
	"context"
	"fmt"
	"time"

	client "github.com/calbera/go-pyth-client"
	"github.com/calbera/go-pyth-client/examples/lib"
)

// Requires a Pyth client to fetch prices from Pyth Benchmarks API. Assumes that all required
// feeds to be queried for `oracleFeeds` are in the `uniqueFeeds` slice.
func GetAllPricesAt(
	ctx context.Context, timestamp time.Time, pythClient client.Benchmarks, qs *lib.QuerySettings,
	pairIndexes map[string]uint64, oracleFeeds map[string][]string, uniqueFeeds []string,
) (lib.PriceUpdates, error) {
	// Query Pyth for the unique price feeds necessary
	priceData, err := pythClient.GetHistoricalPriceUpdatesSync(ctx, timestamp, uniqueFeeds...)
	if err != nil {
		return nil, err
	}

	var (
		pairIndex     uint64
		allPairPrices = make(lib.PriceUpdates, len(oracleFeeds))
	)

	for pair, priceFeeds := range oracleFeeds {
		if pairIndexes != nil {
			pairIndex = pairIndexes[pair]
		} else {
			pairIndex++ // use an arbitrary index if none provided
		}

		switch lib.FeedOption(len(priceFeeds)) {
		case lib.Feed_ZERO:
			continue

		case lib.Feed_SINGULAR:
			allPairPrices[pairIndex] = lib.GetPriceUpdateFromPythStructsPriceFeed(
				priceData[priceFeeds[0]], pairIndex, qs.UseEma, qs.DesiredPrecision,
			)

		case lib.Feed_TRIANGULAR:
			baseP := lib.GetPriceUpdateFromPythStructsPriceFeed(
				priceData[priceFeeds[0]], pairIndex, qs.UseEma, qs.DesiredPrecision,
			)
			quoteP := lib.GetPriceUpdateFromPythStructsPriceFeed(
				priceData[priceFeeds[1]], pairIndex, qs.UseEma, qs.DesiredPrecision,
			)

			// Use the older of the two as the timestamp.
			var timestamp uint64
			if baseP.TimeStamp <= quoteP.TimeStamp {
				timestamp = baseP.TimeStamp
			} else {
				timestamp = quoteP.TimeStamp
			}

			allPairPrices[pairIndex] = &lib.PriceUpdate{
				PairIndex: pairIndex,
				Price: lib.Price(lib.CalculateTriangularPrice(
					int64(baseP.Price), int64(quoteP.Price), qs.DesiredPrecision,
				)),
				Conf: lib.Price(lib.CalculateTriangularConf(
					int64(baseP.Price), int64(quoteP.Price),
					int64(baseP.Conf), int64(quoteP.Conf),
					qs.DesiredPrecision,
				)),
				TimeStamp: timestamp,
			}

		default:
			return nil, fmt.Errorf(
				"unsupported length of oracle price feeds (%d) for pair (%s)",
				len(priceFeeds), pair,
			)
		}
	}

	return allPairPrices, nil
}
