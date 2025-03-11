package http

import (
	"context"
	"net/http"
	"net/url"
)

// DebitRequest .
type DebitRequest struct {
	UserID      string `json:"userId"`
	Amount      int32  `json:"amount"`
	AccountID   string `json:"accountId"`
	OperationID string `json:"operationId"`
}

// DebitResponse .
type DebitResponse struct {
	Status string `json:"status"`
}

// Debit .
func (c *walletClient) Debit(ctx context.Context, body *DebitRequest) (*DebitResponse, *http.Response, error) {
	uri, err := url.JoinPath(c.basePath, "/api/v1/users", body.UserID, "accounts", body.AccountID, "debit")
	if err != nil {
		return nil, nil, err
	}

	return doRequest[DebitRequest, DebitResponse](ctx, c.client, uri, http.MethodPost, body)
}
