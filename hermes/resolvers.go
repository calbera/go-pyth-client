package hermes

import (
	"errors"
	"math/big"

	"github.com/calbera/go-pyth-client/bindings/apyth"
	"github.com/calbera/go-pyth-client/feeds"
	"github.com/calbera/go-pyth-client/types"
	"github.com/ethereum/go-ethereum/common"
)

// Resolve the latestPriceResponse into a LatestPriceData according to the given configuration.
// Throws an error if the pyth response is not correct according to API specification.
func (c *Client) resolveOne(lpr *latestPriceResponse) (*types.LatestPriceData, error) {
	if len(lpr.Parsed) != 1 || len(lpr.Binary.Data) != 1 {
		return nil, errors.New("pyth response does not have exactly 1 response")
	}

	// Set the fields from API response that can be immediately resolved without error.
	lpd := &types.LatestPriceData{
		PriceFeed: &apyth.PythStructsPriceFeed{
			Price: apyth.PythStructsPrice{
				Expo:        lpr.Parsed[0].Price.Expo,
				PublishTime: big.NewInt(lpr.Parsed[0].Price.PublishTime),
			},
			EmaPrice: apyth.PythStructsPrice{
				Expo:        lpr.Parsed[0].EmaPrice.Expo,
				PublishTime: big.NewInt(lpr.Parsed[0].EmaPrice.PublishTime),
			},
		},
	}
	copy(lpd.PriceFeed.Id[:], common.FromHex(lpr.Parsed[0].ID))

	// Decode API response types into golang types.
	priceVal, pOk := new(big.Int).SetString(lpr.Parsed[0].Price.Price, 10)
	if !pOk {
		return nil, errors.New("failed to convert price string to big.Int")
	}
	lpd.PriceFeed.Price.Price = priceVal.Int64()

	confVal, cOk := new(big.Int).SetString(lpr.Parsed[0].Price.Conf, 10)
	if !cOk {
		return nil, errors.New("failed to convert conf string to big.Int")
	}
	lpd.PriceFeed.Price.Conf = confVal.Uint64()

	emaPriceVal, epOk := new(big.Int).SetString(lpr.Parsed[0].EmaPrice.Price, 10)
	if !epOk {
		return nil, errors.New("failed to convert ema price string to big.Int")
	}
	lpd.PriceFeed.EmaPrice.Price = emaPriceVal.Int64()

	emaConfVal, ecOk := new(big.Int).SetString(lpr.Parsed[0].EmaPrice.Conf, 10)
	if !ecOk {
		return nil, errors.New("failed to convert ema conf string to big.Int")
	}
	lpd.PriceFeed.EmaPrice.Conf = emaConfVal.Uint64()

	// Set the update data based on whether the mock is being used or not.
	if c.cfg.UseMock {
		priceUpdateData, err := c.pythABI.Methods[queryPriceFeed].Outputs.Pack(lpd.PriceFeed)
		if err != nil {
			return nil, err
		}
		lpd.UpdateData = priceUpdateData
	} else {
		lpd.UpdateData = common.FromHex(lpr.Binary.Data[0])
	}

	return lpd, nil
}

// Resolves an individual price response into a mapping of feeds to LatestPriceDatas according to
// the given configuration. Throws an error if the pyth response is not correct according to API
// specification.
func (c *Client) resolveMany(lpr *latestPriceResponse, lpds map[string]*types.LatestPriceData) error {
	if len(lpr.Parsed) < 1 || len(lpr.Binary.Data) != 1 {
		// NOTE: Pyth will always return 1 hex string even for multiple feeds.
		return errors.New("pyth response does not have the correct number of response")
	}

	// Handle a nil map for storing results.
	if lpds == nil {
		return errors.New("nil price results map")
	}

	for _, pr := range lpr.Parsed {
		// Set the fields from API response that can be immediately resolved without error.
		lpd := &types.LatestPriceData{
			PriceFeed: &apyth.PythStructsPriceFeed{
				Price: apyth.PythStructsPrice{
					Expo:        pr.Price.Expo,
					PublishTime: big.NewInt(pr.Price.PublishTime),
				},
				EmaPrice: apyth.PythStructsPrice{
					Expo:        pr.EmaPrice.Expo,
					PublishTime: big.NewInt(pr.EmaPrice.PublishTime),
				},
			},
		}
		copy(lpd.PriceFeed.Id[:], common.FromHex(pr.ID))

		// Decode API response types into golang types.
		priceVal, pOk := new(big.Int).SetString(pr.Price.Price, 10)
		if !pOk {
			return errors.New("failed to convert price string to big.Int")
		}
		lpd.PriceFeed.Price.Price = priceVal.Int64()

		confVal, cOk := new(big.Int).SetString(pr.Price.Conf, 10)
		if !cOk {
			return errors.New("failed to convert conf string to big.Int")
		}
		lpd.PriceFeed.Price.Conf = confVal.Uint64()

		emaPriceVal, epOk := new(big.Int).SetString(pr.EmaPrice.Price, 10)
		if !epOk {
			return errors.New("failed to convert ema price string to big.Int")
		}
		lpd.PriceFeed.EmaPrice.Price = emaPriceVal.Int64()

		emaConfVal, ecOk := new(big.Int).SetString(pr.EmaPrice.Conf, 10)
		if !ecOk {
			return errors.New("failed to convert ema conf string to big.Int")
		}
		lpd.PriceFeed.EmaPrice.Conf = emaConfVal.Uint64()

		// Set the update data based on whether the mock is being used or not.
		if c.cfg.UseMock {
			priceUpdateData, err := c.pythABI.Methods[queryPriceFeed].Outputs.Pack(lpd.PriceFeed)
			if err != nil {
				return err
			}
			lpd.UpdateData = priceUpdateData
		} else {
			lpd.UpdateData = common.FromHex(lpr.Binary.Data[0]) // same for all feeds
		}

		// Assign the price data to its corresponding feed ID.
		lpds[feeds.MustGetPriceFeed(c.cfg.feedVersion, pr.ID)] = lpd
	}

	return nil
}

// Resolves a price response from sse streaming into the cached ssePriceData.
// Throws an error if the sse response is not correct according to API specification.
// NOTE: assumes that the ssePriceCached write lock is already held!
func (c *Client) resolveSsePrice(lpr *latestPriceResponse) error {
	if len(lpr.Parsed) < 1 || len(lpr.Binary.Data) != 1 {
		// NOTE: Pyth will always return 1 hex string even for multiple feeds.
		return errors.New("pyth response does not have the correct number of response")
	}

	// Refresh the ssePrice upon the arrived price response
	if err := c.resolveMany(lpr, c.ssePriceCached.latestPrice); err != nil {
		c.logger.Error("skipping msg encountered an error when unmarshalling streaming data")
		return err
	}
	return nil
}
