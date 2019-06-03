package main

import (
	"log"

	"net/http"
)

// Start Server
func main() {
	fs := http.FileServer(http.Dir("./build"))
	log.Println("[wasm-xlxs] server started: http://localhost:8180")
	http.ListenAndServe(":8180", http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		if req.URL.Path == "/test.wasm" {
			resp.Header().Set("content-type", "application/wasm")
		}

		fs.ServeHTTP(resp, req)
	}))
}