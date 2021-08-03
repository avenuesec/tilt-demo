package service

import (
	"context"

	"github.com/avenuesec/tilt-demo/pkg/model"
	"github.com/avenuesec/tilt-demo/services/btc/internal/repository"
)

type BTCPriceService struct {
	price     *repository.CointdeckRepository
	quotation *repository.QuotationRepository
}

func NewBTCPriceService(quotation *repository.QuotationRepository, price *repository.CointdeckRepository) *BTCPriceService {
	return &BTCPriceService{
		price:     price,
		quotation: quotation,
	}
}

func (s *BTCPriceService) Price(ctx context.Context) (*model.Price, error) {
	price, err := s.price.Current(ctx)

	if err != nil {
		return nil, err
	}

	quotation, err := s.quotation.Quotation(ctx)

	if err != nil {
		return nil, err
	}

	return &model.Price{
		USD: price,
		BRL: price * quotation,
	}, nil
}
