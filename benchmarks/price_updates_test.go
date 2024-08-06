package benchmarks_test

import (
	"context"
	"math"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/calbera/go-pyth-client/feeds"

	"github.com/ethereum/go-ethereum/common"
)

func TestGetPriceUpdatesSync(t *testing.T) {
	pythClient := setUp()

	prices, err := pythClient.GetHistoricalPriceUpdatesSync(
		context.Background(), testTime, testPairs...,
	)
	assert.Nil(t, err)
	assert.Equal(t, len(testPairs), len(prices))

	for _, pair := range testPairs {
		assert.Contains(t, prices, pair)

		// Correct Price Feed ID returned
		assert.Equal(
			t,
			prices[pair].Id[:],
			common.Hex2Bytes(feeds.MustGetPriceFeedID(feeds.EVMStable, pair)),
		)

		// Valid, non-zero prices returned
		assert.Greater(t, prices[pair].Price.Price, int64(0))
		assert.Greater(t, prices[pair].EmaPrice.Price, int64(0))

		// Within 5 seconds of the desired time
		assert.LessOrEqual(
			t,
			math.Abs(float64(prices[pair].Price.PublishTime.Int64())-float64(testTime.Unix())),
			float64(5*time.Second),
		)
	}
}