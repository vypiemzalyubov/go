package server

import (
	"context"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

func loggingInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (_ interface{}, err error) {
		log.Debug().Msgf("%s.Request: %+v", info.FullMethod, req)
		res, err := handler(ctx, req)
		if err != nil {
			log.Error().Msgf("%s.Error: %+v", info.FullMethod, err)
		} else {
			log.Debug().Msgf("%s.Response: %+v", info.FullMethod, res)
		}
		return res, err
	}
}
