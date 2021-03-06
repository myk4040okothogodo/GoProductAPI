package server


import (
   "context" 
   "github.com/hashicorp/go-hclog"
   "github.com/myk4040okothogodo/Grpc2/currency/data"
    protos "github.com/myk4040okothogodo/Grpc2/currency/protos/currency/protos"
)

// Currency is a gRPC server it implements the methods defined by the CurrencyServer interface
type Currency struct {
	log hclog.Logger
  rates *data.ExchangeRates
}

// NewCurrency creates a new Currency server
func NewCurrency(r *data.ExchangeRates, l hclog.Logger) *Currency {
	return &Currency{ l, r}
}

// GetRate implements the CurrencyServer GetRate method and returns the currency exchange rate
// for the two given currencies.
func (c *Currency) GetRate(ctx context.Context, rr *protos.RateRequest) (*protos.RateResponse, error) {
	c.log.Info("Handle request for GetRate", "base", rr.GetBase(), "destination", rr.GetDestination())

  rate, err := c.rates.GetRate(rr.GetBase().String(), rr.GetDestination().String())

  if err != nil {
      return nil,err
  }

  return &protos.RateResponse{Rate: rate}, nil

}
