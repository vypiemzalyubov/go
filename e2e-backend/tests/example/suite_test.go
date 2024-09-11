package example

import (
	"testing"

	"e2e-backend/internal/clients/grpccli"
	"e2e-backend/internal/pb/gitlab.ozon.dev/route256/wallet"
	"e2e-backend/internal/steps/grpcsteps"

	"google.golang.org/grpc"

	"github.com/stretchr/testify/suite"
)

type ExampleSiute struct {
	suite.Suite

	client wallet.WalletClient
	conn   *grpc.ClientConn

	UserSteps grpcsteps.UserSteps
}

func TestExampleSiute(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(ExampleSiute))
}

func (s *ExampleSiute) SetupSuite() {
	s.client, s.conn = grpccli.NewWalletClient()

	s.UserSteps = *grpcsteps.NewUserSteps(s.client)
}

func (s *ExampleSiute) TearDownSuite() {
	s.conn.Close()
}

func (s *ExampleSiute) BeforeTest(suiteName, testName string) {
	//  TODO
}

func (s *ExampleSiute) AfterTest(suiteName, testName string) {
	//  TODO
}

func (s *ExampleSiute) SetupTest() {
	// TODO
}

func (s *ExampleSiute) TearDownTest() {
	// TODO
}
