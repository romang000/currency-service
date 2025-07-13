package handler

import (
	"context"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/romapopov1212/currency-service/currency/internal/dto"
	"github.com/romapopov1212/currency-service/currency/internal/repository"
	"github.com/romapopov1212/currency-service/pkg/currency"
	"go.uber.org/zap"
	"time"
)

//go:generate mockgen -source=server.go -destination=mocks/mock_currency_service.go -package=mocks CurrencyService
type CurrencyService interface {
	GetCurrencyRatesInInterval(ctx context.Context, reqDTO *dto.CurrencyRequestDTO) ([]repository.CurrencyRate, error)
	FetchAndSaveCurrencyRates(ctx context.Context, baseCurrency string) error
	SetRateForCurrency(ctx context.Context, rate float64, baseCurrency string, targetCurrency string, date time.Time, ) (bool, error)
}

// todo tests
type CurrencyServer struct {
	currency.UnimplementedCurrencyServiceServer
	service CurrencyService
	logger  *zap.Logger

	requestCount    *prometheus.CounterVec
	requestDuration *prometheus.HistogramVec
	appUptime       prometheus.Gauge
}

func NewCurrencyServer(svc CurrencyService,
	logger *zap.Logger,
	requestCount *prometheus.CounterVec,
	requestDuration *prometheus.HistogramVec,
	appUptime prometheus.Gauge) CurrencyServer {

	return CurrencyServer{
		service:         svc,
		logger:          logger,
		requestCount:    requestCount,
		requestDuration: requestDuration,
		appUptime:       appUptime,
	}
}
