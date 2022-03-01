package middleware

import (
	"github.com/ashkan90/bit-driver-api-gateway/handler"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"testing"
)

func TestCheckAuth_ServingCheckAuthStrategyButItFailsCauseOfEmptyHeader(t *testing.T) {
	var target = "localhost:1221"
	var listenPath = "/some-svc/"
	var path = "path-to-address"
	req := httptest.NewRequest(http.MethodGet, listenPath+path, nil)
	rec := httptest.NewRecorder()

	_handler := CheckAuth(&httputil.ReverseProxy{
		Director: handler.NewHandler(target, path),
	})

	_handler(rec, req)

	assert.Equal(t, "Auth needed", rec.Body.String())
}

func TestCheckAuth_ServingCheckAuthStrategyWorksSuccessfully(t *testing.T) {
	var target = "localhost:1221"
	var path = "path-to-address"
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("Authorization", "Bearer authenticated: true")
	rec := httptest.NewRecorder()

	_handler := CheckAuth(&httputil.ReverseProxy{
		Director: handler.NewHandler(target, path),
	})

	_handler(rec, req)

	assert.Equal(t, "", rec.Body.String())
}
