package http

import (
	"context"
	"net/http"
	"net/url"
)

// GetAccountBalanceRequest .
type GetAccountBalanceRequest struct {
	UserID    string `json:"userId"`
	AccountID string `json:"accountId"`
}

// GetAccountBalanceResponse .
type GetAccountBalanceResponse struct {
	AccountID string `json:"accountId"`
	Amount    int32  `json:"amount"`
}

// GetAccountBalance .
func (c *walletClient) GetAccountBalance(ctx context.Context, body *GetAccountBalanceRequest) (*GetAccountBalanceResponse, *http.Response, error) {
	uri, err := url.JoinPath(c.basePath, "/api/v1/users", body.UserID, "accounts", body.AccountID, "balance")
	if err != nil {
		return nil, nil, err
	}

	return doRequest[GetAccountBalanceRequest, GetAccountBalanceResponse](ctx, c.client, uri, http.MethodGet, nil)
}
