package hermes_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubscribePriceStreaming(t *testing.T) {
	ctx, pythClient := setUp()

	pythClient.SubscribePriceStreaming(ctx, testPairs...)

	prices, err := pythClient.GetCachedLatestPriceUpdates(ctx, testPairs...)
	assert.NoError(t, err)
	assert.Equal(t, 5, len(prices))
	for _, pair := range testPairs {
		assert.Contains(t, prices, pair)
	}
}

func TestSubscribePriceStreaming_EmptyRequests(t *testing.T) {
	ctx, pythClient := setUp()

	pythClient.SubscribePriceStreaming(ctx, testPairs...)

	var empty_pair = []string{}

	prices, err := pythClient.GetCachedLatestPriceUpdates(ctx, empty_pair...)
	assert.Error(t, err)
	assert.Nil(t, prices)
}

func TestSubscribePriceStreaming_PriceFeedNotSubscribed(t *testing.T) {
	ctx, pythClient := setUp()

	pythClient.SubscribePriceStreaming(ctx, testPairs...)

	var feed = []string{"USD/CAD"}

	prices, err := pythClient.GetCachedLatestPriceUpdates(ctx, feed...)
	assert.Error(t, err)
	assert.Nil(t, prices)
}

// To run this benchmark only without other tests: `go test -run=^$ -bench=BenchmarkGetCachedLatestPriceUpdates`
func BenchmarkGetCachedLatestPriceUpdates(b *testing.B) {
	ctx, pythClient := setUp()

	pythClient.SubscribePriceStreaming(ctx, testPairs...)

	for i := 0; i < b.N; i++ {
		prices, err := pythClient.GetCachedLatestPriceUpdates(ctx, testPairs...)
		assert.NoError(b, err)
		assert.Equal(b, 5, len(prices))
	}
}
