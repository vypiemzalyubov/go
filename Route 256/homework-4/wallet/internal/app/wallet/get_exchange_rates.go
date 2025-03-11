package wallet

import (
	context "context"
	"time"

	desc "gitlab.ozon.dev/route256/wallet/pkg/api/wallet"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GetExchangeRates получить курсы валют
func (i *Implementation) GetExchangeRates(ctx context.Context, req *desc.GetExchangeRatesRequest) (*desc.GetExchangeRatesResponse, error) {
	date, err := time.Parse("02/01/2006", req.GetDate())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "date must be like pattern MM/DD/YYYY")
	}

	valutes, err := i.cbrClient.GetExchangeRates(ctx, date)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "can not get valutes from cbr.ru: %v", err)
	}

	var rates []*desc.GetExchangeRatesResponse_ExchangeRate
	for _, valute := range valutes.ValCurs {
		rates = append(rates, &desc.GetExchangeRatesResponse_ExchangeRate{
			Code:  valute.CharCode,
			Value: valute.Value,
			Name:  valute.Name,
		})
	}

	return &desc.GetExchangeRatesResponse{Rates: rates}, nil
}
