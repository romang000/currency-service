package service

import (
	"context"
	"fmt"
	"github.com/romapopov1212/currency-service/gateway/internal/dto"
	"github.com/romapopov1212/currency-service/pkg/currency"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type CurrencyService struct {
	currencyClient currency.CurrencyServiceClient
}

func NewCurrency(currencyClient currency.CurrencyServiceClient) CurrencyService {
	return CurrencyService{
		currencyClient: currencyClient,
	}
}

func (svc *CurrencyService) GetCurrencyRates(
	ctx context.Context,
	request dto.ParsedCurrencyRequest,
) (*dto.CurrencyResponse, error) {
	pbResp, err := svc.currencyClient.GetRate(
		ctx, &currency.GetRateRequest{
			Currency: request.Currency,
			DateFrom: timestamppb.New(request.DateFrom),
			DateTo:   timestamppb.New(request.DateTo),
		},
	)

	if err != nil {
		return nil, fmt.Errorf("currencyClient.GetRate: %s", err)
	}

	resp := &dto.CurrencyResponse{
		Currency: pbResp.GetCurrency(),
		Rates:    make([]dto.CurrencyRate, 0, len(pbResp.Rates)),
	}

	for _, rate := range pbResp.Rates {
		resp.Rates = append(
			resp.Rates, dto.CurrencyRate{
				Rate: rate.Rate,
				Date: rate.Date.AsTime(),
			},
		)
	}
	return resp, nil
}

func (svc *CurrencyService) SetCurrencyRate(
	ctx context.Context,
	request dto.SetCurrencyRequest,
) (*dto.SetCurrencyResponse, error) {
	pbResp, err := svc.currencyClient.SetRate(
		ctx, &currency.SetRateRequest{
			Rate:           float32(request.Rate),
			BaseCurrency:   request.BaseCurrency,
			TargetCurrency: request.TargetCurrency,
			Date:           timestamppb.New(request.Date),
		},
	)

	if err != nil {
		return nil, fmt.Errorf("currencyClient.SetRate: {%w}", err)
	}

	resp := &dto.SetCurrencyResponse{
		Result: pbResp.Success,
	}

	return resp, nil
}
