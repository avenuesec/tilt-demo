package repository

import (
	"context"

	"github.com/avenuesec/tilt-demo/services/quotations/internal/domain"
	"gorm.io/gorm"
)

type HistoryRepository struct {
	db *gorm.DB
}

func NewHistoryRepository(db *gorm.DB) (*HistoryRepository, error) {
	repository := &HistoryRepository{
		db: db,
	}

	repository.db.AutoMigrate(&domain.QuotationHistory{})

	return repository, nil
}

func (r *HistoryRepository) Insert(ctx context.Context, symbol string, price float64) error {
	entry := &domain.QuotationHistory{
		Price:  price,
		Symbol: symbol,
	}

	return r.db.WithContext(ctx).Create(entry).Error
}
