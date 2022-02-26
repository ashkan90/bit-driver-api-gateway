package main

import "net/http"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello from svc-1"))
		w.WriteHeader(200)
	})
	http.ListenAndServe(":8090", nil)
}
