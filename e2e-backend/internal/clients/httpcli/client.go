package httpcli

import (
	"bytes"
	"context"
	"e2e-backend/internal/utils"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

type WalletClient interface {
	// TODO: необходимо определить интерфейс взаимодействия с Wallet http клиентов
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

// do - выполняет http запрос; Req и Resp - структуры запроса и ответа
func do[Req any, Resp any](ctx context.Context, client *http.Client, url string, method string, body *Req) (*Resp, *http.Response, error) {
	b := new(bytes.Buffer)
	err := json.NewEncoder(b).Encode(body)
	if err != nil {
		return nil, nil, err
	}

	req, err := http.NewRequest(method, url, b)
	if err != nil {
		return nil, nil, err
	}
	token, ok := utils.FromCtx(ctx)
	if ok {
		req.Header.Add(utils.HeaderKey, token)
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, res, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
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
