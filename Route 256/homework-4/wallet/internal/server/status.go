package server

import (
	"net/http"
	"sync/atomic"

	"gitlab.ozon.dev/route256/wallet/internal/config"
)

func createStatusServer(cfg config.Status, isReady *atomic.Value) *http.Server {
	mux := http.DefaultServeMux

	mux.HandleFunc(cfg.LivenessPath, livenessHandler)
	mux.HandleFunc(cfg.ReadinessPath, readinessHandler(isReady))

	//nolint:gosec
	statusServer := &http.Server{
		Addr:    cfg.Addr,
		Handler: mux,
	}

	return statusServer
}

func livenessHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func readinessHandler(isReady *atomic.Value) http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		if isReady == nil || !isReady.Load().(bool) {
			http.Error(w, http.StatusText(http.StatusServiceUnavailable), http.StatusServiceUnavailable)

			return
		}
		w.WriteHeader(http.StatusOK)
	}
}
