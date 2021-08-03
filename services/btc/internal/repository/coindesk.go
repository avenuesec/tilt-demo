package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type CoindeckResult struct {
	BPI struct {
		USD struct {
			RateFloat float64 `json:"rate_float"`
		} `json:"USD"`
	} `json:"bpi"`
}

type CointdeckRepository struct {
	url    string
	client *http.Client
}

func NewCointdeckRepository(url string, client *http.Client) *CointdeckRepository {
	return &CointdeckRepository{
		url:    url,
		client: client,
	}
}

func (c *CointdeckRepository) Current(ctx context.Context) (float64, error) {
	url := fmt.Sprintf(c.url, "BTC")

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)

	if err != nil {
		return 0, err
	}

	res, err := c.client.Do(req)

	if err != nil {
		return 0, err
	}

	defer res.Body.Close()

	result := new(CoindeckResult)

	if err := json.NewDecoder(res.Body).Decode(result); err != nil {
		return 0, err
	}

	return result.BPI.USD.RateFloat, nil
}
