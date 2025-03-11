package cbr

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"gitlab.ozon.dev/route256/wallet/internal/pkg/cbr/valutes"
)

func NewClient(cbrURL string) *Client {
	return &Client{
		httpClient: &http.Client{},
		cbrURL:     cbrURL,
	}
}

type Client struct {
	httpClient *http.Client
	cbrURL     string
}

func (c *Client) GetExchangeRates(ctx context.Context, date time.Time) (*valutes.ValCurs, error) {
	if c.cbrURL == "" {
		return nil, errors.New("can not get rates: cbrURL not defined, check environment CBR_URL")
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.cbrURL+"/scripts/XML_daily.asp", nil)
	if err != nil {
		return nil, fmt.Errorf("can not create request error: %v", err)
	}

	if date.IsZero() {
		date = time.Now()
	}
	q := req.URL.Query()
	q.Add("date_req", date.Format("02/01/2006"))
	req.URL.RawQuery = q.Encode()

	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.0.0 Safari/537.36")

	fmt.Println(req.URL.String())

	response, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request to cbr error: %v", err)
	}

	return valutes.Unmarshal(response.Body)
}
