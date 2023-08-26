package graph

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/seanb4t/example-movie-service/internal"
)

func (r *Resolver) ginContext(ctx context.Context) (*gin.Context, error) {
	gc, err := internal.GinContextFromContext(ctx)
	if err != nil {
		return nil, err
	}

	log.Info().
		Str("url", gc.Request.RequestURI).
		Msg("Context retrieved from gin.Context")

	return gc, nil
}
