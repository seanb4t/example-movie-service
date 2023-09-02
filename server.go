package main

import (
	"context"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-contrib/logger"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog"
	"github.com/seanb4t/example-movie-service/graph"
	"github.com/seanb4t/example-movie-service/internal"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

// Logger is the global logger
var log = internal.Logger

// Tracer is the global tracer
var tracer = internal.Tracer

// Defining the Graphql handler
func graphqlHandler() gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	h := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	return func(gctx *gin.Context) {
		_, span := tracer.Start(gctx, "graphql-handler")
		defer span.End()

		defer h.ServeHTTP(gctx.Writer, gctx.Request)
	}
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(gctx *gin.Context) {
		_, span := tracer.Start(gctx, "playground-handler")
		defer span.End()

		h.ServeHTTP(gctx.Writer, gctx.Request)
	}
}

// GinContextToContextMiddleware Middleware for providing the gin context to GraphQL
func GinContextToContextMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), internal.GinContextKey, c)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func main() {

	tracerProvider, err := internal.InitTracer()
	if err != nil {
		log.Err(err).Msg("Failed to init tracer")
	}
	defer func() {
		if err := tracerProvider.Shutdown(context.Background()); err != nil {
			log.Printf("Error shutting down tracer provider: %v", err)
		}
	}()

	zerolog.DefaultContextLogger = &log

	log.Info().Msg("Starting server")

	// Setup Gin
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(requestid.New())
	r.Use(otelgin.Middleware("movie-service", otelgin.WithTracerProvider(tracerProvider)))
	r.Use(GinContextToContextMiddleware())
	r.Use(logger.SetLogger(logger.
		WithLogger(func(g *gin.Context, z zerolog.Logger) zerolog.Logger {
			// Add request context to logger
			ctx := context.WithValue(context.Background(), internal.RequestContextKey, g.Request.Context())
			return log.With().Ctx(ctx).Logger()
		})))
	r.POST("/query", graphqlHandler())
	r.GET("/", playgroundHandler())
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))
	r.Run()
}
