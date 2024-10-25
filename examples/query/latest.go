package query

import (
	"context"
	"fmt"
	"math/big"

	client "github.com/calbera/go-pyth-client"
	"github.com/calbera/go-pyth-client/examples/lib"
	"github.com/calbera/go-pyth-client/types"
)

// Requires a Pyth client to fetch prices from Pyth Hermes API. Assumes that all required feeds
// to be queried for `oracleFeeds` are in the `uniqueFeeds` slice.
func GetAllLatestPrices(
	ctx context.Context, pythClient client.Hermes, qs *lib.QuerySettings,
	pairIndexes map[string]uint64, oracleFeeds map[string][]string, uniqueFeeds []string,
) (lib.PriceUpdates, error) {
	// Query Pyth for the unique price feeds necessary
	var (
		priceData map[string]*types.LatestPriceData
		err       error
	)
	switch qs.RequestType {
	case lib.StreamCached:
		priceData, err = pythClient.GetCachedLatestPriceUpdates(ctx, uniqueFeeds...)
	case lib.LatestSync:
		priceData, err = pythClient.GetLatestPriceUpdatesSync(ctx, uniqueFeeds...)
	case lib.LatestAsync:
		fallthrough
	default:
		priceData, err = pythClient.GetLatestPriceUpdatesAsync(ctx, uniqueFeeds...)
	}
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
			pu := lib.GetPriceUpdateFromPythResult(
				priceData[priceFeeds[0]], pairIndex, qs.UseEma, qs.DesiredPrecision,
			)

			// Set the update fee depending on the request type to the API.
			if qs.RequestType == lib.LatestAsync {
				// If latest async, each feed's update data is stored separately.
				pu.PythUpdateFee = big.NewInt(int64(qs.SingleUpdateFee))
			} else {
				// For other request types, all feeds' update data is stored in each feed's data.
				pu.PythUpdateFee = big.NewInt(int64(qs.SingleUpdateFee) * int64(len(uniqueFeeds)))
			}

			allPairPrices[pairIndex] = pu

		case lib.Feed_TRIANGULAR:
			baseP := lib.GetPriceUpdateFromPythResult(
				priceData[priceFeeds[0]], pairIndex, qs.UseEma, qs.DesiredPrecision,
			)
			quoteP := lib.GetPriceUpdateFromPythResult(
				priceData[priceFeeds[1]], pairIndex, qs.UseEma, qs.DesiredPrecision,
			)

			// Use the older of the two as the timestamp.
			var timestamp uint64
			if baseP.TimeStamp <= quoteP.TimeStamp {
				timestamp = baseP.TimeStamp
			} else {
				timestamp = quoteP.TimeStamp
			}

			pu := &lib.PriceUpdate{
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

			// Set the update data depending on the request type to the API.
			if qs.RequestType == lib.LatestAsync {
				// If latest async, each feed's update data is stored separately.
				pu.PythUpdateData = [][]byte{
					baseP.PythUpdateData[0], quoteP.PythUpdateData[0],
				}
				pu.PythUpdateFee = big.NewInt(int64(qs.SingleUpdateFee) * 2)
			} else {
				// For other request types, all feeds' update data is stored in each feed's data.
				pu.PythUpdateData = baseP.PythUpdateData
				pu.PythUpdateFee = big.NewInt(int64(qs.SingleUpdateFee) * int64(len(uniqueFeeds)))
			}

			allPairPrices[pairIndex] = pu

		default:
			return nil, fmt.Errorf(
				"unsupported length of oracle price feeds (%d) for pair (%s)",
				len(priceFeeds), pair,
			)
		}
	}

	return allPairPrices, nil
}
