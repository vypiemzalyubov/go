package wallet

import (
	context "context"

	desc "gitlab.ozon.dev/route256/wallet/pkg/api/wallet"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// GetAccountBalance получить инфу по балансу счета
func (i *Implementation) GetAccountBalance(ctx context.Context, req *desc.GetAccountBalanceRequest) (*desc.AccountBalanceResponse, error) {
	if err := i.CheckToken(ctx, req.UserId); err != nil {
		return nil, err
	}

	user, err := i.User(ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	account, ok := user.GetAccount(req.GetAccountId())
	if !ok {
		return nil, status.Error(codes.NotFound, "account not found")
	}

	return &desc.AccountBalanceResponse{
		AccountId: account.AccountID,
		Amount:    account.Amount,
	}, nil
}
