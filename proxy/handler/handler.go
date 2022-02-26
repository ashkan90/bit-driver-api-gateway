package handler

import "net/http"

func NewHandler(target string) func(req *http.Request) {
	return func(req *http.Request) {
		req.URL.Host = target
		req.URL.Scheme = "http"
	}
}
