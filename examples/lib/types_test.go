package lib_test

import (
	"math/rand"
	"testing"

	"github.com/calbera/go-pyth-client/examples/lib"
)

// Results from running these benchmarks tend to follow this pattern:
//
// goos: darwin
// goarch: arm64
// pkg: github.com/berachain/bts/pkg/oracle
// BenchmarkPriceCmp-12     	1000000000	      1.116 ns/op	    0 B/op	    0 allocs/op
// BenchmarkPriceCmpUnsafe-12    	1000000000	      1.131 ns/op	    0 B/op	    0 allocs/op
//
// Takeaway: PriceCmp and PriceCmpUnsafe are practically similar in performance.

// BenchmarkPriceCmp benchmarks the PriceCmp function.
func BenchmarkPriceCmp(b *testing.B) {
	// Generate random test data
	prices := generateTestPrices(b.N)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		lib.PriceCmp(prices[i%len(prices)], prices[(i+1)%len(prices)])
	}
}

// BenchmarkPriceCmpUnsafe benchmarks the PriceCmp2 function.
func BenchmarkPriceCmpUnsafe(b *testing.B) {
	// Generate random test data
	prices := generateTestPrices(b.N)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		lib.PriceCmpUnsafe(prices[i%len(prices)], prices[(i+1)%len(prices)])
	}
}

// generateTestPrices generates a slice of random Price values for benchmarking.
func generateTestPrices(n int) []lib.Price {
	prices := make([]lib.Price, n)
	for i := 0; i < n; i++ {
		prices[i] = lib.Price(rand.Int63n(int64(lib.MaxPrice)))
	}
	return prices
}

func TestPriceCmp(t *testing.T) {
	tests := []struct {
		name     string
		a        lib.Price
		b        lib.Price
		expected int
	}{
		{
			name:     "a > b",
			a:        lib.Price(200),
			b:        lib.Price(100),
			expected: 1,
		},
		{
			name:     "a == b",
			a:        lib.Price(100),
			b:        lib.Price(100),
			expected: 0,
		},
		{
			name:     "a < b",
			a:        lib.Price(100),
			b:        lib.Price(200),
			expected: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lib.PriceCmp(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestPriceCmpUnsafe(t *testing.T) {
	tests := []struct {
		name     string
		a        lib.Price
		b        lib.Price
		expected int
	}{
		{
			name:     "a > b",
			a:        lib.Price(200),
			b:        lib.Price(100),
			expected: 100,
		},
		{
			name:     "a == b",
			a:        lib.Price(100),
			b:        lib.Price(100),
			expected: 0,
		},
		{
			name:     "a < b",
			a:        lib.Price(100),
			b:        lib.Price(200),
			expected: -100,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lib.PriceCmpUnsafe(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}
