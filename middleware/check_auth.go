package middleware

import (
	"io"
	"net/http"
	"net/http/httputil"
)

// CheckAuth checks the request header and determines
// this request is allowed to proxy. Works before proxy.
func CheckAuth(reverse *httputil.ReverseProxy) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var header = r.Header.Get("Authorization")

		if header != "Bearer authenticated: true" {
			w.WriteHeader(http.StatusForbidden)
			_, _ = io.WriteString(w, "Auth needed")
			return
		}

		reverse.ServeHTTP(w, r)
	}
}
