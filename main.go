package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		host, _ := os.Hostname()
		fmt.Fprintf(w, "hello from shipd-example v2\nhostname: %s\ntime: %s\npath: %s\n",
			host, time.Now().UTC().Format(time.RFC3339), r.URL.Path)
	})
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	http.ListenAndServe(":"+port, mux)
}
