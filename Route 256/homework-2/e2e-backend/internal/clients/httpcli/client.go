package httpcli

import (
	"bytes"
	"context"
	"e2e-backend/internal/models"
	"e2e-backend/internal/pb/gitlab.ozon.dev/route256/wallet"
	"e2e-backend/internal/utils"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type WalletClient interface {
	CreateUser(ctx context.Context, req *models.User) (*models.User, *http.Response, error)
	LogIn(ctx context.Context, phone, password string) (*wallet.LogInResponse, *http.Response, error)
	GetUser(ctx context.Context, userID string) (*models.User, *http.Response, error)
	CreateAccount(ctx context.Context, userID, description string, amount int32) (*models.Account, *http.Response, error)
	GetAccountBalance(ctx context.Context, userID, accountID string) (*models.Balance, *http.Response, error)
	UpgradeUser(ctx context.Context, userID string) (*wallet.UpgradeUserResponse, *http.Response, error)
	Credit(ctx context.Context, userID, accountID string, amount int32) (*wallet.CreditResponse, *http.Response, error)
	Debit(ctx context.Context, userID, accountID string, amount int32) (*wallet.DebitResponse, *http.Response, error)
}

type walletClient struct {
	client   *http.Client
	basePath string
}

func NewWalletClient(basePath string, timeout time.Duration) WalletClient {
	return &walletClient{
		client:   &http.Client{Timeout: timeout},
		basePath: basePath,
	}
}

func (wc *walletClient) CreateUser(ctx context.Context, req *models.User) (*models.User, *http.Response, error) {
	url := fmt.Sprintf("%s/api/v1/users", wc.basePath)
	return do[models.User, models.User](ctx, wc.client, url, http.MethodPost, req)
}

func (wc *walletClient) LogIn(ctx context.Context, phone, password string) (*wallet.LogInResponse, *http.Response, error) {
	url := fmt.Sprintf("%s/api/v1/login", wc.basePath)
	reqBody := &wallet.LogInRequest{Phone: phone, Password: password}
	return do[wallet.LogInRequest, wallet.LogInResponse](ctx, wc.client, url, http.MethodPost, reqBody)
}

func (wc *walletClient) GetUser(ctx context.Context, userID string) (*models.User, *http.Response, error) {
	url := fmt.Sprintf("%s/users/%s", wc.basePath, userID)
	return do[any, models.User](ctx, wc.client, url, http.MethodGet, nil)
}

func (wc *walletClient) CreateAccount(ctx context.Context, userID, description string, amount int32) (*models.Account, *http.Response, error) {
	url := fmt.Sprintf("%s/api/v1/%s/accounts", wc.basePath, userID)
	reqBody := &wallet.CreateAccountRequest{
		UserId:      userID,
		Description: description,
		Amount:      amount,
	}
	return do[wallet.CreateAccountRequest, models.Account](ctx, wc.client, url, http.MethodPost, reqBody)
}

func (wc *walletClient) GetAccountBalance(ctx context.Context, userID, accountID string) (*models.Balance, *http.Response, error) {
	url := fmt.Sprintf("%s/api/v1/%s/accounts/%s/balance", wc.basePath, userID, accountID)
	return do[interface{}, models.Balance](ctx, wc.client, url, http.MethodGet, nil)
}

func (wc *walletClient) UpgradeUser(ctx context.Context, userID string) (*wallet.UpgradeUserResponse, *http.Response, error) {
	url := fmt.Sprintf("%s/users/%s/upgrade", wc.basePath, userID)
	return do[interface{}, wallet.UpgradeUserResponse](ctx, wc.client, url, http.MethodPost, nil)
}

func (wc *walletClient) Credit(ctx context.Context, userID, accountID string, amount int32) (*wallet.CreditResponse, *http.Response, error) {
	url := fmt.Sprintf("%s/api/v1/%s/accounts/%s/credit", wc.basePath, userID, accountID)
	reqBody := &wallet.CreditRequest{
		UserId:    userID,
		AccountId: accountID,
		Amount:    amount,
	}
	return do[wallet.CreditRequest, wallet.CreditResponse](ctx, wc.client, url, http.MethodPost, reqBody)
}

func (wc *walletClient) Debit(ctx context.Context, userID, accountID string, amount int32) (*wallet.DebitResponse, *http.Response, error) {
	url := fmt.Sprintf("%s/api/v1/%s/accounts/%s/debit", wc.basePath, userID, accountID)
	reqBody := &wallet.DebitRequest{
		UserId:    userID,
		AccountId: accountID,
		Amount:    amount,
	}
	return do[wallet.DebitRequest, wallet.DebitResponse](ctx, wc.client, url, http.MethodPost, reqBody)
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
