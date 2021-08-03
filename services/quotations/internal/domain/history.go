package domain

import "gorm.io/gorm"

type QuotationHistory struct {
	gorm.Model

	Symbol string
	Price  float64
}
