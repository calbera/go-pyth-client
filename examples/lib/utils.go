package lib

import (
	"math"

	"github.com/calbera/go-pyth-client/bindings/apyth"
	"github.com/calbera/go-pyth-client/types"
)

// QueryType is an enum that represents the type of query to be made to the Pythnet Hermes API.
type QueryType string

const (
	// LatestSync is a query type that fetches all latest prices together & synchronously.
	LatestSync QueryType = "latest-sync"
	// LatestAsync is a query type that fetches all latest prices individually, in parallel.
	LatestAsync QueryType = "latest-async"
	// Stream is a query type that subscribes to price updates and returns the latest cached price.
	StreamCached QueryType = "stream"
)

// QuerySettings is a struct that holds the settings for querying Hermes using the Pythnet client.
type QuerySettings struct {
	UseEma           bool
	DesiredPrecision int
	RequestType      QueryType
	SingleUpdateFee  uint
}

// Builds a PriceUpdate for a Pyth oracle result.
//
// NOTE: does not set the PythUpdateFee or the FeedType.
func GetPriceUpdateFromPythResult(
	lpd *types.LatestPriceData, pairIndex uint64, useEma bool, desiredPrecision int,
) *PriceUpdate {
	pu := GetPriceUpdateFromPythStructsPriceFeed(
		lpd.PriceFeed, pairIndex, useEma, desiredPrecision,
	)
	pu.PythUpdateData = [][]byte{lpd.UpdateData}
	return pu
}

// Builds a PriceUpdate for a Pyth price feed struct.
//
// NOTE: does not set the PythUpdateFee or the FeedType.
func GetPriceUpdateFromPythStructsPriceFeed(
	pf *apyth.PythStructsPriceFeed, pairIndex uint64, useEma bool, desiredPrecision int,
) *PriceUpdate {
	var psp apyth.PythStructsPrice
	if useEma {
		psp = pf.EmaPrice
	} else {
		psp = pf.Price
	}

	return &PriceUpdate{
		PairIndex: pairIndex,
		Price: Price(
			NormalizeToPrecision(psp.Price, int(-psp.Expo), desiredPrecision),
		),
		Conf: Price(
			NormalizeToPrecision(int64(psp.Conf), int(-psp.Expo), desiredPrecision),
		),
		TimeStamp: psp.PublishTime.Uint64(),
	}
}

// NormalizeToPrecision normalizes a int64 of decimals to a int64 with the desired precision.
func NormalizeToPrecision(i int64, decimals, precision int) int64 {
	if decimals < precision {
		return i * int64(math.Pow10(precision-decimals))
	} else if decimals > precision {
		return i / int64(math.Pow10(decimals-precision))
	}
	return i
}

// Returns baseP/quoteP in the desired precision. Assumes baseP/quoteP are in the same precision.
func CalculateTriangularPrice(baseP, quoteP int64, desiredPrecision int) int64 {
	return int64(float64(baseP) * math.Pow10(desiredPrecision) / float64(quoteP))
}

// Returns the confidence of baseP/quoteP in the desired precision. Assumes baseP/quoteP
// are in the same precision. Logic taken from Pyth Rust SDK.
// (https://github.com/pyth-network/pyth-sdk-rs/blob/main/pyth-sdk/src/price.rs#L424)
func CalculateTriangularConf(
	baseP, quoteP, baseC, quoteC int64, desiredPrecision int,
) int64 {
	triangularPrice := float64(baseP) / float64(quoteP)
	confMultiplier := math.Sqrt(
		(math.Pow(float64(baseC), 2) / math.Pow(float64(baseP), 2)) +
			(math.Pow(float64(quoteC), 2) / math.Pow(float64(quoteP), 2)),
	)
	return int64(triangularPrice * confMultiplier * math.Pow10(desiredPrecision))
}
