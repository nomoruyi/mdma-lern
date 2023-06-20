package main

import (
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", personHandler)
	mux.HandleFunc("/change-name", changeNameHandler)

	srv := http.Server{
		Handler: mux,
		Addr:    ":8080",
	}

	srv.ListenAndServe()
}
