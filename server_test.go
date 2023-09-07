package main

import (
	"github.com/seanb4t/example-movie-service/internal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http/httptest"
	"testing"
)

func TestDefaultRouteGoesToGraphiql(t *testing.T) {
	tp, err := internal.InitTracer()
	require.Nil(t, err, "err must be nil for tracer init")

	router := setupRouter(tp)
	require.NotNil(t, router, "router must not be nil")

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "<title>GraphQL</title>")
}

func TestEmptyGetQueryFails(t *testing.T) {
	tp, err := internal.InitTracer()
	require.Nil(t, err, "err must be nil for tracer init")

	router := setupRouter(tp)
	require.NotNil(t, router, "router must not be nil")

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/query", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 422, w.Code)
	assert.JSONEq(t, "{\"errors\":[{\"message\":\"no operation provided\",\"extensions\":{\"code\":\"GRAPHQL_VALIDATION_FAILED\"}}],\"data\":null}", w.Body.String())
}

func TestMetricsRoute(t *testing.T) {
	tp, err := internal.InitTracer()
	require.NoError(t, err, "err must be nil for tracer init")

	router := setupRouter(tp)
	require.NotNil(t, router, "router must not be nil")

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/metrics", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Regexp(t, "^# HELP go_gc_duration_seconds", w.Body.String())
}
