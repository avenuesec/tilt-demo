package service

import (
	"context"
	"errors"
	"net/http"

	"github.com/avenuesec/tilt-demo/pkg/config"
	"github.com/avenuesec/tilt-demo/services/quotations/internal/domain"
	"github.com/avenuesec/tilt-demo/services/quotations/internal/repository"
)

const (
	Dolar Currency = "USD"
)

var (
	ErrCurrencyHasNoProvider = errors.New("currency has no provider registed")
)

type Currency string

type QuotationProvider interface {
	CurrentPrice(context.Context) (*domain.CurrencyQuotation, error)
}

type QuotationFactory struct {
	cfg *config.Configuration
}

func NewQuotationFactory(cfg *config.Configuration) *QuotationFactory {
	return &QuotationFactory{
		cfg: cfg,
	}
}

func (f *QuotationFactory) ResolveProvider(currency Currency) (QuotationProvider, error) {
	switch currency {
	case Dolar:
		return repository.NewPTAXRepository(http.DefaultClient, f.cfg.PtaxURL), nil
	}

	return nil, ErrCurrencyHasNoProvider
}
