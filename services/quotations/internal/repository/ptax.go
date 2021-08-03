package repository

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/avenuesec/tilt-demo/services/quotations/internal/domain"
)

var (
	ErrPTAXUnavaliable = errors.New("PTAX is unavaliable")
)

type PTAX struct {
	Buy  float64 `json:"cotacaoCompra"`
	Sell float64 `json:"cotacaoVenda"`
}

type ODataResult struct {
	Value []*PTAX `json:"value"`
}

type PTAXRepository struct {
	client *http.Client
	url    string
}

func NewPTAXRepository(client *http.Client, url string) *PTAXRepository {
	return &PTAXRepository{
		url:    url,
		client: client,
	}
}

func (p *PTAXRepository) CurrentPrice(ctx context.Context) (*domain.CurrencyQuotation, error) {
	date := time.Now().Format("02-01-2006")

	reqURL := fmt.Sprintf("%s/olinda/servico/PTAX/versao/v1/odata/CotacaoDolarDia(dataCotacao=@date)?$format=json&@date='%s'", p.url, date)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL, nil)

	if err != nil {
		return nil, err
	}

	res, err := p.client.Do(req)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	result := new(ODataResult)

	if err := json.NewDecoder(res.Body).Decode(result); err != nil {
		return nil, err
	}

	if len(result.Value) <= 0 {
		return nil, ErrPTAXUnavaliable
	}

	return &domain.CurrencyQuotation{
		Buy:  result.Value[0].Buy,
		Sell: result.Value[0].Sell,
	}, nil
}
