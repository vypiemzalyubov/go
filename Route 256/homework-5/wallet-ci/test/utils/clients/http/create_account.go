package http

import (
	"context"
	"net/http"
	"net/url"
)

// CreateAccountRequest .
type CreateAccountRequest struct {
	UserID      string `json:"userId"`
	Amount      int32  `json:"amount"`
	Description string `json:"description"`
}

// CreateAccountResponse .
type CreateAccountResponse struct {
	Amount      int32  `json:"amount"`
	AccountID   string `json:"accountId"`
	Description string `json:"description"`
}

// CreateAccount .
func (c *walletClient) CreateAccount(ctx context.Context, body *CreateAccountRequest) (*CreateAccountResponse, *http.Response, error) {
	uri, err := url.JoinPath(c.basePath, "/api/v1/users", body.UserID, "accounts")
	if err != nil {
		return nil, nil, err
	}

	return doRequest[CreateAccountRequest, CreateAccountResponse](ctx, c.client, uri, http.MethodPost, body)
}
