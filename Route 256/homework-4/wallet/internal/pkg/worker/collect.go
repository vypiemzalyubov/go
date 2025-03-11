package worker

import (
	"context"
	"sync/atomic"
	"time"

	"github.com/rs/zerolog/log"
	"gitlab.ozon.dev/route256/wallet/internal/config"
	"gitlab.ozon.dev/route256/wallet/internal/pkg/storage"
)

var (
	// Count .
	Count atomic.Int64
	// ForceSignal .
	ForceSignal = make(chan struct{})
)

type Collect struct{}

func NewCollect() *Collect {
	return &Collect{}
}

func (c *Collect) Start(ctx context.Context, repo storage.Storage, conf *config.AppConfig) {
	ticker := time.NewTicker(conf.Jobs.DurationCollectOperation)

	for {
		select {
		case <-ticker.C:

			count, err := repo.CollectOperation(ctx)
			if err != nil {
				log.Err(err).Msg("failed to collect operation")
			}

			log.Info().Msgf("%d all operations collected", count)

			old := Count.Swap(count)

			log.Info().Msgf("%d old operations collected", old)
		case <-ForceSignal:
			log.Info().Msgf("Start force collect operation")
			count, err := repo.CollectOperation(ctx)
			if err != nil {
				log.Err(err).Msg("failed to collect operation")
			}

			log.Info().Msgf("%d all operations collected", count)

			Count.Swap(count)
		case <-ctx.Done():
			ticker.Stop()
		}
	}
}
