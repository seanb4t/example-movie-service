package internal

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
)

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
