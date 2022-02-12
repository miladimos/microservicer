package test

import (
	"fmt"
	"net/http"
)

func RunServer(serverAddress string) error {
	http.HandleFunc("/", rootHandler)
	return http.ListerAndServe(serverAddress, nil)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to server : %s", r.RemoteAddr)
}