package lib_test

import (
	"testing"

	"github.com/calbera/go-pyth-client/examples/lib"
)

func TestNormalizeToPrecision(t *testing.T) {
	tests := []struct {
		i         int64
		decimals  int
		precision int
		expected  lib.Price
	}{
		{100000, 6, 3, 100},
		{100, 3, 6, 100000},
		{1000, 3, 3, 1000},
	}

	for _, test := range tests {
		result := lib.NormalizeToPrecision(test.i, test.decimals, test.precision)
		if result != test.expected {
			t.Errorf("Expected %v, got %v", test.expected, result)
		}
	}
}

func TestCalculateTriangularPrice(t *testing.T) {
	tests := []struct {
		baseP, quoteP    lib.Price
		desiredPrecision int
		expected         lib.Price
	}{
		{3e10, 4e10, 10, 3e10 / 4},
		{1000, 10000, 10, 1e10 / 10},
		{1000, 100, 2, 1000},
		{500, 100, 1, 50},
		{123456, 7890, 5, 1564714},
		{0, 100, 2, 0},
	}

	for _, test := range tests {
		result := lib.CalculateTriangularPrice(test.baseP, test.quoteP, test.desiredPrecision)
		if result != test.expected {
			t.Errorf("CalculateTriangularPrice(%d, %d, %d) = %d; want %d",
				test.baseP, test.quoteP, test.desiredPrecision, result, test.expected)
		}
	}
}

func TestCalculateTriangularConf(t *testing.T) {
	tests := []struct {
		baseP, quoteP, baseC, quoteC lib.Price
		desiredPrecision             int
		expected                     lib.Price
	}{
		{3e10, 4e10, 1e10, 1e10, 10, 5e10 / 16},
		{1000, 10000, 1, 5, 10, 1118033},
		{1000, 100, 10, 10, 2, 100},
		{500, 100, 5, 5, 1, 2},
		{123456, 7890, 100, 100, 5, 19872},
		{100, 1, 1, 1, 2, 10000},
	}

	for _, test := range tests {
		result := lib.CalculateTriangularConf(
			test.baseP, test.quoteP, test.baseC, test.quoteC, test.desiredPrecision,
		)
		if result != test.expected {
			t.Errorf("CalculateTriangularConf(%d, %d, %d, %d, %d) = %d; want %d",
				test.baseP, test.quoteP, test.baseC, test.quoteC, test.desiredPrecision,
				result, test.expected)
		}
	}
}
