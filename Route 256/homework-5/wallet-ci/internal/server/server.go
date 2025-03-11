package server

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"
	"time"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpcrecovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	wallet "gitlab.ozon.dev/route256/wallet/internal/app/wallet"
	"gitlab.ozon.dev/route256/wallet/internal/config"
	"gitlab.ozon.dev/route256/wallet/internal/pkg/kafka"
	"gitlab.ozon.dev/route256/wallet/internal/pkg/storage"
	"gitlab.ozon.dev/route256/wallet/internal/pkg/worker"
	desc "gitlab.ozon.dev/route256/wallet/pkg/api/wallet"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// Server .
type Server struct {
	db *sqlx.DB
}

// NewServer .
func NewServer(db *sqlx.DB) *Server {
	return &Server{
		db: db,
	}
}

// Start method runs server
func (s *Server) Start(cfg *config.AppConfig) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	gatewayServer, err := createGatewayServer(ctx, cfg)

	if err != nil {
		log.Error().Err(err).Msg("Failed create gateway server")
	}
	go func() {
		log.Info().Msgf("Gateway server is running on %s", cfg.Rest.Addr)
		if err = gatewayServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Error().Err(err).Msg("Failed running gateway server")
			cancel()
		}
	}()

	metricsServer := createMetricsServer(cfg.Metrics)
	go func() {
		log.Info().Msgf("Metrics server is running on %s", cfg.Metrics.Addr)
		if err = metricsServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Error().Err(err).Msg("Failed running metrics server")
			cancel()
		}
	}()

	isReady := &atomic.Value{}
	isReady.Store(false)

	statusServer := createStatusServer(cfg.Status, isReady)
	go func() {
		log.Info().Msgf("Status server is running on %s", cfg.Status.Addr)
		if err = statusServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Error().Err(err).Msg("Failed running status server")
		}
	}()

	l, err := net.Listen("tcp", cfg.Grpc.Addr)
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}
	defer l.Close()

	grpcServer := grpc.NewServer(
		grpc.KeepaliveParams(cfg.Grpc.ServerParameters),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_ctxtags.UnaryServerInterceptor(),
			grpc_prometheus.UnaryServerInterceptor,
			loggingInterceptor(),
			grpcrecovery.UnaryServerInterceptor(),
		)),
	)

	repo := storage.New(s.db)
	collect := worker.NewCollect()

	kfk := kafka.NewClient(cfg.Kafka)
	go kfk.ConsumeMessages(ctx)

	desc.RegisterWalletServer(grpcServer, wallet.NewWallet(repo, kfk))
	grpc_prometheus.EnableHandlingTimeHistogram()
	grpc_prometheus.Register(grpcServer)

	go collect.Start(ctx, repo, cfg)

	go func() {
		log.Info().Msg("Turn on reflection.Register")
		reflection.Register(grpcServer)

		log.Info().Msgf("GRPC Server is listening on: %s", cfg.Grpc.Addr)
		if err := grpcServer.Serve(l); err != nil {
			cancel()
			log.Fatal().Err(err).Msg("Failed running gRPC server")
		}
	}()

	go func() {
		time.Sleep(2 * time.Second)
		isReady.Store(true)
		log.Info().Msg("The service is ready to accept requests")
	}()

	//graceful Shutdown
	{
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

		select {
		case v := <-quit:
			log.Info().Msgf("signal.Notify: %v", v)
		case done := <-ctx.Done():
			log.Info().Msgf("ctx.Done: %v", done)
		}

		if err := kfk.Stop(); err != nil {
			log.Error().Err(err).Msg("Failed to stop kafka server")
		}
		isReady.Store(false)
		gracefulStop(ctx, grpcServer, gatewayServer, statusServer, metricsServer)
	}
	return nil
}

func gracefulStop(ctx context.Context, grpcServer *grpc.Server, gatewayServer, statusServer, metricsServer *http.Server) {
	if err := gatewayServer.Shutdown(ctx); err != nil {
		log.Error().Err(err).Msg("gatewayServer.Shutdown")
	} else {
		log.Info().Msg("gatewayServer shut down correctly")
	}

	if err := statusServer.Shutdown(ctx); err != nil {
		log.Error().Err(err).Msg("statusServer.Shutdown")
	} else {
		log.Info().Msg("statusServer shut down correctly")
	}

	if err := metricsServer.Shutdown(ctx); err != nil {
		log.Error().Err(err).Msg("metricsServer.Shutdown")
	} else {
		log.Info().Msg("metricsServer shut down correctly")
	}

	grpcServer.GracefulStop()
	log.Info().Msgf("grpcServer shut down correctly")
}
