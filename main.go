package main

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// for instrumentation details see here: https://prometheus.io/docs/tutorials/instrumenting_http_server_in_go/
var requestCounter = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "request_count",
		Help: "No of request handled by HelloServer handler",
	},
)

func main() {

	prometheus.MustRegister(requestCounter)

	http.HandleFunc("/", HelloServer)
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8080", nil)

}

// HelloServer: Handler for all URL paths, outputs "Hello, $(url path)!" to web page
func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World, %s!", r.URL.Path[1:])
}
