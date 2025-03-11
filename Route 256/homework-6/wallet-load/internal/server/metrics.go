package server

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"gitlab.ozon.dev/route256/wallet/internal/config"
)

func createMetricsServer(cfg config.Metrics) *http.Server {
	mux := http.DefaultServeMux
	mux.Handle(cfg.Path, promhttp.Handler())

	//nolint:gosec
	metricsServer := &http.Server{
		Addr:    cfg.Addr,
		Handler: mux,
	}

	return metricsServer
}
