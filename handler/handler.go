package handler

import (
	"log"
	"net/http"
)

func NewHandler(target, path string) func(req *http.Request) {
	return func(req *http.Request) {
		req.URL.Path = path
		req.URL.Host = target
		req.URL.Scheme = "http"

		log.Printf("An request came through %s path by %s", target, path)
	}
}
