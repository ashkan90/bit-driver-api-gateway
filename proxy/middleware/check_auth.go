package middleware

import (
	"net/http"
	"net/http/httputil"
)

func CheckAuth(reverse *httputil.ReverseProxy) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//var auth = r.Header.Get("Auth")
		//
		//if auth != "31" {
		//	w.WriteHeader(http.StatusForbidden)
		//	io.WriteString(w, "Auth needed\n")
		//	return
		//}

		reverse.ServeHTTP(w, r)
	}
}
