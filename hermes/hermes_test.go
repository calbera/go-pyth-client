package hermes_test

import (
	"context"
	"time"

	"log/slog"

	"github.com/calbera/go-pyth-client/hermes"
)

// This file contains test utils that are shared for tests of Hermes.

var testPairs = []string{"ATOM/USD", "USDC/USD", "ETH/USD", "TIA/USD", "BTC/USD"}

var testConfig = hermes.Config{
	// Offchain parameters
	APIEndpoint: "https://hermes.pyth.network",
	HTTPTimeout: 1 * time.Second,
	MaxRetries:  2,
	Ecosystem:   "EVM-Stable",

	// Onchain parameters
	UseMock: true, // Uses the mock Pyth contract rather than the real one.
}

func setUp() (context.Context, *hermes.Client) {
	// set up Pyth client and subscribe
	pythClient, _ := hermes.NewClient(&testConfig, slog.Default())

	return context.Background(), pythClient
}
