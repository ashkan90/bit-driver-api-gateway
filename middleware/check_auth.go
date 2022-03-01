package middleware

import (
	"io"
	"net/http"
	"net/http/httputil"
)

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
