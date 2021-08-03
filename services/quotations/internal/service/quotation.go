package service

import (
	"context"

	"github.com/avenuesec/tilt-demo/pkg/model"
	"github.com/avenuesec/tilt-demo/services/quotations/internal/repository"
)

type QuotationService struct {
	factory *QuotationFactory
	history *repository.HistoryRepository
}

func NewQuotationService(factory *QuotationFactory, history *repository.HistoryRepository) *QuotationService {
	return &QuotationService{
		factory: factory,
		history: history,
	}
}

func (s *QuotationService) CurrencyPrice(ctx context.Context, currency Currency) (*model.CurrencyQuotation, error) {
	resolver, err := s.factory.ResolveProvider(currency)

	if err != nil {
		return nil, err
	}

	quotation, err := resolver.CurrentPrice(ctx)

	if err != nil {
		return nil, err
	}

	if err := s.history.Insert(ctx, string(currency), quotation.Buy); err != nil {
		return nil, err
	}

	return &model.CurrencyQuotation{
		Buy:  quotation.Buy,
		Sell: quotation.Sell,
	}, nil
}
