package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/avenuesec/tilt-demo/pkg/model"
)

type QuotationRepository struct {
	url    string
	client *http.Client
}

func NewQuotationRepository(url string, client *http.Client) *QuotationRepository {
	return &QuotationRepository{
		url:    url,
		client: client,
	}
}

func (r *QuotationRepository) Quotation(ctx context.Context) (float64, error) {
	url := fmt.Sprintf("http://%s/quotations/USD", r.url)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)

	if err != nil {
		return 0, err
	}

	res, err := r.client.Do(req)

	if err != nil {
		return 0, err
	}

	defer res.Body.Close()

	quotation := new(model.CurrencyQuotation)

	if err := json.NewDecoder(res.Body).Decode(quotation); err != nil {
		return 0, err
	}

	return quotation.Buy, nil
}
