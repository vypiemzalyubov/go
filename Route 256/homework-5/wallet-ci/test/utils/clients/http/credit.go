package http

import (
	"context"
	"net/http"
	"net/url"
)

// CreditRequest .
type CreditRequest struct {
	UserID      string `json:"userId"`
	Amount      int32  `json:"amount"`
	AccountID   string `json:"accountId"`
	OperationID string `json:"operationId"`
}

// CreditResponse .
type CreditResponse struct {
	Status string `json:"status"`
}

// Credit .
func (c *walletClient) Credit(ctx context.Context, body *CreditRequest) (*CreditResponse, *http.Response, error) {
	uri, err := url.JoinPath(c.basePath, "/api/v1/users", body.UserID, "accounts", body.AccountID, "credit")
	if err != nil {
		return nil, nil, err
	}

	return doRequest[CreditRequest, CreditResponse](ctx, c.client, uri, http.MethodPost, body)
}
