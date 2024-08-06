package benchmarks

import (
	"errors"
	"math/big"

	"github.com/calbera/go-pyth-client/bindings/apyth"
	"github.com/calbera/go-pyth-client/feeds"
	"github.com/ethereum/go-ethereum/common"
)

// resolveMany resolves the price data for multiple price feeds.
func resolveMany(
	priceResp *priceResponse, priceResults map[string]*apyth.PythStructsPriceFeed,
) error {
	// Handle a nil map for storing results.
	if priceResults == nil {
		return errors.New("nil price results map")
	}

	// Iterate over the parsed responses and resolve them.
	for _, pr := range priceResp.Parsed {
		// Set the fields from API response that can be immediately resolved without error.
		pf := &apyth.PythStructsPriceFeed{
			Price: apyth.PythStructsPrice{
				Expo:        pr.Price.Expo,
				PublishTime: big.NewInt(pr.Price.PublishTime),
			},
			EmaPrice: apyth.PythStructsPrice{
				Expo:        pr.EmaPrice.Expo,
				PublishTime: big.NewInt(pr.EmaPrice.PublishTime),
			},
		}
		copy(pf.Id[:], common.FromHex(pr.ID))

		// Decode API response types into golang types.
		priceVal, pOk := new(big.Int).SetString(pr.Price.Price, 10)
		if !pOk {
			return errors.New("failed to convert price string to big.Int")
		}
		pf.Price.Price = priceVal.Int64()

		confVal, cOk := new(big.Int).SetString(pr.Price.Conf, 10)
		if !cOk {
			return errors.New("failed to convert conf string to big.Int")
		}
		pf.Price.Conf = confVal.Uint64()

		emaPriceVal, epOk := new(big.Int).SetString(pr.EmaPrice.Price, 10)
		if !epOk {
			return errors.New("failed to convert ema price string to big.Int")
		}
		pf.EmaPrice.Price = emaPriceVal.Int64()

		emaConfVal, ecOk := new(big.Int).SetString(pr.EmaPrice.Conf, 10)
		if !ecOk {
			return errors.New("failed to convert ema conf string to big.Int")
		}
		pf.EmaPrice.Conf = emaConfVal.Uint64()

		priceResults[feeds.MustGetPriceFeed(feeds.EVMStable, pr.ID)] = pf
	}

	return nil
}
