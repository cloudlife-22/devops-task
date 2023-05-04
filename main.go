package main

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", nil)
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":8080", nil)

}

// HelloServer: Handler for all URL paths, outputs "Hello, $(url path)!" to web page
func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}
