package http

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/rs/zerolog/log"
	"gitlab.ozon.dev/route256/wallet/internal/pkg/token"
)

// WalletClient .
type WalletClient interface {
	//CreateAccount Создать счет
	CreateAccount(ctx context.Context, body *CreateAccountRequest) (*CreateAccountResponse, *http.Response, error)
	// GetAccountBalance Получить инфу по балансу
	GetAccountBalance(ctx context.Context, body *GetAccountBalanceRequest) (*GetAccountBalanceResponse, *http.Response, error)
	// Debit Поступление денег
	Debit(ctx context.Context, body *DebitRequest) (*DebitResponse, *http.Response, error)
	// Credit Списание денег
	Credit(ctx context.Context, body *CreditRequest) (*CreditResponse, *http.Response, error)
}

type walletClient struct {
	client   *http.Client
	basePath string
}

// NewWalletClient .
func NewWalletClient(basePath string, timeout time.Duration) WalletClient {
	return &walletClient{
		client:   &http.Client{Timeout: timeout},
		basePath: basePath,
	}
}

func doRequest[Req any, Resp any](ctx context.Context, client *http.Client, url string, method string, body *Req) (*Resp, *http.Response, error) {
	b := new(bytes.Buffer)
	err := json.NewEncoder(b).Encode(body)
	if err != nil {
		return nil, nil, err
	}

	if err != nil {
		return nil, nil, err
	}

	req, err := http.NewRequest(method, url, b)
	if err != nil {
		return nil, nil, err
	}
	session, ok := token.FromCtxToRequest(ctx)
	if ok {
		req.Header.Add(token.HeaderKey, session)
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, res, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Error().Msgf("Bad status code: %d", res.StatusCode)
		return nil, res, err
	}

	data, _ := io.ReadAll(res.Body)
	resp := new(Resp)
	err = json.Unmarshal(data, &resp)
	if err != nil {
		return nil, res, err
	}
	return resp, res, nil
}
