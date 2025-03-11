package config

import (
	"time"

	"github.com/rs/zerolog/log"
)

const (
	maxRetryCount = 10
)

// RunWithRetry .
func RunWithRetry(funcName string, fn func() error) error {
	var err error
	for i := 0; i < maxRetryCount; i++ {
		err = fn()
		if err == nil {
			break
		}

		log.Info().Err(err).Msgf("Retry: call func '%s' an error occurred", funcName)
		time.Sleep(1 * time.Second)
		continue
	}

	return err
}
