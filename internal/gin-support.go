package internal

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
)

// GinContextKey is the key used to store the gin context in the request context
const GinContextKey = "ginContext"

// GinContextFromContext returns the gin context from the context
func GinContextFromContext(ctx context.Context) (*gin.Context, error) {
	ginContext := ctx.Value(GinContextKey)
	if ginContext == nil {
		return nil, fmt.Errorf("unable to get gin.Context")
	}

	gc, ok := ginContext.(*gin.Context)
	if !ok {
		return nil, fmt.Errorf("gin.Context has the wrong type")
	}
	return gc, nil
}
