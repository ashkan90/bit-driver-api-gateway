package middleware

import (
	"net/http"
	"net/http/httputil"
)

func FwdOptions(reverse *httputil.ReverseProxy) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reverse.ServeHTTP(w, r)
	}
}
