package wallet

import (
	"context"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/wiremock/go-wiremock"
	"gitlab.ozon.dev/route256/wallet/internal/config"
	"gitlab.ozon.dev/route256/wallet/internal/pkg/cbr"
	desc "gitlab.ozon.dev/route256/wallet/pkg/api/wallet"
)

const wiremockURL = "http://localhost:8070"

func TestGetExchangeRates_Positive(t *testing.T) {
	if os.Getenv("ENV") != "test" {
		t.Skip("Skipping test as ENV is not set to 'test'")
	}
	t.Run("Simplr get exchange rates", func(t *testing.T) {
		wiremockClient := wiremock.NewClient(wiremockURL)
		defer wiremockClient.Reset()

		wiremockClient.StubFor(wiremock.Get(wiremock.URLPathEqualTo("/scripts/XML_daily.asp")).
			WithQueryParam("date_req", wiremock.Matching("\\d{2}/\\d{2}/\\d{4}")).
			WillReturnResponse(
				wiremock.NewResponse().
					WithHeader("Content-Type", "application/xml").
					WithBody(`
						<ValCurs Date="10/01/2024" name="Foreign Currency Market">
							<Valute ID="R01235">
								<NumCode>840</NumCode>
								<CharCode>USD</CharCode>
								<Nominal>1</Nominal>
								<Name>Доллар США</Name>
								<Value>93,2221</Value>
							</Valute>
							<Valute ID="R01239">
								<NumCode>978</NumCode>
								<CharCode>EUR</CharCode>
								<Nominal>1</Nominal>
								<Name>Евро</Name>
								<Value>104,1735</Value>
							</Valute>
						</ValCurs>
					`).
					WithStatus(http.StatusOK),
			))

		cbrClient := cbr.NewClient(wiremockURL)
		impl := Implementation{
			cbrClient: cbrClient,
			store:     store,
		}

		ctx := context.Background()
		res, err := impl.GetExchangeRates(ctx, &desc.GetExchangeRatesRequest{Date: "10/01/2024"})
		require.NoError(t, err)
		require.NotNil(t, res)
	})
}

func TestGetExchangeRates_Negative(t *testing.T) {
	if os.Getenv("ENV") != "test" {
		t.Skip("Skipping test as ENV is not set to 'test'")
	}
	t.Run("Invalid date", func(t *testing.T) {
		cbrClient := cbr.NewClient(wiremockURL)
		impl := Implementation{
			cbrClient: cbrClient,
			store:     store,
		}

		ctx := context.Background()
		res, err := impl.GetExchangeRates(ctx, &desc.GetExchangeRatesRequest{Date: "invalid-date"})
		require.Error(t, err)
		require.Nil(t, res)
		require.Contains(t, err.Error(), "date must be like pattern MM/DD/YYYY")
	})
	t.Run("Internal CBR error", func(t *testing.T) {
		wiremockClient := wiremock.NewClient(wiremockURL)
		defer wiremockClient.Reset()

		wiremockClient.StubFor(wiremock.Get(wiremock.URLPathEqualTo("/daily_json.js")).
			WillReturnResponse(
				wiremock.NewResponse().
					WithHeader("Content-Type", "application/json").
					WithStatus(http.StatusInternalServerError),
			))

		cbrClient := cbr.NewClient(config.CbrURL)
		impl := Implementation{
			cbrClient: cbrClient,
			store:     store,
		}

		ctx := context.Background()
		res, err := impl.GetExchangeRates(ctx, &desc.GetExchangeRatesRequest{Date: "10/01/2024"})
		require.Error(t, err)
		require.Nil(t, res)
		require.Contains(t, err.Error(), "can not get valutes from cbr.ru")
	})
}
