package handler

import (
	"context"
	"fmt"
	"time"

	"github.com/romapopov1212/currency-service/currency/internal/dto"
	"github.com/romapopov1212/currency-service/pkg/currency"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s CurrencyServer) GetRate(ctx context.Context, request *currency.GetRateRequest) (*currency.GetRateResponse, error) {
	start := time.Now()
	reqDTO := dto.CurrencyRequestDTOFromProtobuf(request, dto.DefaultBaseCurrency)

	// todo метрики в мидлвары
	s.requestCount.WithLabelValues("GetRate").Inc()
	rates, err := s.service.GetCurrencyRatesInInterval(ctx, reqDTO)
	if err != nil {
		return nil, fmt.Errorf("service.GetCurrencyRatesInInterval: %w", err)
	}

	rateRecords := make([]*currency.RateRecord, len(rates))
	for i, rate := range rates {
		rateRecords[i] = &currency.RateRecord{
			Date: timestamppb.New(rate.Date),
			Rate: rate.Rate,
		}
	}

	s.requestDuration.WithLabelValues("GetExchangeRate").Observe(time.Since(start).Seconds())
	return &currency.GetRateResponse{
		Currency: reqDTO.TargetCurrency,
		Rates:    rateRecords,
	}, nil
}

func (s CurrencyServer) SetRate(ctx context.Context, request *currency.SetRateRequest) (*currency.SetRateResponse, error) {
	// todo метрики
	res, err := s.service.SetRateForCurrency(ctx, float64(request.Rate), request.BaseCurrency, request.TargetCurrency, request.Date.AsTime())

	if err != nil {
		return nil, fmt.Errorf("service.SetRateForCurrency: {%w}", err)
	}

	return &currency.SetRateResponse{
		Success: res,
	}, nil
}
