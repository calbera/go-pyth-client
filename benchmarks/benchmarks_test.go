package benchmarks_test

import (
	"log/slog"
	"time"

	"github.com/calbera/go-pyth-client/benchmarks"
)

// This file contains test utils that are shared for tests of Benchmarks.

var (
	testPairs = []string{"ATOM/USD", "USDC/USD", "ETH/USD", "TIA/USD", "BTC/USD"}

	testConfig = benchmarks.Config{
		APIEndpoint: "https://benchmarks.pyth.network",
		HTTPTimeout: 1 * time.Second,
		MaxRetries:  2,
	}

	testTime = time.Now().Add(-2 * time.Hour) // 2 hours ago
)

func setUp() *benchmarks.Client {
	pythClient, _ := benchmarks.NewClient(&testConfig, slog.Default())
	return pythClient
}
