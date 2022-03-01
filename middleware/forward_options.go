package middleware

import (
	"net/http"
	"net/http/httputil"
)

// FwdOptions implemented to bypass requests which has OPTIONS method
// otherwise requests with OPTIONS method has to go by CheckAuth strategy.
func FwdOptions(reverse *httputil.ReverseProxy) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reverse.ServeHTTP(w, r)
	}
}
