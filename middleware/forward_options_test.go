package middleware

import (
	"github.com/ashkan90/bit-driver-api-gateway/handler"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"testing"
)

func TestFwdOptions_ServingSuccessfully(t *testing.T) {
	var target = "localhost:1221"
	var path = "path-to-address"
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	_handler := FwdOptions(&httputil.ReverseProxy{
		Director: handler.NewHandler(target, path),
	})

	_handler(rec, req)

	assert.Equal(t, "", rec.Body.String())
}
